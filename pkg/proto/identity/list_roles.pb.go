// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_roles.proto

package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ListRolesRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DomainId             *wrappers.StringValue `protobuf:"bytes,2,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRolesRequest) Reset()         { *m = ListRolesRequest{} }
func (m *ListRolesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRolesRequest) ProtoMessage()    {}
func (*ListRolesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_roles_1cce66fc26122fd4, []int{0}
}
func (m *ListRolesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesRequest.Unmarshal(m, b)
}
func (m *ListRolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesRequest.Marshal(b, m, deterministic)
}
func (dst *ListRolesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesRequest.Merge(dst, src)
}
func (m *ListRolesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRolesRequest.Size(m)
}
func (m *ListRolesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesRequest proto.InternalMessageInfo

func (m *ListRolesRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListRolesRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

type ListRolesResponse struct {
	Roles                []*Role  `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRolesResponse) Reset()         { *m = ListRolesResponse{} }
func (m *ListRolesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRolesResponse) ProtoMessage()    {}
func (*ListRolesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_roles_1cce66fc26122fd4, []int{1}
}
func (m *ListRolesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesResponse.Unmarshal(m, b)
}
func (m *ListRolesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesResponse.Marshal(b, m, deterministic)
}
func (dst *ListRolesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesResponse.Merge(dst, src)
}
func (m *ListRolesResponse) XXX_Size() int {
	return xxx_messageInfo_ListRolesResponse.Size(m)
}
func (m *ListRolesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesResponse proto.InternalMessageInfo

func (m *ListRolesResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRolesRequest)(nil), "ai.metathings.service.identity.ListRolesRequest")
	proto.RegisterType((*ListRolesResponse)(nil), "ai.metathings.service.identity.ListRolesResponse")
}

func init() { proto.RegisterFile("list_roles.proto", fileDescriptor_list_roles_1cce66fc26122fd4) }

var fileDescriptor_list_roles_1cce66fc26122fd4 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0xff, 0xd0, 0xed, 0xa5, 0xe6, 0x14, 0x8a, 0x94, 0x52, 0x3c, 0xf4, 0xd2, 0x89,
	0x54, 0x10, 0xf4, 0x1b, 0x08, 0x82, 0x10, 0xc1, 0x6b, 0xd9, 0x34, 0xe3, 0x76, 0x30, 0xbb, 0x13,
	0x77, 0x27, 0x0d, 0x9e, 0xfc, 0xea, 0x92, 0x4d, 0x15, 0x4f, 0xd2, 0xdb, 0x83, 0x79, 0x3f, 0xe6,
	0xc7, 0x53, 0xe3, 0x9a, 0x82, 0xac, 0x3d, 0xd7, 0x18, 0xa0, 0xf1, 0x2c, 0x9c, 0x4e, 0x35, 0x81,
	0x45, 0xd1, 0xb2, 0x25, 0x67, 0x02, 0x04, 0xf4, 0x3b, 0xda, 0x20, 0x50, 0x85, 0x4e, 0x48, 0x3e,
	0x27, 0x53, 0xc3, 0x6c, 0x6a, 0xcc, 0x63, 0xbb, 0x6c, 0xdf, 0xf2, 0xce, 0xeb, 0xa6, 0x41, 0xbf,
	0xe7, 0x27, 0x77, 0x86, 0x64, 0xdb, 0x96, 0xb0, 0x61, 0x9b, 0xdb, 0x8e, 0xe4, 0x9d, 0xbb, 0xdc,
	0xf0, 0x32, 0x1e, 0x97, 0x3b, 0x5d, 0x53, 0xa5, 0x85, 0x7d, 0xc8, 0x7f, 0xe3, 0x9e, 0x53, 0xbd,
	0xc4, 0x90, 0xe7, 0x5f, 0x6a, 0xfc, 0x44, 0x41, 0x8a, 0x5e, 0xab, 0xc0, 0x8f, 0x16, 0x83, 0xa4,
	0x37, 0xea, 0xc4, 0x69, 0x8b, 0x59, 0x32, 0x4b, 0x16, 0xa3, 0xd5, 0x15, 0x0c, 0x1a, 0xf0, 0xa3,
	0x01, 0x2f, 0xe2, 0xc9, 0x99, 0x57, 0x5d, 0xb7, 0x58, 0xc4, 0x66, 0x7a, 0xaf, 0x2e, 0x2a, 0xb6,
	0x9a, 0xdc, 0x9a, 0xaa, 0xec, 0xe8, 0x00, 0xec, 0x7c, 0xa8, 0x3f, 0x56, 0xf3, 0x67, 0x75, 0xf9,
	0x47, 0x20, 0x34, 0xec, 0x02, 0xa6, 0x0f, 0xea, 0x34, 0x0e, 0x95, 0x25, 0xb3, 0xe3, 0xc5, 0x68,
	0x75, 0x0d, 0xff, 0x2f, 0x05, 0x3d, 0x5d, 0x0c, 0x48, 0x79, 0x16, 0x1f, 0xde, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x16, 0x6e, 0x5a, 0x4e, 0x70, 0x01, 0x00, 0x00,
}
