// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_roles_for_group_on_project.proto

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

type ListRolesForGroupOnProjectRequest struct {
	ProjectId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	GroupId              *wrappers.StringValue `protobuf:"bytes,2,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRolesForGroupOnProjectRequest) Reset()         { *m = ListRolesForGroupOnProjectRequest{} }
func (m *ListRolesForGroupOnProjectRequest) String() string { return proto.CompactTextString(m) }
func (*ListRolesForGroupOnProjectRequest) ProtoMessage()    {}
func (*ListRolesForGroupOnProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_roles_for_group_on_project_ed686b9997180ad8, []int{0}
}
func (m *ListRolesForGroupOnProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesForGroupOnProjectRequest.Unmarshal(m, b)
}
func (m *ListRolesForGroupOnProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesForGroupOnProjectRequest.Marshal(b, m, deterministic)
}
func (dst *ListRolesForGroupOnProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesForGroupOnProjectRequest.Merge(dst, src)
}
func (m *ListRolesForGroupOnProjectRequest) XXX_Size() int {
	return xxx_messageInfo_ListRolesForGroupOnProjectRequest.Size(m)
}
func (m *ListRolesForGroupOnProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesForGroupOnProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesForGroupOnProjectRequest proto.InternalMessageInfo

func (m *ListRolesForGroupOnProjectRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *ListRolesForGroupOnProjectRequest) GetGroupId() *wrappers.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

type ListRolesForGroupOnProjectResponse struct {
	Roles                []*Role  `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRolesForGroupOnProjectResponse) Reset()         { *m = ListRolesForGroupOnProjectResponse{} }
func (m *ListRolesForGroupOnProjectResponse) String() string { return proto.CompactTextString(m) }
func (*ListRolesForGroupOnProjectResponse) ProtoMessage()    {}
func (*ListRolesForGroupOnProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_roles_for_group_on_project_ed686b9997180ad8, []int{1}
}
func (m *ListRolesForGroupOnProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesForGroupOnProjectResponse.Unmarshal(m, b)
}
func (m *ListRolesForGroupOnProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesForGroupOnProjectResponse.Marshal(b, m, deterministic)
}
func (dst *ListRolesForGroupOnProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesForGroupOnProjectResponse.Merge(dst, src)
}
func (m *ListRolesForGroupOnProjectResponse) XXX_Size() int {
	return xxx_messageInfo_ListRolesForGroupOnProjectResponse.Size(m)
}
func (m *ListRolesForGroupOnProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesForGroupOnProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesForGroupOnProjectResponse proto.InternalMessageInfo

func (m *ListRolesForGroupOnProjectResponse) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRolesForGroupOnProjectRequest)(nil), "ai.metathings.service.identityd.ListRolesForGroupOnProjectRequest")
	proto.RegisterType((*ListRolesForGroupOnProjectResponse)(nil), "ai.metathings.service.identityd.ListRolesForGroupOnProjectResponse")
}

func init() {
	proto.RegisterFile("list_roles_for_group_on_project.proto", fileDescriptor_list_roles_for_group_on_project_ed686b9997180ad8)
}

var fileDescriptor_list_roles_for_group_on_project_ed686b9997180ad8 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0xb6, 0x62, 0xd5, 0xf4, 0xb6, 0xa7, 0x52, 0xc4, 0xd6, 0x42, 0xa1, 0x97, 0x66,
	0xa1, 0x82, 0x17, 0x0f, 0x82, 0x82, 0x52, 0x10, 0x94, 0x15, 0xbc, 0x2e, 0x69, 0x33, 0x4d, 0x47,
	0xb7, 0x99, 0x98, 0xcc, 0xb6, 0xf8, 0x48, 0x3e, 0x95, 0xe0, 0x93, 0xc8, 0x6e, 0xaa, 0x47, 0xc5,
	0xdb, 0x40, 0xf2, 0xff, 0x7c, 0xfc, 0x62, 0x54, 0x62, 0xe0, 0xc2, 0x53, 0x09, 0xa1, 0x58, 0x92,
	0x2f, 0x8c, 0xa7, 0xca, 0x15, 0x64, 0x0b, 0xe7, 0xe9, 0x19, 0x16, 0x2c, 0x9d, 0x27, 0xa6, 0xb4,
	0xaf, 0x50, 0xae, 0x81, 0x15, 0xaf, 0xd0, 0x9a, 0x20, 0x03, 0xf8, 0x0d, 0x2e, 0x40, 0xa2, 0x06,
	0xcb, 0xc8, 0x6f, 0xba, 0x77, 0x62, 0x88, 0x4c, 0x09, 0x59, 0xf3, 0x7d, 0x5e, 0x2d, 0xb3, 0xad,
	0x57, 0xce, 0x81, 0x0f, 0xb1, 0xa0, 0x77, 0x6e, 0x90, 0x57, 0xd5, 0x5c, 0x2e, 0x68, 0x9d, 0xad,
	0xb7, 0xc8, 0x2f, 0xb4, 0xcd, 0x0c, 0x4d, 0x9a, 0xc7, 0xc9, 0x46, 0x95, 0xa8, 0x15, 0x93, 0x0f,
	0xd9, 0xcf, 0xb9, 0xcb, 0x89, 0x9a, 0x16, 0xef, 0xe1, 0x7b, 0x22, 0x4e, 0xef, 0x30, 0x70, 0x5e,
	0x6b, 0x6f, 0xc8, 0xdf, 0xd6, 0xd6, 0x7b, 0xfb, 0x10, 0xa5, 0x39, 0xbc, 0x56, 0x10, 0x38, 0xbd,
	0x16, 0x62, 0x67, 0x2f, 0x50, 0x77, 0x93, 0x41, 0x32, 0xee, 0x4c, 0x8f, 0x65, 0xe4, 0xc9, 0x6f,
	0x9e, 0x7c, 0x64, 0x8f, 0xd6, 0x3c, 0xa9, 0xb2, 0x82, 0xab, 0xf6, 0xe7, 0x47, 0xbf, 0x35, 0x48,
	0xf2, 0xa3, 0x5d, 0x6e, 0xa6, 0xd3, 0x4b, 0x71, 0x18, 0x97, 0x40, 0xdd, 0x6d, 0xfd, 0xa3, 0xe2,
	0xa0, 0x49, 0xcd, 0xf4, 0x50, 0x89, 0xe1, 0x6f, 0xd4, 0xe0, 0xc8, 0x06, 0x48, 0x2f, 0xc4, 0x7e,
	0x33, 0x7d, 0x37, 0x19, 0xec, 0x8d, 0x3b, 0xd3, 0x91, 0xfc, 0x63, 0x66, 0x59, 0xf7, 0xe5, 0x31,
	0x33, 0x6f, 0x37, 0x92, 0xb3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x3e, 0xa1, 0x71, 0xc3,
	0x01, 0x00, 0x00,
}
