// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role.proto

package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Role struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	DomainId string `protobuf:"bytes,3,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor64, []int{0} }

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func init() {
	proto.RegisterType((*Role)(nil), "ai.metathings.service.identity.Role")
}

func init() { proto.RegisterFile("role.proto", fileDescriptor64) }

var fileDescriptor64 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xca, 0xcf, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c,
	0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c,
	0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x54, 0x72, 0xe7, 0x62, 0x09, 0xca, 0xcf, 0x49, 0x15, 0xe2,
	0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12,
	0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0xd2, 0x5c, 0x9c,
	0x29, 0xf9, 0xb9, 0x89, 0x99, 0x79, 0xf1, 0x99, 0x29, 0x12, 0xcc, 0x60, 0x09, 0x0e, 0x88, 0x80,
	0x67, 0x4a, 0x12, 0x1b, 0xd8, 0x3e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x65, 0xea,
	0x15, 0x7d, 0x00, 0x00, 0x00,
}
