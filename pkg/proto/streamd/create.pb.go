// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create.proto

package streamd

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

type CreateRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Pads                 []*OpPad              `protobuf:"bytes,2,rep,name=pads" json:"pads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_create_ef63e165a177e2e1, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateRequest) GetPads() []*OpPad {
	if m != nil {
		return m.Pads
	}
	return nil
}

type CreateResponse struct {
	Stream               *Stream  `protobuf:"bytes,1,opt,name=stream" json:"stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_create_ef63e165a177e2e1, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "ai.metathings.service.streamd.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "ai.metathings.service.streamd.CreateResponse")
}

func init() { proto.RegisterFile("create.proto", fileDescriptor_create_ef63e165a177e2e1) }

var fileDescriptor_create_ef63e165a177e2e1 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x55, 0xa8, 0x3a, 0xb8, 0x85, 0x21, 0x53, 0x55, 0x01, 0xaa, 0x2a, 0x90, 0xba, 0xf4,
	0x82, 0x8a, 0x84, 0x58, 0x98, 0x78, 0x80, 0xa2, 0x54, 0x62, 0xbf, 0x24, 0x87, 0x6b, 0x11, 0xc7,
	0xc6, 0x77, 0x69, 0x06, 0x5e, 0x1e, 0xc9, 0x4e, 0x19, 0xe9, 0x64, 0x5b, 0xfe, 0xbf, 0xfb, 0xbf,
	0x53, 0xb3, 0x2a, 0x10, 0x0a, 0x81, 0x0f, 0x4e, 0x5c, 0x76, 0x8b, 0x06, 0x2c, 0x09, 0xca, 0xc1,
	0xb4, 0x9a, 0x81, 0x29, 0x1c, 0x4d, 0x45, 0xc0, 0x12, 0x08, 0x6d, 0xbd, 0xb8, 0xd3, 0xce, 0xe9,
	0x86, 0xf2, 0x18, 0x2e, 0xbb, 0xcf, 0xbc, 0x0f, 0xe8, 0x3d, 0x05, 0x4e, 0xf8, 0xe2, 0x59, 0x1b,
	0x39, 0x74, 0x25, 0x54, 0xce, 0xe6, 0xb6, 0x37, 0xf2, 0xe5, 0xfa, 0x5c, 0xbb, 0x4d, 0xfc, 0xdc,
	0x1c, 0xb1, 0x31, 0x35, 0x8a, 0x0b, 0x9c, 0xff, 0x5d, 0x07, 0x6e, 0x96, 0x0a, 0xd2, 0x6b, 0xf5,
	0xa3, 0xae, 0xde, 0xa2, 0x54, 0x41, 0xdf, 0x1d, 0xb1, 0x64, 0x8f, 0x6a, 0xdc, 0xa2, 0xa5, 0xf9,
	0x68, 0x39, 0x5a, 0x4f, 0xb7, 0x37, 0x90, 0x2c, 0xe0, 0x64, 0x01, 0x7b, 0x09, 0xa6, 0xd5, 0x1f,
	0xd8, 0x74, 0x54, 0xc4, 0x64, 0xf6, 0xa2, 0xc6, 0x1e, 0x6b, 0x9e, 0x5f, 0x2c, 0x2f, 0xd7, 0xd3,
	0xed, 0x3d, 0xfc, 0xbb, 0x16, 0xec, 0xfc, 0x3b, 0xd6, 0x45, 0x24, 0x56, 0x3b, 0x75, 0x7d, 0x2a,
	0x67, 0xef, 0x5a, 0xa6, 0xec, 0x55, 0x4d, 0x52, 0x70, 0xe8, 0x7f, 0x38, 0x33, 0x6d, 0x1f, 0xcf,
	0x62, 0x80, 0xca, 0x49, 0xd4, 0x7c, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x21, 0x21, 0x64, 0x46,
	0x69, 0x01, 0x00, 0x00,
}
