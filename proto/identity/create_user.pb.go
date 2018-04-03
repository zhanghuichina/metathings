// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_user.proto

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

type CreateUserRequest struct {
	DomainId         *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	DefaultProjectId *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=default_project_id,json=defaultProjectId" json:"default_project_id,omitempty"`
	Name             *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Password         *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	Enabled          *google_protobuf.BoolValue   `protobuf:"bytes,5,opt,name=enabled" json:"enabled,omitempty"`
	Extra            map[string]string            `protobuf:"bytes,6,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *CreateUserRequest) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *CreateUserRequest) GetDefaultProjectId() *google_protobuf.StringValue {
	if m != nil {
		return m.DefaultProjectId
	}
	return nil
}

func (m *CreateUserRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateUserRequest) GetPassword() *google_protobuf.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *CreateUserRequest) GetEnabled() *google_protobuf.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *CreateUserRequest) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type CreateUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *CreateUserResponse) Reset()                    { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()               {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "ai.metathings.service.identity.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "ai.metathings.service.identity.CreateUserResponse")
}

func init() { proto.RegisterFile("create_user.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0xab, 0xda, 0x30,
	0x1c, 0xc7, 0xa9, 0x55, 0xa7, 0xf1, 0xa2, 0x61, 0x87, 0x52, 0x86, 0x13, 0xd9, 0xc1, 0x8b, 0x29,
	0xb8, 0x31, 0x64, 0xec, 0xb0, 0x39, 0x3c, 0xb8, 0xc3, 0x18, 0x1d, 0xdb, 0x55, 0xd2, 0xe6, 0x67,
	0xcd, 0x6c, 0x93, 0x2e, 0x49, 0xed, 0xfc, 0x6b, 0x1f, 0xf8, 0x97, 0x3c, 0x9a, 0xea, 0xf3, 0x81,
	0xf0, 0xc4, 0x5b, 0xda, 0xe6, 0xf3, 0xed, 0x27, 0xbf, 0x6f, 0xd0, 0x20, 0x56, 0x40, 0x0d, 0xac,
	0x0b, 0x0d, 0x8a, 0xe4, 0x4a, 0x1a, 0x89, 0x87, 0x94, 0x93, 0x0c, 0x0c, 0x35, 0x5b, 0x2e, 0x12,
	0x4d, 0x34, 0xa8, 0x3d, 0x8f, 0x81, 0x70, 0x06, 0xc2, 0x70, 0x73, 0xf0, 0x87, 0x89, 0x94, 0x49,
	0x0a, 0x81, 0xdd, 0x1d, 0x15, 0x9b, 0xa0, 0x54, 0x34, 0xcf, 0x41, 0xe9, 0x9a, 0xf7, 0x3f, 0x26,
	0xdc, 0x6c, 0x8b, 0x88, 0xc4, 0x32, 0x0b, 0xb2, 0x92, 0x9b, 0x9d, 0x2c, 0x83, 0x44, 0x4e, 0xed,
	0xc7, 0xe9, 0x9e, 0xa6, 0x9c, 0x51, 0x23, 0x95, 0x0e, 0x9e, 0x96, 0x27, 0x0e, 0x5d, 0x1c, 0xc6,
	0x47, 0x17, 0x0d, 0xbe, 0x59, 0xb3, 0xdf, 0x1a, 0x54, 0x08, 0xff, 0x0a, 0xd0, 0x06, 0x7f, 0x45,
	0x5d, 0x26, 0x33, 0xca, 0xc5, 0x9a, 0x33, 0xcf, 0x19, 0x39, 0x93, 0xde, 0xec, 0x0d, 0xa9, 0x6d,
	0xc8, 0xd9, 0x86, 0xfc, 0x32, 0x8a, 0x8b, 0xe4, 0x0f, 0x4d, 0x0b, 0x58, 0xb4, 0x8f, 0x0f, 0x6f,
	0x1b, 0x23, 0x27, 0xec, 0xd4, 0xd8, 0x8a, 0xe1, 0xef, 0x08, 0x33, 0xd8, 0xd0, 0x22, 0x35, 0xeb,
	0x5c, 0xc9, 0xbf, 0x10, 0x9b, 0x2a, 0xab, 0x71, 0x3b, 0x2b, 0xec, 0x9f, 0xb8, 0x9f, 0x35, 0xb6,
	0x62, 0x78, 0x8e, 0x9a, 0x82, 0x66, 0xe0, 0xb9, 0x77, 0x98, 0x58, 0x02, 0x7f, 0x41, 0x9d, 0x9c,
	0x6a, 0x5d, 0x4a, 0xc5, 0xbc, 0xe6, 0x3d, 0xe7, 0x38, 0x53, 0xf8, 0x03, 0x7a, 0x05, 0x82, 0x46,
	0x29, 0x30, 0xaf, 0x65, 0x03, 0xfc, 0xab, 0x80, 0x85, 0x94, 0x69, 0xad, 0x7e, 0xde, 0x8a, 0x43,
	0xd4, 0x82, 0xff, 0x46, 0x51, 0xaf, 0x3d, 0x72, 0x27, 0xbd, 0xd9, 0x67, 0xf2, 0x72, 0xd5, 0xe4,
	0xaa, 0x02, 0xb2, 0xac, 0xf0, 0xa5, 0x30, 0xea, 0x10, 0xd6, 0x51, 0xfe, 0x1c, 0xa1, 0xcb, 0x4b,
	0xdc, 0x47, 0xee, 0x0e, 0x0e, 0xb6, 0x9c, 0x6e, 0x58, 0x2d, 0xf1, 0x6b, 0xd4, 0xda, 0x57, 0x16,
	0x76, 0xc8, 0xdd, 0xb0, 0x7e, 0xf8, 0xd4, 0x98, 0x3b, 0xe3, 0x1f, 0x08, 0x3f, 0xff, 0x81, 0xce,
	0xa5, 0xd0, 0x50, 0x4d, 0xb5, 0xba, 0x08, 0xa7, 0x7e, 0xdf, 0xdd, 0x52, 0xb4, 0xac, 0x25, 0xa2,
	0xb6, 0x3d, 0xfa, 0xfb, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x75, 0xa2, 0xb5, 0xd4, 0x02,
	0x00, 0x00,
}
