// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_role_to_user_on_domain.proto

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

type AddRoleToUserOnDomainRequest struct {
	DomainId             *wrappers.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	UserId               *wrappers.StringValue `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	RoleId               *wrappers.StringValue `protobuf:"bytes,3,opt,name=role_id,json=roleId" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AddRoleToUserOnDomainRequest) Reset()         { *m = AddRoleToUserOnDomainRequest{} }
func (m *AddRoleToUserOnDomainRequest) String() string { return proto.CompactTextString(m) }
func (*AddRoleToUserOnDomainRequest) ProtoMessage()    {}
func (*AddRoleToUserOnDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_add_role_to_user_on_domain_eea2179082ac5fdd, []int{0}
}
func (m *AddRoleToUserOnDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoleToUserOnDomainRequest.Unmarshal(m, b)
}
func (m *AddRoleToUserOnDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoleToUserOnDomainRequest.Marshal(b, m, deterministic)
}
func (dst *AddRoleToUserOnDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoleToUserOnDomainRequest.Merge(dst, src)
}
func (m *AddRoleToUserOnDomainRequest) XXX_Size() int {
	return xxx_messageInfo_AddRoleToUserOnDomainRequest.Size(m)
}
func (m *AddRoleToUserOnDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoleToUserOnDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoleToUserOnDomainRequest proto.InternalMessageInfo

func (m *AddRoleToUserOnDomainRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *AddRoleToUserOnDomainRequest) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *AddRoleToUserOnDomainRequest) GetRoleId() *wrappers.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRoleToUserOnDomainRequest)(nil), "ai.metathings.service.identity.AddRoleToUserOnDomainRequest")
}

func init() {
	proto.RegisterFile("add_role_to_user_on_domain.proto", fileDescriptor_add_role_to_user_on_domain_eea2179082ac5fdd)
}

var fileDescriptor_add_role_to_user_on_domain_eea2179082ac5fdd = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0x3d, 0x6b, 0xf3, 0x30,
	0x10, 0xc0, 0x71, 0x9c, 0x07, 0xf2, 0xb4, 0xee, 0x96, 0xc9, 0x84, 0x90, 0x9a, 0x4e, 0x5d, 0x22,
	0x43, 0x0b, 0xdd, 0x3a, 0xa4, 0x74, 0xc9, 0x54, 0x70, 0x5f, 0x56, 0x21, 0xe7, 0xae, 0xca, 0x51,
	0x5b, 0xe7, 0x4a, 0xe7, 0x98, 0x7e, 0xda, 0x40, 0x3f, 0x49, 0x89, 0x44, 0xbb, 0x67, 0x13, 0xe8,
	0x7e, 0xe2, 0xaf, 0xcb, 0x4b, 0x03, 0xa0, 0x3d, 0xb7, 0xa8, 0x85, 0xf5, 0x10, 0xd0, 0x6b, 0x76,
	0x1a, 0xb8, 0x33, 0xe4, 0x54, 0xef, 0x59, 0x78, 0xb6, 0x34, 0xa4, 0x3a, 0x14, 0x23, 0x3b, 0x72,
	0x36, 0xa8, 0x80, 0x7e, 0x4f, 0x5b, 0x54, 0x04, 0xe8, 0x84, 0xe4, 0x6b, 0xbe, 0xb4, 0xcc, 0xb6,
	0xc5, 0x2a, 0x4e, 0x37, 0xc3, 0x7b, 0x35, 0x7a, 0xd3, 0xf7, 0xe8, 0x43, 0xf2, 0xf3, 0x3b, 0x4b,
	0xb2, 0x1b, 0x1a, 0xb5, 0xe5, 0xae, 0xea, 0x46, 0x92, 0x0f, 0x1e, 0x2b, 0xcb, 0xab, 0x78, 0xb9,
	0xda, 0x9b, 0x96, 0xc0, 0x08, 0xfb, 0x50, 0xfd, 0x1d, 0x93, 0xbb, 0x3a, 0x64, 0xf9, 0x62, 0x0d,
	0x50, 0x73, 0x8b, 0x2f, 0xfc, 0x1a, 0xd0, 0x3f, 0xb9, 0xc7, 0xd8, 0x55, 0xe3, 0xe7, 0x80, 0x41,
	0x66, 0xeb, 0xfc, 0x3c, 0x85, 0x6a, 0x82, 0x22, 0x2b, 0xb3, 0xeb, 0x8b, 0x9b, 0x85, 0x4a, 0x31,
	0xea, 0x37, 0x46, 0x3d, 0x8b, 0x27, 0x67, 0xdf, 0x4c, 0x3b, 0xe0, 0xc3, 0xf4, 0xfb, 0x70, 0x39,
	0x29, 0xb3, 0xfa, 0x2c, 0xb1, 0x0d, 0xcc, 0xee, 0xf3, 0xff, 0xf1, 0xcf, 0x04, 0xc5, 0xe4, 0x84,
	0x07, 0xa6, 0x47, 0x94, 0x78, 0x5c, 0x1d, 0x41, 0xf1, 0xef, 0x14, 0x7e, 0x44, 0x1b, 0x68, 0xa6,
	0x71, 0xea, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x34, 0x57, 0xeb, 0x72, 0x84, 0x01, 0x00, 0x00,
}
