// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_core.proto

package cored

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

type CreateCoreRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateCoreRequest) Reset()         { *m = CreateCoreRequest{} }
func (m *CreateCoreRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCoreRequest) ProtoMessage()    {}
func (*CreateCoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_create_core_f428acebdfc7b4c9, []int{0}
}
func (m *CreateCoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCoreRequest.Unmarshal(m, b)
}
func (m *CreateCoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCoreRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCoreRequest.Merge(dst, src)
}
func (m *CreateCoreRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCoreRequest.Size(m)
}
func (m *CreateCoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCoreRequest proto.InternalMessageInfo

func (m *CreateCoreRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type CreateCoreResponse struct {
	Core                 *Core    `protobuf:"bytes,1,opt,name=core" json:"core,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCoreResponse) Reset()         { *m = CreateCoreResponse{} }
func (m *CreateCoreResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCoreResponse) ProtoMessage()    {}
func (*CreateCoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_create_core_f428acebdfc7b4c9, []int{1}
}
func (m *CreateCoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCoreResponse.Unmarshal(m, b)
}
func (m *CreateCoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCoreResponse.Marshal(b, m, deterministic)
}
func (dst *CreateCoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCoreResponse.Merge(dst, src)
}
func (m *CreateCoreResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCoreResponse.Size(m)
}
func (m *CreateCoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCoreResponse proto.InternalMessageInfo

func (m *CreateCoreResponse) GetCore() *Core {
	if m != nil {
		return m.Core
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateCoreRequest)(nil), "ai.metathings.service.cored.CreateCoreRequest")
	proto.RegisterType((*CreateCoreResponse)(nil), "ai.metathings.service.cored.CreateCoreResponse")
}

func init() { proto.RegisterFile("create_core.proto", fileDescriptor_create_core_f428acebdfc7b4c9) }

var fileDescriptor_create_core_f428acebdfc7b4c9 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xcd, 0x4a, 0x43, 0x31,
	0x10, 0x85, 0x29, 0x14, 0x17, 0x71, 0xd5, 0xbb, 0x92, 0x2a, 0xa2, 0x5d, 0xb9, 0xe9, 0x44, 0x14,
	0x7d, 0x81, 0xe2, 0xca, 0x5d, 0x05, 0xb7, 0x32, 0x37, 0x1d, 0xd3, 0xe0, 0x4d, 0x26, 0x4e, 0x26,
	0xbd, 0xaf, 0x2f, 0x4d, 0xeb, 0xcf, 0xca, 0x5d, 0x20, 0xdf, 0x39, 0xe7, 0x1b, 0x33, 0x73, 0x42,
	0xa8, 0xf4, 0xe6, 0x58, 0x08, 0xb2, 0xb0, 0x72, 0x77, 0x8e, 0x01, 0x22, 0x29, 0xea, 0x36, 0x24,
	0x5f, 0xa0, 0x90, 0xec, 0x82, 0x23, 0xd8, 0x13, 0x9b, 0xf9, 0xa5, 0x67, 0xf6, 0x03, 0xd9, 0x86,
	0xf6, 0xf5, 0xdd, 0x8e, 0x82, 0x39, 0x93, 0x94, 0x43, 0x78, 0xfe, 0xe8, 0x83, 0x6e, 0x6b, 0x0f,
	0x8e, 0xa3, 0x8d, 0x63, 0xd0, 0x0f, 0x1e, 0xad, 0xe7, 0x65, 0xfb, 0x5c, 0xee, 0x70, 0x08, 0x1b,
	0x54, 0x96, 0x62, 0x7f, 0x9e, 0xc7, 0x9c, 0xf9, 0x15, 0x58, 0x3c, 0x99, 0xd9, 0xaa, 0x59, 0xad,
	0x58, 0x68, 0x4d, 0x9f, 0x95, 0x8a, 0x76, 0xb7, 0x66, 0x9a, 0x30, 0xd2, 0xd9, 0xe4, 0x6a, 0x72,
	0x73, 0x7a, 0x77, 0x01, 0x07, 0x0f, 0xf8, 0xf6, 0x80, 0x17, 0x95, 0x90, 0xfc, 0x2b, 0x0e, 0x95,
	0xd6, 0x8d, 0x5c, 0x3c, 0x9b, 0xee, 0x6f, 0x4d, 0xc9, 0x9c, 0x0a, 0x75, 0x0f, 0x66, 0xba, 0x9f,
	0x3a, 0xf6, 0x5c, 0xc3, 0x3f, 0xc7, 0x42, 0x0b, 0x36, 0xbc, 0x3f, 0x69, 0x43, 0xf7, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x0b, 0x81, 0xd9, 0x6c, 0x30, 0x01, 0x00, 0x00,
}
