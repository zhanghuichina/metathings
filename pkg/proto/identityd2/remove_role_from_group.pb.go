// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remove_role_from_group.proto

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

type RemoveRoleFromGroupRequest struct {
	Group                *OpGroup `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Role                 *OpRole  `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRoleFromGroupRequest) Reset()         { *m = RemoveRoleFromGroupRequest{} }
func (m *RemoveRoleFromGroupRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRoleFromGroupRequest) ProtoMessage()    {}
func (*RemoveRoleFromGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_remove_role_from_group_16bb161b8e47b1d3, []int{0}
}
func (m *RemoveRoleFromGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRoleFromGroupRequest.Unmarshal(m, b)
}
func (m *RemoveRoleFromGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRoleFromGroupRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveRoleFromGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRoleFromGroupRequest.Merge(dst, src)
}
func (m *RemoveRoleFromGroupRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRoleFromGroupRequest.Size(m)
}
func (m *RemoveRoleFromGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRoleFromGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRoleFromGroupRequest proto.InternalMessageInfo

func (m *RemoveRoleFromGroupRequest) GetGroup() *OpGroup {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *RemoveRoleFromGroupRequest) GetRole() *OpRole {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoveRoleFromGroupRequest)(nil), "ai.metathings.service.identityd2.RemoveRoleFromGroupRequest")
}

func init() {
	proto.RegisterFile("remove_role_from_group.proto", fileDescriptor_remove_role_from_group_16bb161b8e47b1d3)
}

var fileDescriptor_remove_role_from_group_16bb161b8e47b1d3 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcf, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x71, 0x52, 0xb4, 0x87, 0xed, 0x2d, 0xa7, 0x12, 0x04, 0x83, 0xa7, 0x7a, 0xe8, 0x06,
	0x2a, 0xf8, 0x00, 0x22, 0xf6, 0x28, 0xec, 0x0b, 0x84, 0x6d, 0x77, 0xdc, 0x0e, 0xee, 0x76, 0xe2,
	0xec, 0x24, 0xc5, 0x67, 0xf2, 0xa1, 0x04, 0x9f, 0x44, 0xb2, 0x81, 0x5c, 0xf5, 0xb6, 0x7b, 0xf8,
	0xfe, 0xfc, 0x46, 0xdd, 0x30, 0x44, 0x1a, 0xa0, 0x65, 0x0a, 0xd0, 0xbe, 0x31, 0xc5, 0xd6, 0x33,
	0xf5, 0x9d, 0xee, 0x98, 0x84, 0xca, 0xda, 0xa2, 0x8e, 0x20, 0x56, 0x4e, 0x78, 0xf6, 0x49, 0x27,
	0xe0, 0x01, 0x8f, 0xa0, 0xd1, 0xc1, 0x59, 0x50, 0x3e, 0xdd, 0xae, 0x7a, 0xf4, 0x28, 0xa7, 0xfe,
	0xa0, 0x8f, 0x14, 0x9b, 0x78, 0x41, 0x79, 0xa7, 0x4b, 0xe3, 0x69, 0x9b, 0xe7, 0xdb, 0xc1, 0x06,
	0x74, 0x56, 0x88, 0x53, 0x33, 0x3f, 0xa7, 0x72, 0xb5, 0x8a, 0xe4, 0x20, 0x4c, 0x9f, 0xbb, 0xaf,
	0x42, 0x55, 0x26, 0x3b, 0x0c, 0x05, 0x78, 0x61, 0x8a, 0xfb, 0x11, 0x61, 0xe0, 0xa3, 0x87, 0x24,
	0xe5, 0x5e, 0x5d, 0x67, 0xd4, 0xba, 0xa8, 0x8b, 0xcd, 0x6a, 0x77, 0xaf, 0xff, 0x52, 0xe9, 0xd7,
	0x2e, 0x07, 0x9e, 0x96, 0x3f, 0xdf, 0xb7, 0x8b, 0xba, 0x30, 0xd3, 0xbe, 0x7c, 0x56, 0x57, 0xe3,
	0x9d, 0xeb, 0x45, 0xee, 0x6c, 0xfe, 0xd3, 0x19, 0x41, 0x73, 0x26, 0xaf, 0x0f, 0xcb, 0x8c, 0x7e,
	0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xb8, 0xb3, 0x9b, 0x3b, 0x01, 0x00, 0x00,
}