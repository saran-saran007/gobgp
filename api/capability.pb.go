// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capability.proto

package gobgpapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AddPathMode int32

const (
	AddPathMode_MODE_NONE    AddPathMode = 0
	AddPathMode_MODE_RECEIVE AddPathMode = 1
	AddPathMode_MODE_SEND    AddPathMode = 2
	AddPathMode_MODE_BOTH    AddPathMode = 3
)

var AddPathMode_name = map[int32]string{
	0: "MODE_NONE",
	1: "MODE_RECEIVE",
	2: "MODE_SEND",
	3: "MODE_BOTH",
}
var AddPathMode_value = map[string]int32{
	"MODE_NONE":    0,
	"MODE_RECEIVE": 1,
	"MODE_SEND":    2,
	"MODE_BOTH":    3,
}

func (x AddPathMode) String() string {
	return proto.EnumName(AddPathMode_name, int32(x))
}
func (AddPathMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type MultiProtocolCapability struct {
	Family *Family `protobuf:"bytes,1,opt,name=family" json:"family,omitempty"`
}

func (m *MultiProtocolCapability) Reset()                    { *m = MultiProtocolCapability{} }
func (m *MultiProtocolCapability) String() string            { return proto.CompactTextString(m) }
func (*MultiProtocolCapability) ProtoMessage()               {}
func (*MultiProtocolCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MultiProtocolCapability) GetFamily() *Family {
	if m != nil {
		return m.Family
	}
	return nil
}

type RouteRefreshCapability struct {
}

func (m *RouteRefreshCapability) Reset()                    { *m = RouteRefreshCapability{} }
func (m *RouteRefreshCapability) String() string            { return proto.CompactTextString(m) }
func (*RouteRefreshCapability) ProtoMessage()               {}
func (*RouteRefreshCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type CarryingLabelInfoCapability struct {
}

func (m *CarryingLabelInfoCapability) Reset()                    { *m = CarryingLabelInfoCapability{} }
func (m *CarryingLabelInfoCapability) String() string            { return proto.CompactTextString(m) }
func (*CarryingLabelInfoCapability) ProtoMessage()               {}
func (*CarryingLabelInfoCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type ExtendedNexthopCapabilityTuple struct {
	NlriFamily *Family `protobuf:"bytes,1,opt,name=nlri_family,json=nlriFamily" json:"nlri_family,omitempty"`
	// Nexthop AFI must be either
	// gobgp.IPv4 or
	// gobgp.IPv6.
	NexthopFamily *Family `protobuf:"bytes,2,opt,name=nexthop_family,json=nexthopFamily" json:"nexthop_family,omitempty"`
}

func (m *ExtendedNexthopCapabilityTuple) Reset()                    { *m = ExtendedNexthopCapabilityTuple{} }
func (m *ExtendedNexthopCapabilityTuple) String() string            { return proto.CompactTextString(m) }
func (*ExtendedNexthopCapabilityTuple) ProtoMessage()               {}
func (*ExtendedNexthopCapabilityTuple) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ExtendedNexthopCapabilityTuple) GetNlriFamily() *Family {
	if m != nil {
		return m.NlriFamily
	}
	return nil
}

func (m *ExtendedNexthopCapabilityTuple) GetNexthopFamily() *Family {
	if m != nil {
		return m.NexthopFamily
	}
	return nil
}

type ExtendedNexthopCapability struct {
	Tuples []*ExtendedNexthopCapabilityTuple `protobuf:"bytes,1,rep,name=tuples" json:"tuples,omitempty"`
}

func (m *ExtendedNexthopCapability) Reset()                    { *m = ExtendedNexthopCapability{} }
func (m *ExtendedNexthopCapability) String() string            { return proto.CompactTextString(m) }
func (*ExtendedNexthopCapability) ProtoMessage()               {}
func (*ExtendedNexthopCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ExtendedNexthopCapability) GetTuples() []*ExtendedNexthopCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type GracefulRestartCapabilityTuple struct {
	Family *Family `protobuf:"bytes,1,opt,name=family" json:"family,omitempty"`
	Flags  uint32  `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
}

func (m *GracefulRestartCapabilityTuple) Reset()                    { *m = GracefulRestartCapabilityTuple{} }
func (m *GracefulRestartCapabilityTuple) String() string            { return proto.CompactTextString(m) }
func (*GracefulRestartCapabilityTuple) ProtoMessage()               {}
func (*GracefulRestartCapabilityTuple) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GracefulRestartCapabilityTuple) GetFamily() *Family {
	if m != nil {
		return m.Family
	}
	return nil
}

func (m *GracefulRestartCapabilityTuple) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type GracefulRestartCapability struct {
	Flags  uint32                            `protobuf:"varint,1,opt,name=flags" json:"flags,omitempty"`
	Time   uint32                            `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Tuples []*GracefulRestartCapabilityTuple `protobuf:"bytes,3,rep,name=tuples" json:"tuples,omitempty"`
}

func (m *GracefulRestartCapability) Reset()                    { *m = GracefulRestartCapability{} }
func (m *GracefulRestartCapability) String() string            { return proto.CompactTextString(m) }
func (*GracefulRestartCapability) ProtoMessage()               {}
func (*GracefulRestartCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *GracefulRestartCapability) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *GracefulRestartCapability) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *GracefulRestartCapability) GetTuples() []*GracefulRestartCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type FourOctetASNumberCapability struct {
	As uint32 `protobuf:"varint,1,opt,name=as" json:"as,omitempty"`
}

