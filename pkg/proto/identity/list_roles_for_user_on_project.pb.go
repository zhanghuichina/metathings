// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_roles_for_user_on_project.proto

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

type ListRolesForUserOnProjectRequest struct {
	ProjectId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	UserId               *wrappers.StringValue `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRolesForUserOnProjectRequest) Reset()         { *m = ListRolesForUserOnProjectRequest{} }
func (m *ListRolesForUserOnProjectRequest) String() string { return proto.CompactTextString(m) }
func (*ListRolesForUserOnProjectRequest) ProtoMessage()    {}
func (*ListRolesForUserOnProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_roles_for_user_on_project_34b9abec53cb2bc8, []int{0}
}
func (m *ListRolesForUserOnProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesForUserOnProjectRequest.Unmarshal(m, b)
}
func (m *ListRolesForUserOnProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesForUserOnProjectRequest.Marshal(b, m, deterministic)
}
func (dst *ListRolesForUserOnProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesForUserOnProjectRequest.Merge(dst, src)
}
func (m *ListRolesForUserOnProjectRequest) XXX_Size() int {
	return xxx_messageInfo_ListRolesForUserOnProjectRequest.Size(m)
}
func (m *ListRolesForUserOnProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesForUserOnProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesForUserOnProjectRequest proto.InternalMessageInfo

func (m *ListRolesForUserOnProjectRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *ListRolesForUserOnProjectRequest) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

type ListRolesForUserOnProjectResponse struct {
	Roles                []*Role  `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRolesForUserOnProjectResponse) Reset()         { *m = ListRolesForUserOnProjectResponse{} }
func (m *ListRolesForUserOnProjectResponse) String() string { return proto.CompactTextString(m) }
func (*ListRolesForUserOnProjectResponse) ProtoMessage()    {}
func (*ListRolesForUserOnProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_roles_for_user_on_project_34b9abec53cb2bc8, []int{1}
}
func (m *ListRolesForUserOnProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesForUserOnProjectResponse.Unmarshal(m, b)
}
func (m *ListRolesForUserOnProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesForUserOnProjectResponse.Marshal(b, m, deterministic)
}
func (dst *ListRolesForUserOnProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesForUserOnProjectResponse.Merge(dst, src)
}
func (m *ListRolesForUserOnProjectResponse) XXX_Size() int {
	return xxx_messageInfo_ListRolesForUserOnProjectResponse.Size(m)
}
func (m *ListRolesForUserOnProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesForUserOnProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesForUserOnProjectResponse proto.InternalMessageInfo

func (m *ListRolesForUserOnProjectResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRolesForUserOnProjectRequest)(nil), "ai.metathings.service.identity.ListRolesForUserOnProjectRequest")
	proto.RegisterType((*ListRolesForUserOnProjectResponse)(nil), "ai.metathings.service.identity.ListRolesForUserOnProjectResponse")
}

func init() {
	proto.RegisterFile("list_roles_for_user_on_project.proto", fileDescriptor_list_roles_for_user_on_project_34b9abec53cb2bc8)
}

var fileDescriptor_list_roles_for_user_on_project_34b9abec53cb2bc8 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xf3, 0x50,
	0x10, 0x85, 0x49, 0x7f, 0xfe, 0x8a, 0xb7, 0xbb, 0xac, 0x4a, 0x91, 0x1a, 0x4b, 0x17, 0xdd, 0xf4,
	0x06, 0x2a, 0xb8, 0x10, 0xdc, 0x28, 0x08, 0x05, 0x41, 0x89, 0xe8, 0xf6, 0x92, 0x36, 0xd3, 0x74,
	0x34, 0xbd, 0x13, 0x67, 0x26, 0x2d, 0xbe, 0x90, 0xaf, 0x25, 0xf8, 0x24, 0x92, 0xa4, 0xba, 0x2c,
	0xb8, 0x3b, 0x30, 0xf3, 0x1d, 0x3e, 0x8e, 0x19, 0x17, 0x28, 0xea, 0x98, 0x0a, 0x10, 0xb7, 0x22,
	0x76, 0x95, 0x00, 0x3b, 0xf2, 0xae, 0x64, 0x7a, 0x81, 0xa5, 0xda, 0x92, 0x49, 0x29, 0x1c, 0xa6,
	0x68, 0x37, 0xa0, 0xa9, 0xae, 0xd1, 0xe7, 0x62, 0x05, 0x78, 0x8b, 0x4b, 0xb0, 0x98, 0x81, 0x57,
	0xd4, 0xf7, 0xc1, 0x30, 0x27, 0xca, 0x0b, 0x88, 0x9b, 0xef, 0x45, 0xb5, 0x8a, 0x77, 0x9c, 0x96,
	0x25, 0xb0, 0xb4, 0xfc, 0xe0, 0x22, 0x47, 0x5d, 0x57, 0x0b, 0xbb, 0xa4, 0x4d, 0xbc, 0xd9, 0xa1,
	0xbe, 0xd2, 0x2e, 0xce, 0x69, 0xda, 0x1c, 0xa7, 0xdb, 0xb4, 0xc0, 0x2c, 0x55, 0x62, 0x89, 0x7f,
	0xe3, 0x9e, 0x33, 0xb5, 0x58, 0x9b, 0x47, 0x1f, 0x81, 0x89, 0xee, 0x50, 0x34, 0xa9, 0x5d, 0x6f,
	0x89, 0x9f, 0x04, 0xf8, 0xde, 0x3f, 0xb4, 0x9e, 0x09, 0xbc, 0x55, 0x20, 0x1a, 0xde, 0x18, 0xb3,
	0x37, 0x77, 0x98, 0xf5, 0x83, 0x28, 0x98, 0xf4, 0x66, 0x27, 0xb6, 0xb5, 0xb3, 0x3f, 0x76, 0xf6,
	0x51, 0x19, 0x7d, 0xfe, 0x9c, 0x16, 0x15, 0x5c, 0x77, 0xbf, 0x3e, 0x4f, 0x3b, 0x51, 0x90, 0x1c,
	0xef, 0xb9, 0x79, 0x16, 0x5e, 0x99, 0xa3, 0x66, 0x06, 0xcc, 0xfa, 0x9d, 0x3f, 0x34, 0x74, 0x6b,
	0x68, 0x9e, 0x8d, 0x9c, 0x39, 0x3b, 0xe0, 0x29, 0x25, 0x79, 0x81, 0xf0, 0xd2, 0xfc, 0x6f, 0x46,
	0xef, 0x07, 0xd1, 0xbf, 0x49, 0x6f, 0x36, 0xb6, 0x87, 0x17, 0xb6, 0x75, 0x5b, 0xd2, 0x22, 0x8b,
	0x6e, 0xa3, 0x71, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x20, 0x6b, 0xe3, 0xbc, 0x01, 0x00,
	0x00,
}
