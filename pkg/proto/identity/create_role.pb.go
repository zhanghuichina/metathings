// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_role.proto

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

type CreateRoleRequest struct {
	Name     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DomainId *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
}

func (m *CreateRoleRequest) Reset()                    { *m = CreateRoleRequest{} }
func (m *CreateRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRoleRequest) ProtoMessage()               {}
func (*CreateRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *CreateRoleRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateRoleRequest) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

type CreateRoleResponse struct {
	Role *Role `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
}

func (m *CreateRoleResponse) Reset()                    { *m = CreateRoleResponse{} }
func (m *CreateRoleResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateRoleResponse) ProtoMessage()               {}
func (*CreateRoleResponse) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{1} }

func (m *CreateRoleResponse) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRoleRequest)(nil), "ai.metathings.service.identity.CreateRoleRequest")
	proto.RegisterType((*CreateRoleResponse)(nil), "ai.metathings.service.identity.CreateRoleResponse")
}

func init() { proto.RegisterFile("create_role.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc7, 0x71, 0xb6, 0x3c, 0x94, 0xc7, 0x78, 0xea, 0x9e, 0x4a, 0x91, 0x5a, 0x8a, 0x07, 0x2f,
	0x9d, 0x05, 0x05, 0xd1, 0xab, 0x9e, 0xbc, 0x78, 0x58, 0xc1, 0x6b, 0xc9, 0xee, 0x8e, 0xe9, 0x60,
	0x92, 0x59, 0x93, 0xd9, 0x2e, 0xbe, 0x03, 0xdf, 0xa5, 0xe0, 0x2b, 0x91, 0x66, 0xeb, 0x9f, 0x93,
	0xde, 0x02, 0xf9, 0x7d, 0xc9, 0x87, 0xa8, 0x49, 0x1d, 0x50, 0x0b, 0xae, 0x03, 0x5b, 0x84, 0x36,
	0xb0, 0x70, 0x3e, 0xd7, 0x04, 0x0e, 0x45, 0xcb, 0x86, 0xbc, 0x89, 0x10, 0x31, 0x6c, 0xa9, 0x46,
	0xa0, 0x06, 0xbd, 0x90, 0xbc, 0xcc, 0xe6, 0x86, 0xd9, 0x58, 0x2c, 0xd2, 0xba, 0xea, 0x1e, 0x8b,
	0x3e, 0xe8, 0xb6, 0xc5, 0x10, 0x87, 0x7e, 0x76, 0x61, 0x48, 0x36, 0x5d, 0x05, 0x35, 0xbb, 0xc2,
	0xf5, 0x24, 0x4f, 0xdc, 0x17, 0x86, 0x57, 0xe9, 0x72, 0xb5, 0xd5, 0x96, 0x1a, 0x2d, 0x1c, 0x62,
	0xf1, 0x75, 0xdc, 0x77, 0xea, 0xdb, 0xb0, 0x7c, 0xcd, 0xd4, 0xe4, 0x26, 0xc9, 0x4a, 0xb6, 0x58,
	0xe2, 0x73, 0x87, 0x51, 0xf2, 0x4b, 0xf5, 0xcf, 0x6b, 0x87, 0xd3, 0x6c, 0x91, 0x9d, 0x1e, 0x9e,
	0x1d, 0xc1, 0x00, 0x81, 0x4f, 0x08, 0xdc, 0x4b, 0x20, 0x6f, 0x1e, 0xb4, 0xed, 0xf0, 0x7a, 0xfc,
	0xfe, 0x76, 0x3c, 0x5a, 0x64, 0x65, 0x2a, 0xf2, 0x2b, 0x75, 0xd0, 0xb0, 0xd3, 0xe4, 0xd7, 0xd4,
	0x4c, 0x47, 0x7f, 0xe7, 0xe5, 0xff, 0x61, 0x7e, 0xdb, 0x2c, 0xef, 0x54, 0xfe, 0x53, 0x12, 0x5b,
	0xf6, 0x11, 0x77, 0x94, 0x1d, 0x77, 0x4f, 0x39, 0x81, 0xdf, 0xff, 0x0c, 0x52, 0x9b, 0x8a, 0x6a,
	0x9c, 0xde, 0x3b, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x7a, 0x73, 0x6c, 0x7a, 0x01, 0x00,
	0x00,
}