// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role.proto

package identityd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Role struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	DomainId             string   `protobuf:"bytes,3,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_ce2085af35c6ac6e, []int{0}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

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
	proto.RegisterType((*Role)(nil), "ai.metathings.service.identityd.Role")
}

func init() { proto.RegisterFile("role.proto", fileDescriptor_role_ce2085af35c6ac6e) }

var fileDescriptor_role_ce2085af35c6ac6e = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xca, 0xcf, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4f, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c,
	0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c,
	0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x4c, 0x51, 0x72, 0xe7, 0x62, 0x09, 0xca, 0xcf, 0x49, 0x15,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11,
	0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0xd2, 0x5c,
	0x9c, 0x29, 0xf9, 0xb9, 0x89, 0x99, 0x79, 0xf1, 0x99, 0x29, 0x12, 0xcc, 0x60, 0x09, 0x0e, 0x88,
	0x80, 0x67, 0x4a, 0x12, 0x1b, 0xd8, 0x42, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0xbb,
	0x7f, 0x44, 0x7e, 0x00, 0x00, 0x00,
}