// Code generated by protoc-gen-go. DO NOT EDIT.
// source: start.proto

package camera

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

type StartConfig struct {
	Url                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Device               *wrappers.StringValue `protobuf:"bytes,2,opt,name=device" json:"device,omitempty"`
	Dimension            *Dimension            `protobuf:"bytes,3,opt,name=dimension" json:"dimension,omitempty"`
	Bitrate              *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=bitrate" json:"bitrate,omitempty"`
	Framerate            *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=framerate" json:"framerate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StartConfig) Reset()         { *m = StartConfig{} }
func (m *StartConfig) String() string { return proto.CompactTextString(m) }
func (*StartConfig) ProtoMessage()    {}
func (*StartConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_start_206a6e0f4c7e769e, []int{0}
}
func (m *StartConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartConfig.Unmarshal(m, b)
}
func (m *StartConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartConfig.Marshal(b, m, deterministic)
}
func (dst *StartConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartConfig.Merge(dst, src)
}
func (m *StartConfig) XXX_Size() int {
	return xxx_messageInfo_StartConfig.Size(m)
}
func (m *StartConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StartConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StartConfig proto.InternalMessageInfo

func (m *StartConfig) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *StartConfig) GetDevice() *wrappers.StringValue {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *StartConfig) GetDimension() *Dimension {
	if m != nil {
		return m.Dimension
	}
	return nil
}

func (m *StartConfig) GetBitrate() *wrappers.UInt32Value {
	if m != nil {
		return m.Bitrate
	}
	return nil
}

func (m *StartConfig) GetFramerate() *wrappers.UInt32Value {
	if m != nil {
		return m.Framerate
	}
	return nil
}

type StartRequest struct {
	Config               *StartConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_start_206a6e0f4c7e769e, []int{1}
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

func (m *StartRequest) GetConfig() *StartConfig {
	if m != nil {
		return m.Config
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
	return fileDescriptor_start_206a6e0f4c7e769e, []int{2}
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
	proto.RegisterType((*StartConfig)(nil), "ai.metathings.service.camera.StartConfig")
	proto.RegisterType((*StartRequest)(nil), "ai.metathings.service.camera.StartRequest")
	proto.RegisterType((*StartResponse)(nil), "ai.metathings.service.camera.StartResponse")
}

func init() { proto.RegisterFile("start.proto", fileDescriptor_start_206a6e0f4c7e769e) }

var fileDescriptor_start_206a6e0f4c7e769e = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4f, 0x02, 0x31,
	0x10, 0x85, 0x03, 0xe8, 0x1a, 0x0a, 0x5e, 0x7a, 0xda, 0x10, 0xa2, 0x84, 0x98, 0xa8, 0x07, 0xba,
	0x09, 0x18, 0x0e, 0xc6, 0x93, 0xe8, 0x81, 0x83, 0x97, 0x25, 0x9a, 0x78, 0xec, 0x2e, 0x43, 0x69,
	0xdc, 0x6d, 0xd7, 0x76, 0x16, 0xfe, 0x99, 0x7f, 0xc7, 0xc4, 0x5f, 0x62, 0xe8, 0x16, 0xf0, 0x22,
	0x7a, 0x6a, 0x9b, 0xbe, 0xef, 0xcd, 0x9b, 0x47, 0x5a, 0x16, 0xb9, 0x41, 0x56, 0x18, 0x8d, 0x9a,
	0x76, 0xb9, 0x64, 0x39, 0x20, 0xc7, 0xa5, 0x54, 0xc2, 0x32, 0x0b, 0x66, 0x25, 0x53, 0x60, 0x29,
	0xcf, 0xc1, 0xf0, 0xce, 0x99, 0xd0, 0x5a, 0x64, 0x10, 0x39, 0x6d, 0x52, 0x2e, 0xa2, 0xb5, 0xe1,
	0x45, 0x01, 0xc6, 0x56, 0x74, 0x67, 0x2c, 0x24, 0x2e, 0xcb, 0x84, 0xa5, 0x3a, 0x8f, 0xf2, 0xb5,
	0xc4, 0x37, 0xbd, 0x8e, 0x84, 0x1e, 0xb8, 0xcf, 0xc1, 0x8a, 0x67, 0x72, 0xce, 0x51, 0x1b, 0x1b,
	0xed, 0xae, 0x9e, 0x6b, 0x57, 0xfe, 0xd5, 0xab, 0xff, 0x51, 0x27, 0xad, 0xd9, 0x26, 0xd3, 0x44,
	0xab, 0x85, 0x14, 0x74, 0x4c, 0x1a, 0xa5, 0xc9, 0xc2, 0x5a, 0xaf, 0x76, 0xd5, 0x1a, 0x76, 0x59,
	0x95, 0x81, 0x6d, 0x33, 0xb0, 0x19, 0x1a, 0xa9, 0xc4, 0x0b, 0xcf, 0x4a, 0xb8, 0x0f, 0xbe, 0x3e,
	0xcf, 0xeb, 0xbd, 0x5a, 0xbc, 0x01, 0xe8, 0x0d, 0x09, 0xe6, 0xb0, 0x89, 0x1f, 0xd6, 0xff, 0x46,
	0x63, 0xaf, 0xa5, 0x8f, 0xa4, 0x39, 0x97, 0x39, 0x28, 0x2b, 0xb5, 0x0a, 0x1b, 0x0e, 0xbc, 0x64,
	0x87, 0x5a, 0x61, 0x0f, 0x5b, 0x79, 0xbc, 0x27, 0xe9, 0x98, 0x9c, 0x24, 0x12, 0x0d, 0x47, 0x08,
	0x8f, 0x7e, 0x99, 0xfe, 0x3c, 0x55, 0x38, 0x1a, 0x56, 0xd3, 0xb7, 0x62, 0x7a, 0x4b, 0x9a, 0x0b,
	0xe3, 0x7c, 0x11, 0xc2, 0xe3, 0x7f, 0x90, 0x7b, 0x79, 0xff, 0x95, 0xb4, 0x5d, 0x6f, 0x31, 0xbc,
	0x97, 0x60, 0x91, 0x4e, 0x49, 0x90, 0xba, 0x0a, 0x7d, 0x77, 0xd7, 0x87, 0xf7, 0xf8, 0xd1, 0xf9,
	0xae, 0x48, 0x6f, 0xd0, 0x7f, 0x22, 0xa7, 0xde, 0xda, 0x16, 0x5a, 0x59, 0xa0, 0x77, 0x24, 0xa8,
	0x30, 0xef, 0x7d, 0x71, 0xd8, 0x7b, 0xe2, 0x8e, 0xd8, 0x33, 0x49, 0xe0, 0x56, 0x19, 0x7d, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xd4, 0xef, 0x11, 0x12, 0x7c, 0x02, 0x00, 0x00,
}
