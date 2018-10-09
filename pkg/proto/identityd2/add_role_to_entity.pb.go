// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_role_to_entity.proto

package identityd2

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

type AddRoleToEntityRequest struct {
	Entity               *OpEntity `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
	Role                 *OpRole   `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddRoleToEntityRequest) Reset()         { *m = AddRoleToEntityRequest{} }
func (m *AddRoleToEntityRequest) String() string { return proto.CompactTextString(m) }
func (*AddRoleToEntityRequest) ProtoMessage()    {}
func (*AddRoleToEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_add_role_to_entity_82d932657cecfb14, []int{0}
}
func (m *AddRoleToEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoleToEntityRequest.Unmarshal(m, b)
}
func (m *AddRoleToEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoleToEntityRequest.Marshal(b, m, deterministic)
}
func (dst *AddRoleToEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoleToEntityRequest.Merge(dst, src)
}
func (m *AddRoleToEntityRequest) XXX_Size() int {
	return xxx_messageInfo_AddRoleToEntityRequest.Size(m)
}
func (m *AddRoleToEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoleToEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoleToEntityRequest proto.InternalMessageInfo

func (m *AddRoleToEntityRequest) GetEntity() *OpEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *AddRoleToEntityRequest) GetRole() *OpRole {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRoleToEntityRequest)(nil), "ai.metathings.service.identityd2.AddRoleToEntityRequest")
}

func init() {
	proto.RegisterFile("add_role_to_entity.proto", fileDescriptor_add_role_to_entity_82d932657cecfb14)
}

var fileDescriptor_add_role_to_entity_82d932657cecfb14 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4c, 0x49, 0x89,
	0x2f, 0xca, 0xcf, 0x49, 0x8d, 0x2f, 0xc9, 0x8f, 0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x48, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c, 0xc9, 0xc8,
	0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0xcb, 0x4c, 0x81, 0xa8,
	0x4a, 0x31, 0x92, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf,
	0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0x6b, 0xd7, 0x2d, 0x4b, 0xcc,
	0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21, 0x26, 0x4b, 0x71, 0xe7, 0xe6,
	0xa7, 0xa4, 0xe6, 0x40, 0x38, 0x4a, 0xab, 0x18, 0xb9, 0xc4, 0x1c, 0x53, 0x52, 0x82, 0xf2, 0x73,
	0x52, 0x43, 0xf2, 0x5d, 0xc1, 0x46, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x79, 0x71,
	0xb1, 0x41, 0xec, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xd2, 0xd2, 0x23, 0xe4, 0x24, 0x3d,
	0xff, 0x02, 0x88, 0x11, 0x4e, 0x6c, 0x8f, 0xee, 0xcb, 0x33, 0x29, 0x30, 0x06, 0x41, 0x4d, 0x10,
	0x72, 0xe1, 0x62, 0x01, 0xf9, 0x52, 0x82, 0x09, 0x6c, 0x92, 0x06, 0x31, 0x26, 0x81, 0x9c, 0x04,
	0x37, 0x07, 0xac, 0x3b, 0x89, 0x0d, 0xec, 0x66, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10,
	0x7f, 0xc9, 0xec, 0x36, 0x01, 0x00, 0x00,
}
