package backendpool

import (
	"hash/fnv"
	"sync"

	"github.com/improbable-eng/kedge/pkg/metrics"
	pb "github.com/improbable-eng/kedge/protogen/kedge/config/grpc/backends"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// dynamic is a Pool to which you can update or remove routes.
type dynamic struct {
	backends       map[string]*backend
	mu             sync.RWMutex
	backendFactory func(backend *pb.Backend) (*backend, error)
	logger         logrus.FieldLogger
}

func (s *dynamic) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, be := range s.backends {
		be.Close()
	}
	return nil
}

// NewDynamic creates a pool with a dynamic allocator
func NewDynamic(logger logrus.FieldLogger) *dynamic {
	s := &dynamic{backends: make(map[string]*backend), backendFactory: newBackend, logger: logger}
	return s
}

func (s *dynamic) Conn(backendName string) (*grpc.ClientConn, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	be, ok := s.backends[backendName]
	if !ok {
		return nil, ErrUnknownBackend
	}
	return be.Conn()
}

// AddOrUpdate checks tries to perform the least destructive operation of adding a new backend.
//
// If a backend of a given name already exists, and the configuration hasn't changed, no new work will be done.
// If a backend requires changes, the previous one will be removed and closed.
func (s *dynamic) AddOrUpdate(config *pb.Backend, logTestResolution bool) (changed bool, err error) {
	s.mu.RLock()
	existing, ok := s.backends[config.Name]
	s.mu.RUnlock()

	defer func() {
		if changed && logTestResolution {
			go s.backends[config.Name].LogTestResolution(
				s.logger.WithField("protocol", "grpc").WithField("backend", config.Name),
			)
		}
	}()

	if !ok {
		err = s.addNewBackend(config)
		if err != nil {
			return changed, err
		}
		changed = true
		s.logger.Infof("Adding new grpc backend: %v", config.Name)
		metrics.BackendGRPCConfigurationCounter.WithLabelValues(config.Name, metrics.ConfiguationActionCreate).Inc()
		return changed, nil
	}

	var updated bool
	updated, err = s.updateBackendWithDiffing(existing, config)
	if err != nil {
		return changed, err
	}
	if updated {
		changed = true
		s.logger.Infof("Updated grpc backend: %v", config.Name)
		metrics.BackendGRPCConfigurationCounter.WithLabelValues(config.Name, metrics.ConfiguationActionChange).Inc()
	}
	return changed, nil
}

func (s *dynamic) addNewBackend(config *pb.Backend) error {
	be, err := s.backendFactory(config)
	if err != nil {
		return err
	}
	s.mu.Lock()
	s.backends[config.Name] = be
	s.mu.Unlock()
	return nil
}

func (s *dynamic) updateBackendWithDiffing(existing *backend, config *pb.Backend) (changed bool, err error) {
	if configsAreTheSame(existing.config, config) {
		return false, nil
	}
	if err := s.addNewBackend(config); err != nil {
		return true, err
	}
	// Make sure we clear up resources.
	existing.Close()
	return true, nil
}

// Remove removes and shuts down a previously active backend.
func (s *dynamic) Remove(backendName string) error {
	s.mu.RLock()
	existing, ok := s.backends[backendName]
	s.mu.RUnlock()
	if !ok {
		return ErrUnknownBackend
	}
	s.mu.Lock()
	delete(s.backends, backendName)
	s.mu.Unlock()
	existing.Close()

	s.logger.Infof("Removed grpc backend: %v", backendName)
	metrics.BackendGRPCConfigurationCounter.WithLabelValues(backendName, metrics.ConfiguationActionDelete).Inc()
	return nil
}

// Configs returns a map of all active backends and their configuration.
func (s *dynamic) Configs() map[string]*pb.Backend {
	ret := make(map[string]*pb.Backend)
	s.mu.RLock()
	for k, v := range s.backends {
		ret[k] = v.config
	}
	s.mu.RUnlock()
	return ret
}

func configsAreTheSame(c1 *pb.Backend, c2 *pb.Backend) bool {
	h1 := fnv.New64a()
	h2 := fnv.New64a()
	h1.Write([]byte(c1.String()))
	h2.Write([]byte(c2.String()))
	return h1.Sum64() == h2.Sum64()
}
