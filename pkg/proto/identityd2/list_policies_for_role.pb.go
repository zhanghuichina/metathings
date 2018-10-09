// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_policies_for_role.proto

package identityd2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListPoliciesForRoleRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListPoliciesForRoleRequest) Reset()         { *m = ListPoliciesForRoleRequest{} }
func (m *ListPoliciesForRoleRequest) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesForRoleRequest) ProtoMessage()    {}
func (*ListPoliciesForRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_policies_for_role_032cf43aface0f21, []int{0}
}
func (m *ListPoliciesForRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesForRoleRequest.Unmarshal(m, b)
}
func (m *ListPoliciesForRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesForRoleRequest.Marshal(b, m, deterministic)
}
func (dst *ListPoliciesForRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesForRoleRequest.Merge(dst, src)
}
func (m *ListPoliciesForRoleRequest) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesForRoleRequest.Size(m)
}
func (m *ListPoliciesForRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesForRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesForRoleRequest proto.InternalMessageInfo

func (m *ListPoliciesForRoleRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type ListPoliciesForRoleResponse struct {
	Policies             []*Policy `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListPoliciesForRoleResponse) Reset()         { *m = ListPoliciesForRoleResponse{} }
func (m *ListPoliciesForRoleResponse) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesForRoleResponse) ProtoMessage()    {}
func (*ListPoliciesForRoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_policies_for_role_032cf43aface0f21, []int{1}
}
func (m *ListPoliciesForRoleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesForRoleResponse.Unmarshal(m, b)
}
func (m *ListPoliciesForRoleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesForRoleResponse.Marshal(b, m, deterministic)
}
func (dst *ListPoliciesForRoleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesForRoleResponse.Merge(dst, src)
}
func (m *ListPoliciesForRoleResponse) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesForRoleResponse.Size(m)
}
func (m *ListPoliciesForRoleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesForRoleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesForRoleResponse proto.InternalMessageInfo

func (m *ListPoliciesForRoleResponse) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func init() {
	proto.RegisterType((*ListPoliciesForRoleRequest)(nil), "ai.metathings.service.identityd2.ListPoliciesForRoleRequest")
	proto.RegisterType((*ListPoliciesForRoleResponse)(nil), "ai.metathings.service.identityd2.ListPoliciesForRoleResponse")
}

func init() {
	proto.RegisterFile("list_policies_for_role.proto", fileDescriptor_list_policies_for_role_032cf43aface0f21)
}

var fileDescriptor_list_policies_for_role_032cf43aface0f21 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0x49, 0xbf, 0xf2, 0x21, 0x9b, 0x5b, 0x4e, 0x25, 0x16, 0x09, 0x3d, 0xe5, 0x20, 0x5b,
	0x88, 0x7f, 0x41, 0x3c, 0x79, 0x90, 0x15, 0xbc, 0x86, 0x74, 0x77, 0x1a, 0x07, 0xb6, 0x3b, 0xeb,
	0xcc, 0x44, 0xe9, 0xbf, 0x17, 0x13, 0xeb, 0x49, 0xf0, 0x38, 0x30, 0xef, 0xf3, 0xbe, 0x8f, 0xd9,
	0x46, 0x14, 0xed, 0x33, 0x45, 0xf4, 0x08, 0xd2, 0x1f, 0x89, 0x7b, 0xa6, 0x08, 0x36, 0x33, 0x29,
	0x55, 0xcd, 0x80, 0xf6, 0x04, 0x3a, 0xe8, 0x2b, 0xa6, 0x51, 0xac, 0x00, 0xbf, 0xa3, 0x07, 0x8b,
	0x01, 0x92, 0xa2, 0x9e, 0x43, 0x57, 0xdf, 0x8c, 0x44, 0x63, 0x84, 0xfd, 0xfc, 0x7f, 0x98, 0x8e,
	0xfb, 0x0f, 0x1e, 0x72, 0x06, 0x96, 0x85, 0x50, 0x97, 0x27, 0x0a, 0x10, 0x97, 0x63, 0xe7, 0x4d,
	0xfd, 0x88, 0xa2, 0x4f, 0xdf, 0x6d, 0x0f, 0xc4, 0x8e, 0x22, 0x38, 0x78, 0x9b, 0x40, 0xb4, 0xba,
	0x35, 0x2b, 0x0c, 0x9b, 0xa2, 0x29, 0xda, 0xb2, 0xdb, 0xda, 0x85, 0x6b, 0x2f, 0x5c, 0xfb, 0xac,
	0x8c, 0x69, 0x7c, 0x19, 0xe2, 0x04, 0x6e, 0x85, 0xc1, 0xad, 0x79, 0x8a, 0xe0, 0xca, 0x00, 0xe2,
	0x19, 0xb3, 0x22, 0x25, 0xb7, 0xfe, 0x5a, 0xbe, 0xf3, 0xe6, 0xfa, 0xd7, 0x12, 0xc9, 0x94, 0x04,
	0xaa, 0x7b, 0x73, 0x75, 0xb1, 0xdd, 0x14, 0xcd, 0xbf, 0xb6, 0xec, 0x5a, 0xfb, 0x97, 0xa5, 0x9d,
	0x61, 0x67, 0xf7, 0x93, 0x3c, 0xfc, 0x9f, 0x77, 0xdd, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0xdb,
	0x67, 0x17, 0xe4, 0x3f, 0x01, 0x00, 0x00,
}
