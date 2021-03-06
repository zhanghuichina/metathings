// Code generated by protoc-gen-go. DO NOT EDIT.
// source: check_role_in_user_on_project.proto

package identityd

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

type CheckRoleInUserOnProjectRequest struct {
	ProjectId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	UserId               *wrappers.StringValue `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	RoleId               *wrappers.StringValue `protobuf:"bytes,3,opt,name=role_id,json=roleId" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CheckRoleInUserOnProjectRequest) Reset()         { *m = CheckRoleInUserOnProjectRequest{} }
func (m *CheckRoleInUserOnProjectRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRoleInUserOnProjectRequest) ProtoMessage()    {}
func (*CheckRoleInUserOnProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_check_role_in_user_on_project_3ae05ad25ebc6bc6, []int{0}
}
func (m *CheckRoleInUserOnProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRoleInUserOnProjectRequest.Unmarshal(m, b)
}
func (m *CheckRoleInUserOnProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRoleInUserOnProjectRequest.Marshal(b, m, deterministic)
}
func (dst *CheckRoleInUserOnProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRoleInUserOnProjectRequest.Merge(dst, src)
}
func (m *CheckRoleInUserOnProjectRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRoleInUserOnProjectRequest.Size(m)
}
func (m *CheckRoleInUserOnProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRoleInUserOnProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRoleInUserOnProjectRequest proto.InternalMessageInfo

func (m *CheckRoleInUserOnProjectRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *CheckRoleInUserOnProjectRequest) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *CheckRoleInUserOnProjectRequest) GetRoleId() *wrappers.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckRoleInUserOnProjectRequest)(nil), "ai.metathings.service.identityd.CheckRoleInUserOnProjectRequest")
}

func init() {
	proto.RegisterFile("check_role_in_user_on_project.proto", fileDescriptor_check_role_in_user_on_project_3ae05ad25ebc6bc6)
}

var fileDescriptor_check_role_in_user_on_project_3ae05ad25ebc6bc6 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x4b, 0xf4, 0x40,
	0x10, 0x86, 0xc9, 0x7d, 0x70, 0x1f, 0xc6, 0x2e, 0x55, 0x38, 0xc4, 0x1c, 0xda, 0xd8, 0xdc, 0x06,
	0x14, 0xec, 0x6c, 0xbc, 0x2a, 0x95, 0x12, 0xd1, 0x76, 0xd9, 0x64, 0xc7, 0xcd, 0x78, 0xc9, 0x4e,
	0xdc, 0x9d, 0x5c, 0xf0, 0xd7, 0x0a, 0xe2, 0x0f, 0x91, 0x64, 0x4f, 0xfb, 0xeb, 0x06, 0xe6, 0x7d,
	0x5e, 0x9e, 0x99, 0xf8, 0xb2, 0x6e, 0xa0, 0xde, 0x49, 0x47, 0x2d, 0x48, 0xb4, 0x72, 0xf0, 0xe0,
	0x24, 0x59, 0xd9, 0x3b, 0x7a, 0x83, 0x9a, 0x45, 0xef, 0x88, 0x29, 0xc9, 0x14, 0x8a, 0x0e, 0x58,
	0x71, 0x83, 0xd6, 0x78, 0xe1, 0xc1, 0xed, 0xb1, 0x06, 0x81, 0x1a, 0x2c, 0x23, 0x7f, 0xe8, 0xd5,
	0xb9, 0x21, 0x32, 0x2d, 0xe4, 0x73, 0xbc, 0x1a, 0x5e, 0xf3, 0xd1, 0xa9, 0xbe, 0x07, 0xe7, 0x43,
	0xc1, 0xea, 0xd6, 0x20, 0x37, 0x43, 0x25, 0x6a, 0xea, 0xf2, 0x6e, 0x44, 0xde, 0xd1, 0x98, 0x1b,
	0xda, 0xcc, 0xcb, 0xcd, 0x5e, 0xb5, 0xa8, 0x15, 0x93, 0xf3, 0xf9, 0xdf, 0x18, 0xb8, 0x8b, 0xef,
	0x28, 0xce, 0xb6, 0x93, 0x60, 0x49, 0x2d, 0x14, 0xf6, 0xd9, 0x83, 0x7b, 0xb0, 0x8f, 0xc1, 0xad,
	0x84, 0xf7, 0x01, 0x3c, 0x27, 0xdb, 0x38, 0x3e, 0xd8, 0x4a, 0xd4, 0x69, 0xb4, 0x8e, 0xae, 0x4e,
	0xaf, 0xcf, 0x44, 0x10, 0x12, 0xbf, 0x42, 0xe2, 0x89, 0x1d, 0x5a, 0xf3, 0xa2, 0xda, 0x01, 0xee,
	0x97, 0x5f, 0x9f, 0xd9, 0x62, 0x1d, 0x95, 0x27, 0x07, 0xae, 0xd0, 0xc9, 0x5d, 0xfc, 0x7f, 0x3e,
	0x1d, 0x75, 0xba, 0x38, 0xa2, 0x61, 0x39, 0x41, 0x01, 0x0f, 0x1f, 0xd4, 0xe9, 0xbf, 0x63, 0xf0,
	0x09, 0x2a, 0x74, 0xb5, 0x9c, 0x53, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0x47, 0xb7,
	0x7c, 0x8d, 0x01, 0x00, 0x00,
}