func (m *FourOctetASNumberCapability) Reset()                    { *m = FourOctetASNumberCapability{} }
func (m *FourOctetASNumberCapability) String() string            { return proto.CompactTextString(m) }
func (*FourOctetASNumberCapability) ProtoMessage()               {}
func (*FourOctetASNumberCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *FourOctetASNumberCapability) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

type AddPathCapabilityTuple struct {
	Family *Family     `protobuf:"bytes,1,opt,name=family" json:"family,omitempty"`
	Mode   AddPathMode `protobuf:"varint,2,opt,name=mode,enum=gobgpapi.AddPathMode" json:"mode,omitempty"`
}

func (m *AddPathCapabilityTuple) Reset()                    { *m = AddPathCapabilityTuple{} }
func (m *AddPathCapabilityTuple) String() string            { return proto.CompactTextString(m) }
func (*AddPathCapabilityTuple) ProtoMessage()               {}
func (*AddPathCapabilityTuple) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *AddPathCapabilityTuple) GetFamily() *Family {
	if m != nil {
		return m.Family
	}
	return nil
}

func (m *AddPathCapabilityTuple) GetMode() AddPathMode {
	if m != nil {
		return m.Mode
	}
	return AddPathMode_MODE_NONE
}

type AddPathCapability struct {
	Tuples []*AddPathCapabilityTuple `protobuf:"bytes,1,rep,name=tuples" json:"tuples,omitempty"`
}

func (m *AddPathCapability) Reset()                    { *m = AddPathCapability{} }
func (m *AddPathCapability) String() string            { return proto.CompactTextString(m) }
func (*AddPathCapability) ProtoMessage()               {}
func (*AddPathCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *AddPathCapability) GetTuples() []*AddPathCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type EnhancedRouteRefreshCapability struct {
}

func (m *EnhancedRouteRefreshCapability) Reset()         { *m = EnhancedRouteRefreshCapability{} }
func (m *EnhancedRouteRefreshCapability) String() string { return proto.CompactTextString(m) }
func (*EnhancedRouteRefreshCapability) ProtoMessage()    {}
func (*EnhancedRouteRefreshCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{10}
}

