// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_role_to_group.proto

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

type AddRoleToGroupRequest struct {
	Group                *OpGroup `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Role                 *OpRole  `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRoleToGroupRequest) Reset()         { *m = AddRoleToGroupRequest{} }
func (m *AddRoleToGroupRequest) String() string { return proto.CompactTextString(m) }
func (*AddRoleToGroupRequest) ProtoMessage()    {}
func (*AddRoleToGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_add_role_to_group_81a7d318e457d83e, []int{0}
}
func (m *AddRoleToGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoleToGroupRequest.Unmarshal(m, b)
}
func (m *AddRoleToGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoleToGroupRequest.Marshal(b, m, deterministic)
}
func (dst *AddRoleToGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoleToGroupRequest.Merge(dst, src)
}
func (m *AddRoleToGroupRequest) XXX_Size() int {
	return xxx_messageInfo_AddRoleToGroupRequest.Size(m)
}
func (m *AddRoleToGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoleToGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoleToGroupRequest proto.InternalMessageInfo

func (m *AddRoleToGroupRequest) GetGroup() *OpGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *AddRoleToGroupRequest) GetRole() *OpRole {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRoleToGroupRequest)(nil), "ai.metathings.service.identityd2.AddRoleToGroupRequest")
}

func init() {
	proto.RegisterFile("add_role_to_group.proto", fileDescriptor_add_role_to_group_81a7d318e457d83e)
}

var fileDescriptor_add_role_to_group_81a7d318e457d83e = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x40, 0xd9, 0xa0, 0x29, 0x36, 0xdd, 0x81, 0x18, 0xd2, 0x78, 0x58, 0xc5, 0x22, 0x7b, 0x10,
	0xc1, 0x5e, 0x11, 0x52, 0x0a, 0x87, 0xfd, 0xb1, 0xc9, 0x0c, 0x9b, 0xc1, 0xdd, 0xcc, 0xb9, 0x3b,
	0x97, 0xc3, 0x1f, 0xf2, 0xb7, 0x04, 0xbf, 0x44, 0x6e, 0x0f, 0xae, 0x35, 0xdd, 0x4c, 0xf1, 0xde,
	0xbc, 0xd1, 0xb7, 0x16, 0xa0, 0x89, 0xec, 0xb1, 0x11, 0x6e, 0x5c, 0xe4, 0xae, 0x35, 0x6d, 0x64,
	0xe1, 0xa2, 0xb4, 0x64, 0x02, 0x8a, 0x95, 0x23, 0x9d, 0x5c, 0x32, 0x09, 0xe3, 0x99, 0x0e, 0x68,
	0x08, 0xf0, 0x24, 0x24, 0x5f, 0xb0, 0x5d, 0x3d, 0x39, 0x92, 0x63, 0xb7, 0x37, 0x07, 0x0e, 0x55,
	0xe8, 0x49, 0x3e, 0xb8, 0xaf, 0x1c, 0x6f, 0x32, 0xbe, 0x39, 0x5b, 0x4f, 0x60, 0x85, 0x63, 0xaa,
	0xa6, 0x71, 0x34, 0xaf, 0x16, 0x81, 0x01, 0xfd, 0xb8, 0xdc, 0x7f, 0x2b, 0x7d, 0xf3, 0x0c, 0x50,
	0xb3, 0xc7, 0x77, 0xde, 0x0d, 0xf7, 0x6b, 0xfc, 0xec, 0x30, 0x49, 0xb1, 0xd3, 0xd7, 0xb9, 0x67,
	0xa9, 0x4a, 0xb5, 0x5e, 0x6c, 0x1f, 0xcc, 0x7f, 0x41, 0xe6, 0xad, 0xcd, 0x82, 0x97, 0xf9, 0xef,
	0xcf, 0xdd, 0xac, 0x54, 0xf5, 0xc8, 0x17, 0xaf, 0xfa, 0x6a, 0x78, 0x70, 0x39, 0xcb, 0x9e, 0xf5,
	0x25, 0x9e, 0x21, 0x67, 0xd2, 0x64, 0x7a, 0x3f, 0xcf, 0xbd, 0x8f, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x43, 0xdf, 0x17, 0x28, 0x31, 0x01, 0x00, 0x00,
}