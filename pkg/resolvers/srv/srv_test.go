package srvresolver

import (
	"testing"

	"github.com/improbable-eng/go-srvlb/srv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDomain = "domain.org"

type testLookup struct {
	t *testing.T
}

func (l *testLookup) Lookup(domainName string) ([]*srv.Target, error) {
	require.Equal(l.t, testDomain, domainName)
	return []*srv.Target{
		{
			DialAddr: "1.1.1.1:80",
			Ttl:      resolutionTTL,
		},
		{
			DialAddr: "1.1.1.2",
			Ttl:      resolutionTTL,
		},
		{
			DialAddr: "1.1.1.10:81",
			Ttl:      resolutionTTL,
		},
	}, nil
}

func TestPortOverrideSRVResolver_Lookup(t *testing.T) {
	l := &testLookup{
		t: t,
	}

	p := newPortOverrideSRVResolver(99, l)
	targets, err := p.Lookup(testDomain)
	require.NoError(t, err)

	assert.Equal(t, "1.1.1.1:99", targets[0].DialAddr)
	assert.Equal(t, resolutionTTL, targets[0].Ttl)
	assert.Equal(t, "1.1.1.2:99", targets[1].DialAddr)
	assert.Equal(t, resolutionTTL, targets[0].Ttl)
	assert.Equal(t, "1.1.1.10:99", targets[2].DialAddr)
	assert.Equal(t, resolutionTTL, targets[0].Ttl)
}