type LongLivedGracefulRestartCapabilityTuple struct {
	Family *Family `protobuf:"bytes,1,opt,name=family" json:"family,omitempty"`
	Flags  uint32  `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
	Time   uint32  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
}

func (m *LongLivedGracefulRestartCapabilityTuple) Reset() {
	*m = LongLivedGracefulRestartCapabilityTuple{}
}
func (m *LongLivedGracefulRestartCapabilityTuple) String() string { return proto.CompactTextString(m) }
func (*LongLivedGracefulRestartCapabilityTuple) ProtoMessage()    {}
func (*LongLivedGracefulRestartCapabilityTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{11}
}

func (m *LongLivedGracefulRestartCapabilityTuple) GetFamily() *Family {
	if m != nil {
		return m.Family
	}
	return nil
}

func (m *LongLivedGracefulRestartCapabilityTuple) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *LongLivedGracefulRestartCapabilityTuple) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type LongLivedGracefulRestartCapability struct {
	Tuples []*LongLivedGracefulRestartCapabilityTuple `protobuf:"bytes,1,rep,name=tuples" json:"tuples,omitempty"`
}

func (m *LongLivedGracefulRestartCapability) Reset()         { *m = LongLivedGracefulRestartCapability{} }
func (m *LongLivedGracefulRestartCapability) String() string { return proto.CompactTextString(m) }
func (*LongLivedGracefulRestartCapability) ProtoMessage()    {}
func (*LongLivedGracefulRestartCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{12}
}

func (m *LongLivedGracefulRestartCapability) GetTuples() []*LongLivedGracefulRestartCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type RouteRefreshCiscoCapability struct {
}

func (m *RouteRefreshCiscoCapability) Reset()                    { *m = RouteRefreshCiscoCapability{} }
func (m *RouteRefreshCiscoCapability) String() string            { return proto.CompactTextString(m) }
func (*RouteRefreshCiscoCapability) ProtoMessage()               {}
func (*RouteRefreshCiscoCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

type FQDNCapability struct {
	HostNameLen   uint32 `protobuf:"varint,1,opt,name=host_name_len,json=hostNameLen" json:"host_name_len,omitempty"`
	HostName      string `protobuf:"bytes,2,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	DomainNameLen uint32 `protobuf:"varint,3,opt,name=domain_name_len,json=domainNameLen" json:"domain_name_len,omitempty"`
	DomainName    string `protobuf:"bytes,4,opt,name=domain_name,json=domainName" json:"domain_name,omitempty"`
}

func (m *FQDNCapability) Reset()                    { *m = FQDNCapability{} }
func (m *FQDNCapability) String() string            { return proto.CompactTextString(m) }
func (*FQDNCapability) ProtoMessage()               {}
func (*FQDNCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *FQDNCapability) GetHostNameLen() uint32 {
	if m != nil {
		return m.HostNameLen
	}
	return 0
}

func (m *FQDNCapability) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *FQDNCapability) GetDomainNameLen() uint32 {
	if m != nil {
		return m.DomainNameLen
	}
	return 0
}

func (m *FQDNCapability) GetDomainName() string {
	if m != nil {
		return m.DomainName
	}
	return ""
}

type UnknownCapability struct {
	Code  uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *UnknownCapability) Reset()                    { *m = UnknownCapability{} }
func (m *UnknownCapability) String() string            { return proto.CompactTextString(m) }
func (*UnknownCapability) ProtoMessage()               {}
func (*UnknownCapability) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *UnknownCapability) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UnknownCapability) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*MultiProtocolCapability)(nil), "gobgpapi.MultiProtocolCapability")
	proto.RegisterType((*RouteRefreshCapability)(nil), "gobgpapi.RouteRefreshCapability")
	proto.RegisterType((*CarryingLabelInfoCapability)(nil), "gobgpapi.CarryingLabelInfoCapability")
	proto.RegisterType((*ExtendedNexthopCapabilityTuple)(nil), "gobgpapi.ExtendedNexthopCapabilityTuple")
	proto.RegisterType((*ExtendedNexthopCapability)(nil), "gobgpapi.ExtendedNexthopCapability")
	proto.RegisterType((*GracefulRestartCapabilityTuple)(nil), "gobgpapi.GracefulRestartCapabilityTuple")
	proto.RegisterType((*GracefulRestartCapability)(nil), "gobgpapi.GracefulRestartCapability")
	proto.RegisterType((*FourOctetASNumberCapability)(nil), "gobgpapi.FourOctetASNumberCapability")
	proto.RegisterType((*AddPathCapabilityTuple)(nil), "gobgpapi.AddPathCapabilityTuple")
	proto.RegisterType((*AddPathCapability)(nil), "gobgpapi.AddPathCapability")
	proto.RegisterType((*EnhancedRouteRefreshCapability)(nil), "gobgpapi.EnhancedRouteRefreshCapability")
	proto.RegisterType((*LongLivedGracefulRestartCapabilityTuple)(nil), "gobgpapi.LongLivedGracefulRestartCapabilityTuple")
	proto.RegisterType((*LongLivedGracefulRestartCapability)(nil), "gobgpapi.LongLivedGracefulRestartCapability")
	proto.RegisterType((*RouteRefreshCiscoCapability)(nil), "gobgpapi.RouteRefreshCiscoCapability")
	proto.RegisterType((*FQDNCapability)(nil), "gobgpapi.FQDNCapability")
	proto.RegisterType((*UnknownCapability)(nil), "gobgpapi.UnknownCapability")
	proto.RegisterEnum("gobgpapi.AddPathMode", AddPathMode_name, AddPathMode_value)
}

