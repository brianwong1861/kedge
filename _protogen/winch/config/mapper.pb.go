// Code generated by protoc-gen-go. DO NOT EDIT.
// source: winch/config/mapper.proto

/*
Package winch_config is a generated protocol buffer package.

It is generated from these files:
	winch/config/mapper.proto

It has these top-level messages:
	MapperConfig
	Route
	DirectRoute
	RegexpRoute
*/
package winch_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// / Config is the top level configuration message for a winch mapper.
type MapperConfig struct {
	Routes []*Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
}

func (m *MapperConfig) Reset()                    { *m = MapperConfig{} }
func (m *MapperConfig) String() string            { return proto.CompactTextString(m) }
func (*MapperConfig) ProtoMessage()               {}
func (*MapperConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MapperConfig) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	// Types that are valid to be assigned to Type:
	//	*Route_Direct
	//	*Route_Regexp
	Type isRoute_Type `protobuf_oneof:"type"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isRoute_Type interface {
	isRoute_Type()
}

type Route_Direct struct {
	Direct *DirectRoute `protobuf:"bytes,1,opt,name=direct,oneof"`
}
type Route_Regexp struct {
	Regexp *RegexpRoute `protobuf:"bytes,2,opt,name=regexp,oneof"`
}

func (*Route_Direct) isRoute_Type() {}
func (*Route_Regexp) isRoute_Type() {}

func (m *Route) GetType() isRoute_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Route) GetDirect() *DirectRoute {
	if x, ok := m.GetType().(*Route_Direct); ok {
		return x.Direct
	}
	return nil
}

func (m *Route) GetRegexp() *RegexpRoute {
	if x, ok := m.GetType().(*Route_Regexp); ok {
		return x.Regexp
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Route) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Route_OneofMarshaler, _Route_OneofUnmarshaler, _Route_OneofSizer, []interface{}{
		(*Route_Direct)(nil),
		(*Route_Regexp)(nil),
	}
}

func _Route_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Route)
	// type
	switch x := m.Type.(type) {
	case *Route_Direct:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Direct); err != nil {
			return err
		}
	case *Route_Regexp:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Regexp); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Route.Type has unexpected type %T", x)
	}
	return nil
}

func _Route_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Route)
	switch tag {
	case 1: // type.direct
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DirectRoute)
		err := b.DecodeMessage(msg)
		m.Type = &Route_Direct{msg}
		return true, err
	case 2: // type.regexp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RegexpRoute)
		err := b.DecodeMessage(msg)
		m.Type = &Route_Regexp{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Route_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Route)
	// type
	switch x := m.Type.(type) {
	case *Route_Direct:
		s := proto.Size(x.Direct)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Route_Regexp:
		s := proto.Size(x.Regexp)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// / Simplest routing mechanism using just direct mapping between dns and kedge target.
type DirectRoute struct {
	Key      string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	KedgeUrl string `protobuf:"bytes,2,opt,name=kedge_url,json=kedgeUrl" json:"kedge_url,omitempty"`
}

func (m *DirectRoute) Reset()                    { *m = DirectRoute{} }
func (m *DirectRoute) String() string            { return proto.CompactTextString(m) }
func (*DirectRoute) ProtoMessage()               {}
func (*DirectRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DirectRoute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DirectRoute) GetKedgeUrl() string {
	if m != nil {
		return m.KedgeUrl
	}
	return ""
}

type RegexpRoute struct {
	// Regexp expression that will be applied on given DNS.
	Exp string `protobuf:"bytes,1,opt,name=exp" json:"exp,omitempty"`
	// Kedge URL to be used if we have a match. It can be a string including two variables using
	// go template:
	// - {{ .Cluster }} if cluster_group_name is specified.
	KedgeUrl string `protobuf:"bytes,2,opt,name=kedge_url,json=kedgeUrl" json:"kedge_url,omitempty"`
	// If specified, target cluster name will be fetched from named regexp group.
	ClusterGroupName string `protobuf:"bytes,3,opt,name=cluster_group_name,json=clusterGroupName" json:"cluster_group_name,omitempty"`
}

func (m *RegexpRoute) Reset()                    { *m = RegexpRoute{} }
func (m *RegexpRoute) String() string            { return proto.CompactTextString(m) }
func (*RegexpRoute) ProtoMessage()               {}
func (*RegexpRoute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegexpRoute) GetExp() string {
	if m != nil {
		return m.Exp
	}
	return ""
}

func (m *RegexpRoute) GetKedgeUrl() string {
	if m != nil {
		return m.KedgeUrl
	}
	return ""
}

func (m *RegexpRoute) GetClusterGroupName() string {
	if m != nil {
		return m.ClusterGroupName
	}
	return ""
}

func init() {
	proto.RegisterType((*MapperConfig)(nil), "winch.config.MapperConfig")
	proto.RegisterType((*Route)(nil), "winch.config.Route")
	proto.RegisterType((*DirectRoute)(nil), "winch.config.DirectRoute")
	proto.RegisterType((*RegexpRoute)(nil), "winch.config.RegexpRoute")
}

func init() { proto.RegisterFile("winch/config/mapper.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x18, 0x86, 0x4d, 0xab, 0xc1, 0x5e, 0x3a, 0xc8, 0xb9, 0xa4, 0x2e, 0x96, 0xb8, 0x14, 0xb4, 0x09,
	0xb4, 0xe0, 0xe2, 0x56, 0x05, 0x1d, 0xd4, 0xe1, 0xc0, 0xb9, 0xa4, 0xe9, 0xe7, 0xf5, 0x68, 0x92,
	0x3b, 0x2e, 0x17, 0xd3, 0xf8, 0x67, 0x05, 0x7f, 0x89, 0xdc, 0x97, 0x50, 0xa3, 0x38, 0xb8, 0x85,
	0xef, 0x79, 0xdf, 0xe7, 0x0d, 0x1c, 0x19, 0x55, 0x22, 0x4f, 0x36, 0x51, 0x22, 0xf3, 0x57, 0xc1,
	0xa3, 0x2c, 0x56, 0x0a, 0x74, 0xa8, 0xb4, 0x34, 0x92, 0x0e, 0x11, 0x85, 0x0d, 0x3a, 0xbb, 0xe6,
	0xc2, 0x6c, 0xca, 0x55, 0x98, 0xc8, 0x2c, 0xca, 0x2a, 0x61, 0xb6, 0xb2, 0x8a, 0xb8, 0x9c, 0x62,
	0x74, 0xfa, 0x16, 0xa7, 0x62, 0x1d, 0x1b, 0xa9, 0x8b, 0x68, 0xff, 0xd9, 0x58, 0x82, 0x1b, 0x32,
	0x7c, 0x42, 0xeb, 0x2d, 0x7a, 0xe8, 0x25, 0x71, 0xb5, 0x2c, 0x0d, 0x14, 0xbe, 0x33, 0xee, 0x4f,
	0xbc, 0xd9, 0x69, 0xd8, 0x9d, 0x09, 0x99, 0x65, 0xac, 0x8d, 0x04, 0x35, 0x39, 0xc2, 0x03, 0x9d,
	0x13, 0x77, 0x2d, 0x34, 0x24, 0xc6, 0x77, 0xc6, 0xce, 0xc4, 0x9b, 0x8d, 0x7e, 0xb6, 0xee, 0x90,
	0x61, 0xf4, 0xe1, 0x80, 0xb5, 0x51, 0x5b, 0xd2, 0xc0, 0x61, 0xa7, 0xfc, 0xde, 0x5f, 0x25, 0x86,
	0x6c, 0x5f, 0x6a, 0xa2, 0x0b, 0x97, 0x1c, 0x9a, 0x5a, 0x41, 0xf0, 0x48, 0xbc, 0x8e, 0x95, 0xfa,
	0xa4, 0xbf, 0x85, 0x1a, 0xd7, 0x07, 0x0b, 0xf7, 0xf3, 0xe3, 0xbc, 0x37, 0x76, 0x98, 0x3d, 0xd1,
	0x0b, 0x32, 0xd8, 0xc2, 0x9a, 0xc3, 0xb2, 0xd4, 0x29, 0x0e, 0x7d, 0xf3, 0x63, 0x04, 0x2f, 0x3a,
	0x0d, 0xde, 0x89, 0xd7, 0x99, 0xb3, 0x36, 0xfb, 0x5b, 0xbf, 0x6c, 0xb0, 0x53, 0xff, 0xb2, 0xd1,
	0x2b, 0x42, 0x93, 0xb4, 0x2c, 0x0c, 0xe8, 0x25, 0xd7, 0xb2, 0x54, 0xcb, 0x3c, 0xce, 0xc0, 0xef,
	0xdb, 0x34, 0x3b, 0x69, 0xc9, 0xbd, 0x05, 0xcf, 0x71, 0x06, 0x2b, 0x17, 0x1f, 0x62, 0xfe, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xd5, 0x26, 0x21, 0xd4, 0xeb, 0x01, 0x00, 0x00,
}