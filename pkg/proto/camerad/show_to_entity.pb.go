// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show_to_entity.proto

package camerad

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
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

type ShowToEntityResponse struct {
	Camera               *Camera  `protobuf:"bytes,1,opt,name=camera" json:"camera,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowToEntityResponse) Reset()         { *m = ShowToEntityResponse{} }
func (m *ShowToEntityResponse) String() string { return proto.CompactTextString(m) }
func (*ShowToEntityResponse) ProtoMessage()    {}
func (*ShowToEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_show_to_entity_9c533d9f88622c2a, []int{0}
}
func (m *ShowToEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowToEntityResponse.Unmarshal(m, b)
}
func (m *ShowToEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowToEntityResponse.Marshal(b, m, deterministic)
}
func (dst *ShowToEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowToEntityResponse.Merge(dst, src)
}
func (m *ShowToEntityResponse) XXX_Size() int {
	return xxx_messageInfo_ShowToEntityResponse.Size(m)
}
func (m *ShowToEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowToEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowToEntityResponse proto.InternalMessageInfo

func (m *ShowToEntityResponse) GetCamera() *Camera {
	if m != nil {
		return m.Camera
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowToEntityResponse)(nil), "ai.metathings.service.camerad.ShowToEntityResponse")
}

func init() {
	proto.RegisterFile("show_to_entity.proto", fileDescriptor_show_to_entity_9c533d9f88622c2a)
}

var fileDescriptor_show_to_entity_9c533d9f88622c2a = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0xf1, 0xe2, 0xc1, 0xed, 0x64, 0x3c, 0x14, 0x43, 0x4b, 0x29, 0x14, 0xba, 0x58, 0x82,
	0x16, 0xba, 0x65, 0x0a, 0x79, 0x01, 0x27, 0x99, 0x8d, 0x6c, 0x5f, 0x24, 0x11, 0xcb, 0x27, 0xa4,
	0xb3, 0x45, 0xde, 0x3e, 0x20, 0x99, 0x8c, 0x99, 0xee, 0x8e, 0xfb, 0xbf, 0xbb, 0xaf, 0xa8, 0xbc,
	0xc2, 0xd0, 0x11, 0x76, 0x30, 0x93, 0xa6, 0x1b, 0xb3, 0x0e, 0x09, 0xcb, 0x77, 0xa1, 0x99, 0x01,
	0x12, 0xa4, 0xf4, 0x2c, 0x3d, 0xf3, 0xe0, 0x56, 0x3d, 0x00, 0x1b, 0x84, 0x01, 0x27, 0xc6, 0xfa,
	0x43, 0x22, 0xca, 0x09, 0x78, 0x0c, 0xf7, 0xcb, 0x85, 0x07, 0x27, 0xac, 0x05, 0xe7, 0x13, 0x5e,
	0xff, 0x4b, 0x4d, 0x6a, 0xe9, 0xd9, 0x80, 0x86, 0x9b, 0xa0, 0xe9, 0x8a, 0x81, 0x4b, 0x6c, 0xe2,
	0xb2, 0x59, 0xc5, 0xa4, 0x47, 0x41, 0xe8, 0x3c, 0x7f, 0xb4, 0x1b, 0xf7, 0x9a, 0x1e, 0xa4, 0xe9,
	0xeb, 0x5c, 0x54, 0x47, 0x85, 0xe1, 0x84, 0x87, 0xa8, 0xd6, 0x82, 0xb7, 0x38, 0x7b, 0x28, 0x77,
	0x45, 0x9e, 0x72, 0x6f, 0xd9, 0x67, 0xf6, 0xf3, 0xf2, 0xfb, 0xcd, 0x9e, 0xda, 0xb2, 0x7d, 0xac,
	0xed, 0x06, 0xf5, 0x79, 0xbc, 0xfe, 0x77, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xbc, 0x1e, 0x9e,
	0xfa, 0x00, 0x00, 0x00,
}