func init() { proto.RegisterFile("capability.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0xa8, 0xda, 0x71, 0x13, 0xdc, 0x15, 0x94, 0x94, 0xa8, 0x25, 0xda, 0x03, 0x04,
	0x24, 0x22, 0x35, 0x1c, 0xe0, 0x82, 0x44, 0x49, 0x1c, 0x88, 0x94, 0x38, 0xc5, 0x2d, 0xdc, 0x50,
	0xd8, 0xd8, 0x9b, 0xc4, 0xc2, 0xde, 0xb5, 0xec, 0x75, 0x69, 0x0e, 0x9c, 0xb9, 0xf0, 0x05, 0x7c,
	0x2d, 0xb2, 0xd7, 0xb1, 0x4d, 0xaa, 0xb4, 0x15, 0x12, 0xb7, 0x9d, 0x9d, 0x37, 0x6f, 0xde, 0x9b,
	0x1d, 0x1b, 0x34, 0x8b, 0xf8, 0x64, 0xea, 0xb8, 0x8e, 0x58, 0xb6, 0xfd, 0x80, 0x0b, 0x8e, 0xb6,
	0xe7, 0x7c, 0x3a, 0xf7, 0x89, 0xef, 0x3c, 0x52, 0x93, 0x93, 0xbc, 0xc6, 0x5d, 0x78, 0x38, 0x8a,
	0x5c, 0xe1, 0x9c, 0xc6, 0x91, 0xc5, 0xdd, 0x6e, 0x56, 0x87, 0x5a, 0xb0, 0x35, 0x23, 0x9e, 0xe3,
	0x2e, 0xeb, 0x4a, 0x53, 0x69, 0xa9, 0x1d, 0xad, 0xbd, 0xa2, 0x68, 0xf7, 0x93, 0x7b, 0x33, 0xcd,
	0xe3, 0x3a, 0xec, 0x9b, 0x3c, 0x12, 0xd4, 0xa4, 0xb3, 0x80, 0x86, 0x8b, 0x9c, 0x03, 0x1f, 0x42,
	0xa3, 0x4b, 0x82, 0x60, 0xe9, 0xb0, 0xf9, 0x90, 0x4c, 0xa9, 0x3b, 0x60, 0x33, 0x5e, 0x48, 0xff,
	0x52, 0xe0, 0x48, 0xbf, 0x14, 0x94, 0xd9, 0xd4, 0x36, 0xe8, 0xa5, 0x58, 0x70, 0x3f, 0xcf, 0x9e,
	0x47, 0xbe, 0x4b, 0xd1, 0x31, 0xa8, 0xcc, 0x0d, 0x9c, 0xc9, 0x0d, 0x52, 0x20, 0x06, 0xc9, 0x33,
	0x7a, 0x05, 0x35, 0x26, 0xc9, 0x56, 0x55, 0xa5, 0x0d, 0x55, 0xd5, 0x14, 0x27, 0x43, 0xfc, 0x05,
	0x0e, 0x36, 0xaa, 0x41, 0x6f, 0x61, 0x4b, 0xc4, 0x8a, 0xc2, 0xba, 0xd2, 0x2c, 0xb7, 0xd4, 0x4e,
	0x2b, 0x67, 0xbb, 0xde, 0x82, 0x99, 0xd6, 0xe1, 0xaf, 0x70, 0xf4, 0x3e, 0x20, 0x16, 0x9d, 0x45,
	0xae, 0x49, 0x43, 0x41, 0x02, 0xb1, 0x6e, 0xf6, 0xd6, 0x23, 0x47, 0xf7, 0xe1, 0xee, 0xcc, 0x25,
	0xf3, 0x30, 0xb1, 0x56, 0x35, 0x65, 0x80, 0x7f, 0x2a, 0x70, 0xb0, 0xb1, 0x45, 0x5e, 0xa3, 0x14,
	0x6a, 0x10, 0x82, 0x8a, 0x70, 0x3c, 0x9a, 0x12, 0x25, 0xe7, 0x82, 0xd7, 0xf2, 0xba, 0xd7, 0xeb,
	0x1d, 0x64, 0x5e, 0x5f, 0x40, 0xa3, 0xcf, 0xa3, 0x60, 0x6c, 0x09, 0x2a, 0x4e, 0xce, 0x8c, 0xc8,
	0x9b, 0xd2, 0xa0, 0x20, 0xa5, 0x06, 0x25, 0xb2, 0xd2, 0x51, 0x22, 0x21, 0xf6, 0x60, 0xff, 0xc4,
	0xb6, 0x4f, 0x89, 0x58, 0xfc, 0xfb, 0x48, 0x9e, 0x41, 0xc5, 0xe3, 0xb6, 0x34, 0x52, 0xeb, 0x3c,
	0xc8, 0x71, 0x29, 0xf3, 0x88, 0xdb, 0xd4, 0x4c, 0x20, 0x78, 0x04, 0x7b, 0x57, 0xda, 0xa1, 0xd7,
	0x6b, 0x0f, 0xdc, 0xbc, 0xc2, 0xb0, 0xc9, 0x6c, 0x13, 0x8e, 0x74, 0xb6, 0x20, 0xcc, 0xa2, 0xf6,
	0x86, 0xef, 0xe0, 0x07, 0x3c, 0x1d, 0x72, 0x36, 0x1f, 0x3a, 0x17, 0xd4, 0xfe, 0xbf, 0x3b, 0x90,
	0xbd, 0x67, 0x39, 0x7f, 0x4f, 0xcc, 0x01, 0xdf, 0xdc, 0x1e, 0x0d, 0xd6, 0x06, 0x70, 0x9c, 0x77,
	0xbe, 0xa5, 0xf8, 0x6c, 0x22, 0x87, 0xd0, 0xf8, 0x6b, 0x12, 0x4e, 0x68, 0x15, 0xbf, 0xfb, 0xdf,
	0x0a, 0xd4, 0xfa, 0x1f, 0x7b, 0x46, 0xa1, 0x39, 0x86, 0xea, 0x82, 0x87, 0x62, 0xc2, 0x88, 0x47,
	0x27, 0x2e, 0x65, 0xe9, 0x72, 0xa8, 0xf1, 0xa5, 0x41, 0x3c, 0x3a, 0xa4, 0x0c, 0x35, 0x60, 0x27,
	0xc3, 0x24, 0xa6, 0x77, 0xcc, 0xed, 0x55, 0x1e, 0x3d, 0x81, 0x7b, 0x36, 0xf7, 0x88, 0xc3, 0x72,
	0x0a, 0x39, 0x82, 0xaa, 0xbc, 0x5e, 0x91, 0x3c, 0x06, 0xb5, 0x80, 0xab, 0x57, 0x12, 0x1a, 0xc8,
	0x31, 0xf8, 0x0d, 0xec, 0x7d, 0x62, 0xdf, 0x18, 0xff, 0xce, 0x0a, 0xf2, 0x10, 0x54, 0xac, 0x78,
	0xb9, 0xa4, 0xaa, 0xe4, 0x1c, 0xcf, 0xff, 0x82, 0xb8, 0x91, 0x94, 0xb2, 0x6b, 0xca, 0xe0, 0xf9,
	0x10, 0xd4, 0xc2, 0xc2, 0xa1, 0x2a, 0xec, 0x8c, 0xc6, 0x3d, 0x7d, 0x62, 0x8c, 0x0d, 0x5d, 0xbb,
	0x83, 0x34, 0xd8, 0x4d, 0x42, 0x53, 0xef, 0xea, 0x83, 0xcf, 0xba, 0xa6, 0x64, 0x80, 0x33, 0xdd,
	0xe8, 0x69, 0xa5, 0x2c, 0x7c, 0x37, 0x3e, 0xff, 0xa0, 0x95, 0xa7, 0x5b, 0xc9, 0x6f, 0xfa, 0xe5,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x77, 0x70, 0x63, 0xd1, 0x05, 0x00, 0x00,
}
