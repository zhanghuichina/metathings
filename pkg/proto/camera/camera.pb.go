// Code generated by protoc-gen-go. DO NOT EDIT.
// source: camera.proto

package camera

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CameraState int32

const (
	CameraState_CAMERA_STATE_UNKNOWN     CameraState = 0
	CameraState_CAMERA_STATE_STOP        CameraState = 1
	CameraState_CAMERA_STATE_STARTING    CameraState = 2
	CameraState_CAMERA_STATE_TERMINATING CameraState = 3
	CameraState_CAMERA_STATE_RUNNING     CameraState = 4
)

var CameraState_name = map[int32]string{
	0: "CAMERA_STATE_UNKNOWN",
	1: "CAMERA_STATE_STOP",
	2: "CAMERA_STATE_STARTING",
	3: "CAMERA_STATE_TERMINATING",
	4: "CAMERA_STATE_RUNNING",
}
var CameraState_value = map[string]int32{
	"CAMERA_STATE_UNKNOWN":     0,
	"CAMERA_STATE_STOP":        1,
	"CAMERA_STATE_STARTING":    2,
	"CAMERA_STATE_TERMINATING": 3,
	"CAMERA_STATE_RUNNING":     4,
}

func (x CameraState) String() string {
	return proto.EnumName(CameraState_name, int32(x))
}
func (CameraState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_camera_953e8c915faee521, []int{0}
}

type CameraConfig struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Device               string   `protobuf:"bytes,2,opt,name=device" json:"device,omitempty"`
	Width                uint32   `protobuf:"varint,3,opt,name=width" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	Bitrate              uint32   `protobuf:"varint,5,opt,name=bitrate" json:"bitrate,omitempty"`
	Framerate            uint32   `protobuf:"varint,6,opt,name=framerate" json:"framerate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CameraConfig) Reset()         { *m = CameraConfig{} }
func (m *CameraConfig) String() string { return proto.CompactTextString(m) }
func (*CameraConfig) ProtoMessage()    {}
func (*CameraConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_camera_953e8c915faee521, []int{0}
}
func (m *CameraConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CameraConfig.Unmarshal(m, b)
}
func (m *CameraConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CameraConfig.Marshal(b, m, deterministic)
}
func (dst *CameraConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CameraConfig.Merge(dst, src)
}
func (m *CameraConfig) XXX_Size() int {
	return xxx_messageInfo_CameraConfig.Size(m)
}
func (m *CameraConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CameraConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CameraConfig proto.InternalMessageInfo

func (m *CameraConfig) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CameraConfig) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *CameraConfig) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *CameraConfig) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *CameraConfig) GetBitrate() uint32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

func (m *CameraConfig) GetFramerate() uint32 {
	if m != nil {
		return m.Framerate
	}
	return 0
}

