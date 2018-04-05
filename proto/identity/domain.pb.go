// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domain.proto

package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Domain struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Enabled     bool   `protobuf:"varint,4,opt,name=enabled" json:"enabled,omitempty"`
}

func (m *Domain) Reset()                    { *m = Domain{} }
func (m *Domain) String() string            { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()               {}
func (*Domain) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

func (m *Domain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Domain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Domain) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Domain) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*Domain)(nil), "ai.metathings.service.identity.Domain")
}

func init() { proto.RegisterFile("domain.proto", fileDescriptor26) }

var fileDescriptor26 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x31, 0x0a, 0xc2, 0x40,
	0x10, 0x45, 0x49, 0x0c, 0x51, 0x47, 0xb1, 0x98, 0x6a, 0x2b, 0x09, 0x56, 0xa9, 0xb6, 0xf1, 0x0a,
	0x9e, 0x20, 0x37, 0xd8, 0x64, 0x06, 0x33, 0xe0, 0xce, 0x86, 0xdd, 0x41, 0xf0, 0xf6, 0xc2, 0x82,
	0x60, 0xf7, 0xff, 0xfb, 0xbf, 0x78, 0x70, 0xa6, 0x14, 0x83, 0xa8, 0xdf, 0x72, 0xb2, 0x84, 0xd7,
	0x20, 0x3e, 0xb2, 0x05, 0x5b, 0x45, 0x9f, 0xc5, 0x17, 0xce, 0x6f, 0x59, 0xd8, 0x0b, 0xb1, 0x9a,
	0xd8, 0xe7, 0xb6, 0x42, 0xff, 0xa8, 0x7f, 0xbc, 0x40, 0x2b, 0xe4, 0x9a, 0xa1, 0x19, 0x8f, 0x53,
	0x2b, 0x84, 0x08, 0x9d, 0x86, 0xc8, 0xae, 0xad, 0xa4, 0x66, 0x1c, 0xe0, 0x44, 0x5c, 0x96, 0x2c,
	0x9b, 0x49, 0x52, 0xb7, 0xab, 0xd3, 0x3f, 0x42, 0x07, 0x7b, 0xd6, 0x30, 0xbf, 0x98, 0x5c, 0x37,
	0x34, 0xe3, 0x61, 0xfa, 0xd5, 0xb9, 0xaf, 0x42, 0xf7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63,
	0x5d, 0x4a, 0xc9, 0xa0, 0x00, 0x00, 0x00,
}
