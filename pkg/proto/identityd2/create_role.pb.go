// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_role.proto

package identityd2

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

type CreateRoleRequest struct {
	Id                   *wrappers.StringValue            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Domain               *OpDomain                        `protobuf:"bytes,2,opt,name=domain" json:"domain,omitempty"`
	Name                 *wrappers.StringValue            `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Alias                *wrappers.StringValue            `protobuf:"bytes,4,opt,name=alias" json:"alias,omitempty"`
	Description          *wrappers.StringValue            `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Extra                map[string]*wrappers.StringValue `protobuf:"bytes,7,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CreateRoleRequest) Reset()         { *m = CreateRoleRequest{} }
func (m *CreateRoleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoleRequest) ProtoMessage()    {}
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_create_role_9f755e863ae7f66b, []int{0}
}
func (m *CreateRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleRequest.Unmarshal(m, b)
}
func (m *CreateRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleRequest.Merge(dst, src)
}
func (m *CreateRoleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoleRequest.Size(m)
}
func (m *CreateRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleRequest proto.InternalMessageInfo

func (m *CreateRoleRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CreateRoleRequest) GetDomain() *OpDomain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *CreateRoleRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateRoleRequest) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *CreateRoleRequest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CreateRoleRequest) GetExtra() map[string]*wrappers.StringValue {
	if m != nil {
		return m.Extra
	}
	return nil
}

type CreateRoleResponse struct {
	Role                 *Role    `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoleResponse) Reset()         { *m = CreateRoleResponse{} }
func (m *CreateRoleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRoleResponse) ProtoMessage()    {}
func (*CreateRoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_create_role_9f755e863ae7f66b, []int{1}
}
func (m *CreateRoleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleResponse.Unmarshal(m, b)
}
func (m *CreateRoleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleResponse.Marshal(b, m, deterministic)
}
func (dst *CreateRoleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleResponse.Merge(dst, src)
}
func (m *CreateRoleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRoleResponse.Size(m)
}
func (m *CreateRoleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleResponse proto.InternalMessageInfo

func (m *CreateRoleResponse) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRoleRequest)(nil), "ai.metathings.service.identityd2.CreateRoleRequest")
	proto.RegisterMapType((map[string]*wrappers.StringValue)(nil), "ai.metathings.service.identityd2.CreateRoleRequest.ExtraEntry")
	proto.RegisterType((*CreateRoleResponse)(nil), "ai.metathings.service.identityd2.CreateRoleResponse")
}

func init() { proto.RegisterFile("create_role.proto", fileDescriptor_create_role_9f755e863ae7f66b) }

var fileDescriptor_create_role_9f755e863ae7f66b = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0xab, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0xd2, 0x54, 0xdd, 0x5c, 0x7c, 0x7b, 0x0a, 0x45, 0x34, 0xbc, 0x83, 0x3c, 0xc4,
	0xb7, 0x81, 0x08, 0xf2, 0xe8, 0xa1, 0x07, 0xb5, 0x17, 0x2f, 0x4a, 0x94, 0x5e, 0x65, 0x9b, 0x1d,
	0xd3, 0xa1, 0x9b, 0xdd, 0xb8, 0xbb, 0x69, 0xed, 0x5f, 0x2b, 0xf8, 0x6f, 0x78, 0x91, 0x6c, 0x62,
	0x2d, 0x08, 0x46, 0x6f, 0x93, 0x64, 0x7e, 0xdf, 0x7c, 0xf3, 0x4d, 0xc8, 0x55, 0x65, 0x80, 0x3b,
	0xf8, 0x64, 0xb4, 0x04, 0xd6, 0x1a, 0xed, 0x34, 0xcd, 0x38, 0xb2, 0x06, 0x1c, 0x77, 0x3b, 0x54,
	0xb5, 0x65, 0x16, 0xcc, 0x01, 0x2b, 0x60, 0x28, 0x40, 0x39, 0x74, 0x27, 0x51, 0x2c, 0x1e, 0xd7,
	0x5a, 0xd7, 0x12, 0x72, 0xdf, 0xbf, 0xed, 0x3e, 0xe7, 0x47, 0xc3, 0xdb, 0x16, 0x8c, 0x1d, 0x14,
	0x16, 0x2f, 0x6b, 0x74, 0xbb, 0x6e, 0xcb, 0x2a, 0xdd, 0xe4, 0xcd, 0x11, 0xdd, 0x5e, 0x1f, 0xf3,
	0x5a, 0xdf, 0xfa, 0x8f, 0xb7, 0x07, 0x2e, 0x51, 0x70, 0xa7, 0x8d, 0xcd, 0xcf, 0xe5, 0xc8, 0x25,
	0x8d, 0x16, 0x20, 0x87, 0x87, 0xeb, 0x1f, 0x11, 0xb9, 0x7a, 0xed, 0xcd, 0x95, 0x5a, 0x42, 0x09,
	0x5f, 0x3a, 0xb0, 0x8e, 0x3e, 0x27, 0x21, 0x8a, 0x34, 0xc8, 0x82, 0x9b, 0xa4, 0x78, 0xc4, 0x06,
	0x1f, 0xec, 0x97, 0x0f, 0xf6, 0xc1, 0x19, 0x54, 0xf5, 0x86, 0xcb, 0x0e, 0xca, 0x10, 0x05, 0x7d,
	0x4b, 0xe6, 0x42, 0x37, 0x1c, 0x55, 0x1a, 0x7a, 0xe2, 0x19, 0x9b, 0xda, 0x8d, 0xbd, 0x6b, 0xdf,
	0x78, 0xe2, 0xd5, 0xfc, 0xfb, 0xb7, 0x27, 0x61, 0x16, 0x94, 0xa3, 0x02, 0xbd, 0x23, 0x33, 0xc5,
	0x1b, 0x48, 0xa3, 0xe9, 0xd9, 0x67, 0xd6, 0x13, 0x74, 0x49, 0x62, 0x2e, 0x91, 0xdb, 0x74, 0xf6,
	0x1f, 0xe8, 0x80, 0xd0, 0x15, 0x49, 0x04, 0xd8, 0xca, 0x60, 0xeb, 0x50, 0xab, 0x34, 0xfe, 0x87,
	0xc5, 0x2f, 0x01, 0xfa, 0x91, 0xc4, 0xf0, 0xd5, 0x19, 0x9e, 0xde, 0xcb, 0xa2, 0x9b, 0xa4, 0x58,
	0x4d, 0x07, 0xf0, 0x47, 0xe6, 0x6c, 0xdd, 0x0b, 0xac, 0x95, 0x33, 0xa7, 0x72, 0x10, 0x5b, 0x6c,
	0x08, 0xf9, 0xfd, 0x92, 0x3e, 0x24, 0xd1, 0x1e, 0x4e, 0xfe, 0x28, 0x0f, 0xca, 0xbe, 0xa4, 0x05,
	0x89, 0x0f, 0xbd, 0x97, 0x31, 0xf6, 0xbf, 0xfb, 0x1d, 0x5a, 0x97, 0xe1, 0x5d, 0x50, 0xde, 0x6f,
	0xb5, 0xc4, 0x0a, 0xc1, 0x5e, 0xbf, 0x27, 0xf4, 0xd2, 0x88, 0x6d, 0xb5, 0xb2, 0x7d, 0x92, 0xb3,
	0xfe, 0x47, 0x1d, 0xef, 0xff, 0x74, 0x7a, 0x19, 0x4f, 0x7b, 0x66, 0x3b, 0xf7, 0xc3, 0x5f, 0xfc,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xab, 0xef, 0x56, 0xf2, 0x02, 0x00, 0x00,
}