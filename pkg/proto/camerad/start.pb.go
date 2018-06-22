// Code generated by protoc-gen-go. DO NOT EDIT.
// source: start.proto

package camerad

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

type StartRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_start_2ee9d24024e560bf, []int{0}
}
func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (dst *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(dst, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

type StartResponse struct {
	Camera               *Camera  `protobuf:"bytes,1,opt,name=camera" json:"camera,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_start_2ee9d24024e560bf, []int{1}
}
func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (dst *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(dst, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

func (m *StartResponse) GetCamera() *Camera {
	if m != nil {
		return m.Camera
	}
	return nil
}

func init() {
	proto.RegisterType((*StartRequest)(nil), "ai.metathings.service.camerad.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "ai.metathings.service.camerad.StartResponse")
}

func init() { proto.RegisterFile("start.proto", fileDescriptor_start_2ee9d24024e560bf) }

var fileDescriptor_start_2ee9d24024e560bf = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x3f, 0x4f, 0xc3, 0x40,
	0x0c, 0xc5, 0x95, 0x0c, 0x19, 0xae, 0x65, 0xc9, 0x84, 0x2a, 0xfe, 0x54, 0x95, 0x90, 0x58, 0xea,
	0x93, 0x00, 0xb1, 0xb1, 0x00, 0x33, 0x43, 0x2a, 0xb1, 0x3b, 0x89, 0xb9, 0x5a, 0x24, 0x71, 0xb8,
	0x73, 0x92, 0x8f, 0x8b, 0xc4, 0x27, 0x41, 0xe4, 0x02, 0x23, 0x93, 0x6d, 0xd9, 0xef, 0xbd, 0x9f,
	0xcd, 0x2a, 0x28, 0x7a, 0x85, 0xde, 0x8b, 0x4a, 0x7e, 0x8e, 0x0c, 0x2d, 0x29, 0xea, 0x91, 0x3b,
	0x17, 0x20, 0x90, 0x1f, 0xb9, 0x22, 0xa8, 0xb0, 0x25, 0x8f, 0xf5, 0xe6, 0xc2, 0x89, 0xb8, 0x86,
	0xec, 0x7c, 0x5c, 0x0e, 0x6f, 0x76, 0xf2, 0xd8, 0xf7, 0xe4, 0x43, 0x94, 0x6f, 0xee, 0x1d, 0xeb,
	0x71, 0x28, 0xa1, 0x92, 0xd6, 0xb6, 0x13, 0xeb, 0xbb, 0x4c, 0xd6, 0xc9, 0x7e, 0x5e, 0xee, 0x47,
	0x6c, 0xb8, 0x46, 0x15, 0x1f, 0xec, 0x5f, 0xbb, 0xe8, 0xd6, 0x31, 0x20, 0x4e, 0xbb, 0x67, 0xb3,
	0x3e, 0xfc, 0x30, 0x15, 0xf4, 0x31, 0x50, 0xd0, 0xfc, 0xce, 0xa4, 0x5c, 0x9f, 0x26, 0xdb, 0xe4,
	0x7a, 0x75, 0x73, 0x06, 0x11, 0x01, 0x7e, 0x11, 0xe0, 0xa0, 0x9e, 0x3b, 0xf7, 0x8a, 0xcd, 0x40,
	0x8f, 0xd9, 0xd7, 0xe7, 0x65, 0xba, 0x4d, 0x8a, 0x94, 0xeb, 0xdd, 0x8b, 0x39, 0x59, 0x5c, 0x42,
	0x2f, 0x5d, 0xa0, 0xfc, 0xc1, 0x64, 0x31, 0x66, 0xb1, 0xba, 0x82, 0x7f, 0x9f, 0x85, 0xa7, 0xb9,
	0x16, 0x8b, 0xa8, 0xcc, 0xe6, 0xc4, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0xa4, 0x4e,
	0x74, 0x30, 0x01, 0x00, 0x00,
}