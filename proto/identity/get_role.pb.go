// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_role.proto

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

type GetRoleRequest struct {
	RoleId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=role_id,json=roleId" json:"role_id,omitempty"`
}

func (m *GetRoleRequest) Reset()                    { *m = GetRoleRequest{} }
func (m *GetRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRoleRequest) ProtoMessage()               {}
func (*GetRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func (m *GetRoleRequest) GetRoleId() *google_protobuf.StringValue {
	if m != nil {
		return m.RoleId
	}
	return nil
}

type GetRoleResponse struct {
	Role *Role `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
}

func (m *GetRoleResponse) Reset()                    { *m = GetRoleResponse{} }
func (m *GetRoleResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRoleResponse) ProtoMessage()               {}
func (*GetRoleResponse) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

func (m *GetRoleResponse) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRoleRequest)(nil), "ai.metathings.service.identity.GetRoleRequest")
	proto.RegisterType((*GetRoleResponse)(nil), "ai.metathings.service.identity.GetRoleResponse")
}

func init() { proto.RegisterFile("get_role.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x48, 0x84, 0x15, 0x2a, 0xe4, 0x24, 0x45, 0x6a, 0x29, 0x1e, 0xbc, 0x74, 0x16,
	0x14, 0xc4, 0x8b, 0x17, 0x2f, 0x22, 0x1e, 0x84, 0x08, 0x5e, 0xcb, 0xa6, 0x19, 0xb7, 0x83, 0x9b,
	0x4c, 0xdc, 0x9d, 0x6d, 0xf0, 0x69, 0x05, 0x9f, 0x44, 0xb2, 0xad, 0xf1, 0xe6, 0x6d, 0xe0, 0x9f,
	0xef, 0x9f, 0x6f, 0xd4, 0xc4, 0xa2, 0xac, 0x3c, 0x3b, 0x84, 0xce, 0xb3, 0x70, 0x31, 0x33, 0x04,
	0x0d, 0x8a, 0x91, 0x0d, 0xb5, 0x36, 0x40, 0x40, 0xbf, 0xa5, 0x35, 0x02, 0xd5, 0xd8, 0x0a, 0xc9,
	0xe7, 0x74, 0x66, 0x99, 0xad, 0x43, 0x9d, 0xb6, 0xab, 0xf8, 0xa6, 0x7b, 0x6f, 0xba, 0x0e, 0x7d,
	0xd8, 0xf1, 0xd3, 0x1b, 0x4b, 0xb2, 0x89, 0x15, 0xac, 0xb9, 0xd1, 0x4d, 0x4f, 0xf2, 0xce, 0xbd,
	0xb6, 0xbc, 0x4c, 0xe1, 0x72, 0x6b, 0x1c, 0xd5, 0x46, 0xd8, 0x07, 0x3d, 0x8e, 0x7b, 0x4e, 0xfd,
	0x39, 0x2c, 0x9e, 0xd5, 0xe4, 0x01, 0xa5, 0x64, 0x87, 0x25, 0x7e, 0x44, 0x0c, 0x52, 0xdc, 0xa9,
	0xa3, 0x21, 0x5f, 0x51, 0x7d, 0x9a, 0xcd, 0xb3, 0xcb, 0xe3, 0xab, 0x33, 0xd8, 0x79, 0xc0, 0xaf,
	0x07, 0xbc, 0x88, 0xa7, 0xd6, 0xbe, 0x1a, 0x17, 0xf1, 0x3e, 0xff, 0xfe, 0x3a, 0x3f, 0x98, 0x67,
	0x65, 0x3e, 0x40, 0x8f, 0xf5, 0xe2, 0x49, 0x9d, 0x8c, 0x85, 0xa1, 0xe3, 0x36, 0x60, 0x71, 0xab,
	0x0e, 0x87, 0x70, 0x5f, 0x77, 0x01, 0xff, 0xbf, 0x0d, 0x89, 0x4d, 0x44, 0x95, 0xa7, 0x93, 0xd7,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0xd3, 0xa3, 0x9e, 0x3a, 0x01, 0x00, 0x00,
}
