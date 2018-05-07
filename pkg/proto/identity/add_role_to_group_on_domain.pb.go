// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add_role_to_group_on_domain.proto

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

type AddRoleToGroupOnDomainRequest struct {
	DomainId             *wrappers.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	GroupId              *wrappers.StringValue `protobuf:"bytes,2,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	RoleId               *wrappers.StringValue `protobuf:"bytes,3,opt,name=role_id,json=roleId" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AddRoleToGroupOnDomainRequest) Reset()         { *m = AddRoleToGroupOnDomainRequest{} }
func (m *AddRoleToGroupOnDomainRequest) String() string { return proto.CompactTextString(m) }
func (*AddRoleToGroupOnDomainRequest) ProtoMessage()    {}
func (*AddRoleToGroupOnDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_add_role_to_group_on_domain_876586b8d9117112, []int{0}
}
func (m *AddRoleToGroupOnDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRoleToGroupOnDomainRequest.Unmarshal(m, b)
}
func (m *AddRoleToGroupOnDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRoleToGroupOnDomainRequest.Marshal(b, m, deterministic)
}
func (dst *AddRoleToGroupOnDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRoleToGroupOnDomainRequest.Merge(dst, src)
}
func (m *AddRoleToGroupOnDomainRequest) XXX_Size() int {
	return xxx_messageInfo_AddRoleToGroupOnDomainRequest.Size(m)
}
func (m *AddRoleToGroupOnDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRoleToGroupOnDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRoleToGroupOnDomainRequest proto.InternalMessageInfo

func (m *AddRoleToGroupOnDomainRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *AddRoleToGroupOnDomainRequest) GetGroupId() *wrappers.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *AddRoleToGroupOnDomainRequest) GetRoleId() *wrappers.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRoleToGroupOnDomainRequest)(nil), "ai.metathings.service.identity.AddRoleToGroupOnDomainRequest")
}

func init() {
	proto.RegisterFile("add_role_to_group_on_domain.proto", fileDescriptor_add_role_to_group_on_domain_876586b8d9117112)
}

var fileDescriptor_add_role_to_group_on_domain_876586b8d9117112 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0x49, 0x3f, 0x68, 0xfb, 0xc5, 0x5b, 0x4e, 0xa5, 0x68, 0xad, 0x9e, 0xbc, 0x74, 0x03,
	0x0a, 0xde, 0x44, 0x2a, 0x82, 0xe4, 0x24, 0x44, 0xf1, 0xba, 0x6c, 0x3a, 0xe3, 0x76, 0x30, 0xd9,
	0x89, 0x9b, 0x49, 0x83, 0xbf, 0x56, 0xe8, 0x2f, 0x91, 0xee, 0xa2, 0xf7, 0xde, 0x06, 0xe6, 0x7d,
	0x86, 0x67, 0xde, 0xf4, 0xc2, 0x00, 0x68, 0xcf, 0x35, 0x6a, 0x61, 0x6d, 0x3d, 0xf7, 0xad, 0x66,
	0xa7, 0x81, 0x1b, 0x43, 0x4e, 0xb5, 0x9e, 0x85, 0xb3, 0x85, 0x21, 0xd5, 0xa0, 0x18, 0xd9, 0x92,
	0xb3, 0x9d, 0xea, 0xd0, 0xef, 0x68, 0x83, 0x8a, 0x00, 0x9d, 0x90, 0x7c, 0xcd, 0x17, 0x96, 0xd9,
	0xd6, 0x98, 0x87, 0x74, 0xd5, 0xbf, 0xe7, 0x83, 0x37, 0x6d, 0x8b, 0xbe, 0x8b, 0xfc, 0xfc, 0xd6,
	0x92, 0x6c, 0xfb, 0x4a, 0x6d, 0xb8, 0xc9, 0x9b, 0x81, 0xe4, 0x83, 0x87, 0xdc, 0xf2, 0x2a, 0x2c,
	0x57, 0x3b, 0x53, 0x13, 0x18, 0x61, 0xdf, 0xe5, 0x7f, 0x63, 0xe4, 0x2e, 0xf7, 0x49, 0x7a, 0xb6,
	0x06, 0x28, 0xb9, 0xc6, 0x57, 0x7e, 0x3a, 0xa8, 0x3d, 0xbb, 0xc7, 0x20, 0x56, 0xe2, 0x67, 0x8f,
	0x9d, 0x64, 0xeb, 0xf4, 0x7f, 0x34, 0xd5, 0x04, 0xb3, 0x64, 0x99, 0x5c, 0x9d, 0x5c, 0x9f, 0xaa,
	0x68, 0xa3, 0x7e, 0x6d, 0xd4, 0x8b, 0x78, 0x72, 0xf6, 0xcd, 0xd4, 0x3d, 0x3e, 0x8c, 0xf7, 0xdf,
	0xe7, 0xa3, 0x65, 0x52, 0x4e, 0x23, 0x56, 0x40, 0x76, 0x9f, 0x4e, 0xe3, 0xd7, 0x04, 0xb3, 0xd1,
	0x11, 0x17, 0x26, 0x81, 0x2a, 0x20, 0xbb, 0x4b, 0x27, 0xa1, 0x3e, 0x82, 0xd9, 0xbf, 0x23, 0xf8,
	0xf1, 0x01, 0x2a, 0xa0, 0x1a, 0x87, 0xd4, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x4f,
	0xa2, 0xe5, 0x88, 0x01, 0x00, 0x00,
}