type Camera struct {
	State                CameraState   `protobuf:"varint,1,opt,name=state,enum=ai.metathings.service.camera.CameraState" json:"state,omitempty"`
	Config               *CameraConfig `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Camera) Reset()         { *m = Camera{} }
func (m *Camera) String() string { return proto.CompactTextString(m) }
func (*Camera) ProtoMessage()    {}
func (*Camera) Descriptor() ([]byte, []int) {
	return fileDescriptor_camera_953e8c915faee521, []int{1}
}
func (m *Camera) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Camera.Unmarshal(m, b)
}
func (m *Camera) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Camera.Marshal(b, m, deterministic)
}
func (dst *Camera) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Camera.Merge(dst, src)
}
func (m *Camera) XXX_Size() int {
	return xxx_messageInfo_Camera.Size(m)
}
func (m *Camera) XXX_DiscardUnknown() {
	xxx_messageInfo_Camera.DiscardUnknown(m)
}

var xxx_messageInfo_Camera proto.InternalMessageInfo

func (m *Camera) GetState() CameraState {
	if m != nil {
		return m.State
	}
	return CameraState_CAMERA_STATE_UNKNOWN
}

func (m *Camera) GetConfig() *CameraConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*CameraConfig)(nil), "ai.metathings.service.camera.CameraConfig")
	proto.RegisterType((*Camera)(nil), "ai.metathings.service.camera.Camera")
	proto.RegisterEnum("ai.metathings.service.camera.CameraState", CameraState_name, CameraState_value)
}

func init() { proto.RegisterFile("camera.proto", fileDescriptor_camera_953e8c915faee521) }

var fileDescriptor_camera_953e8c915faee521 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4f, 0x32, 0x31,
	0x10, 0xc6, 0xdf, 0xf2, 0x67, 0xdf, 0x30, 0xa0, 0x59, 0x27, 0x60, 0x6a, 0xc2, 0x81, 0x70, 0x42,
	0x0e, 0x3d, 0xe0, 0x07, 0x30, 0x2b, 0x21, 0x86, 0x18, 0x8a, 0x29, 0x4b, 0x3c, 0x92, 0x05, 0x0a,
	0xdb, 0x44, 0x58, 0xd3, 0xad, 0xfa, 0x29, 0x3c, 0xf8, 0x05, 0xfc, 0xac, 0xa6, 0xed, 0x1a, 0x45,
	0x13, 0xe3, 0x6d, 0x9e, 0xe7, 0x37, 0xcf, 0xec, 0xcc, 0x16, 0x1a, 0xab, 0x64, 0x27, 0x75, 0xc2,
	0x1e, 0x74, 0x66, 0x32, 0x6c, 0x27, 0x8a, 0xed, 0xa4, 0x49, 0x4c, 0xaa, 0xf6, 0xdb, 0x9c, 0xe5,
	0x52, 0x3f, 0xa9, 0x95, 0x64, 0xbe, 0xa7, 0xfb, 0x46, 0xa0, 0x31, 0x74, 0xe5, 0x30, 0xdb, 0x6f,
	0xd4, 0x16, 0x43, 0x28, 0x3f, 0xea, 0x7b, 0x4a, 0x3a, 0xa4, 0x57, 0x13, 0xb6, 0xc4, 0x53, 0x08,
	0xd6, 0xd2, 0x66, 0x68, 0xc9, 0x99, 0x85, 0xc2, 0x26, 0x54, 0x9f, 0xd5, 0xda, 0xa4, 0xb4, 0xdc,
	0x21, 0xbd, 0x23, 0xe1, 0x85, 0xed, 0x4e, 0xa5, 0xda, 0xa6, 0x86, 0x56, 0x9c, 0x5d, 0x28, 0xa4,
	0xf0, 0x7f, 0xa9, 0x8c, 0x4e, 0x8c, 0xa4, 0x55, 0x07, 0x3e, 0x24, 0xb6, 0xa1, 0xb6, 0xd1, 0x6e,
	0x05, 0x23, 0x69, 0xe0, 0xd8, 0xa7, 0xd1, 0x7d, 0x21, 0x10, 0xf8, 0x05, 0xf1, 0x12, 0xaa, 0xb9,
	0xb1, 0x4d, 0x76, 0xb9, 0xe3, 0xc1, 0x39, 0xfb, 0xed, 0x32, 0xe6, 0x43, 0x33, 0x1b, 0x10, 0x3e,
	0x87, 0x57, 0x10, 0xac, 0xdc, 0x95, 0xee, 0x92, 0xfa, 0xa0, 0xff, 0x97, 0x09, 0xfe, 0xbf, 0x88,
	0x22, 0xd9, 0x7f, 0x25, 0x50, 0xff, 0x32, 0x1a, 0x29, 0x34, 0x87, 0xd1, 0x64, 0x24, 0xa2, 0xc5,
	0x2c, 0x8e, 0xe2, 0xd1, 0x62, 0xce, 0x6f, 0xf8, 0xf4, 0x8e, 0x87, 0xff, 0xb0, 0x05, 0x27, 0x07,
	0x64, 0x16, 0x4f, 0x6f, 0x43, 0x82, 0x67, 0xd0, 0xfa, 0x66, 0x47, 0x22, 0x1e, 0xf3, 0xeb, 0xb0,
	0x84, 0x6d, 0xa0, 0x07, 0x28, 0x1e, 0x89, 0xc9, 0x98, 0x47, 0x8e, 0x96, 0x7f, 0x7c, 0x49, 0xcc,
	0x39, 0xb7, 0xa4, 0xb2, 0x0c, 0xdc, 0x4b, 0x5f, 0xbc, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xcd,
	0xbf, 0x8d, 0xf9, 0x01, 0x00, 0x00,
}
