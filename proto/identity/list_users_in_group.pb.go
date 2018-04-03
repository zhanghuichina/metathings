// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_users_in_group.proto

package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListUsersInGroupRequest struct {
	GroupId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
}

func (m *ListUsersInGroupRequest) Reset()                    { *m = ListUsersInGroupRequest{} }
func (m *ListUsersInGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*ListUsersInGroupRequest) ProtoMessage()               {}
func (*ListUsersInGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{0} }

func (m *ListUsersInGroupRequest) GetGroupId() *google_protobuf.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

type ListUsersInGroupResponse struct {
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *ListUsersInGroupResponse) Reset()                    { *m = ListUsersInGroupResponse{} }
func (m *ListUsersInGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*ListUsersInGroupResponse) ProtoMessage()               {}
func (*ListUsersInGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{1} }

func (m *ListUsersInGroupResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*ListUsersInGroupRequest)(nil), "ai.metathings.service.identity.ListUsersInGroupRequest")
	proto.RegisterType((*ListUsersInGroupResponse)(nil), "ai.metathings.service.identity.ListUsersInGroupResponse")
}

func init() { proto.RegisterFile("list_users_in_group.proto", fileDescriptor40) }

var fileDescriptor40 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4b, 0xc3, 0x40,
	0x14, 0xc6, 0x89, 0x62, 0x95, 0xeb, 0xd6, 0xc5, 0x58, 0xa4, 0x86, 0xe2, 0xd0, 0xa5, 0x17, 0xa8,
	0xe0, 0xe0, 0x22, 0xb8, 0x48, 0xc1, 0x29, 0x62, 0x07, 0x97, 0x70, 0x69, 0x9e, 0xd7, 0x87, 0xc9,
	0xdd, 0x79, 0xef, 0x5d, 0x83, 0x7f, 0xad, 0xe0, 0x5f, 0x22, 0xb9, 0xa8, 0x8b, 0xe0, 0xf6, 0xc1,
	0xdd, 0xef, 0x7b, 0x3f, 0x3e, 0x71, 0xd6, 0x20, 0x71, 0x19, 0x08, 0x3c, 0x95, 0x68, 0x4a, 0xed,
	0x6d, 0x70, 0xd2, 0x79, 0xcb, 0x76, 0x32, 0x53, 0x28, 0x5b, 0x60, 0xc5, 0x3b, 0x34, 0x9a, 0x24,
	0x81, 0xdf, 0xe3, 0x16, 0x24, 0xd6, 0x60, 0x18, 0xf9, 0x7d, 0x3a, 0xd3, 0xd6, 0xea, 0x06, 0xf2,
	0xf8, 0xbb, 0x0a, 0x2f, 0x79, 0xe7, 0x95, 0x73, 0xe0, 0x69, 0xe0, 0xa7, 0xd7, 0x1a, 0x79, 0x17,
	0x2a, 0xb9, 0xb5, 0x6d, 0xde, 0x76, 0xc8, 0xaf, 0xb6, 0xcb, 0xb5, 0x5d, 0xc6, 0xc7, 0xe5, 0x5e,
	0x35, 0x58, 0x2b, 0xb6, 0x9e, 0xf2, 0xdf, 0xf8, 0xcd, 0x89, 0xde, 0x66, 0xc8, 0xf3, 0x67, 0x71,
	0xfa, 0x80, 0xc4, 0x4f, 0xbd, 0xdf, 0xda, 0xdc, 0xf7, 0x76, 0x05, 0xbc, 0x05, 0x20, 0x9e, 0xdc,
	0x8a, 0x93, 0x68, 0x5b, 0x62, 0x9d, 0x26, 0x59, 0xb2, 0x18, 0xaf, 0xce, 0xe5, 0x60, 0x24, 0x7f,
	0x8c, 0xe4, 0x23, 0x7b, 0x34, 0x7a, 0xa3, 0x9a, 0x00, 0x77, 0xa3, 0xcf, 0x8f, 0x8b, 0x83, 0x2c,
	0x29, 0x8e, 0x23, 0xb5, 0xae, 0xe7, 0x1b, 0x91, 0xfe, 0xed, 0x26, 0x67, 0x0d, 0xc1, 0xe4, 0x46,
	0x1c, 0xc5, 0x4d, 0xd2, 0x24, 0x3b, 0x5c, 0x8c, 0x57, 0x97, 0xf2, 0xff, 0x2d, 0x64, 0x5f, 0x52,
	0x0c, 0x48, 0x35, 0x8a, 0xe7, 0xaf, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xda, 0xcc, 0xd1, 0x8d,
	0x5b, 0x01, 0x00, 0x00,
}
