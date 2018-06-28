// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat.proto

package cored

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type HeartbeatEntity struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	HeartbeatAt          *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=heartbeat_at,json=heartbeatAt" json:"heartbeat_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HeartbeatEntity) Reset()         { *m = HeartbeatEntity{} }
func (m *HeartbeatEntity) String() string { return proto.CompactTextString(m) }
func (*HeartbeatEntity) ProtoMessage()    {}
func (*HeartbeatEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_477d8f1e7fbd4fff, []int{0}
}
func (m *HeartbeatEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatEntity.Unmarshal(m, b)
}
func (m *HeartbeatEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatEntity.Marshal(b, m, deterministic)
}
func (dst *HeartbeatEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatEntity.Merge(dst, src)
}
func (m *HeartbeatEntity) XXX_Size() int {
	return xxx_messageInfo_HeartbeatEntity.Size(m)
}
func (m *HeartbeatEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatEntity.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatEntity proto.InternalMessageInfo

func (m *HeartbeatEntity) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *HeartbeatEntity) GetHeartbeatAt() *timestamp.Timestamp {
	if m != nil {
		return m.HeartbeatAt
	}
	return nil
}

type HeartbeatRequest struct {
	Session              *wrappers.UInt64Value `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	Entities             []*HeartbeatEntity    `protobuf:"bytes,2,rep,name=entities" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_477d8f1e7fbd4fff, []int{1}
}
func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (dst *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(dst, src)
}
func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}
func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

func (m *HeartbeatRequest) GetSession() *wrappers.UInt64Value {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *HeartbeatRequest) GetEntities() []*HeartbeatEntity {
	if m != nil {
		return m.Entities
	}
	return nil
}

func init() {
	proto.RegisterType((*HeartbeatEntity)(nil), "ai.metathings.service.cored.HeartbeatEntity")
	proto.RegisterType((*HeartbeatRequest)(nil), "ai.metathings.service.cored.HeartbeatRequest")
}

func init() { proto.RegisterFile("heartbeat.proto", fileDescriptor_heartbeat_477d8f1e7fbd4fff) }

var fileDescriptor_heartbeat_477d8f1e7fbd4fff = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x84, 0x2a, 0x1b, 0xa1, 0x92, 0x53, 0x88, 0x62, 0x43, 0x4f, 0x3d, 0xb4, 0x1b,
	0xa8, 0xd2, 0x9b, 0x82, 0x82, 0x50, 0xaf, 0xf1, 0xcf, 0x55, 0x36, 0xc9, 0xb8, 0x19, 0x4c, 0xb2,
	0x71, 0x77, 0xd2, 0xe0, 0xc5, 0x27, 0xf1, 0xdd, 0x04, 0x9f, 0x44, 0x4c, 0x4c, 0x0e, 0x15, 0x7b,
	0x1b, 0x98, 0xef, 0xc7, 0xf7, 0x9b, 0x61, 0xe3, 0x0c, 0x84, 0xa6, 0x18, 0x04, 0xf1, 0x4a, 0x2b,
	0x52, 0xee, 0xb1, 0x40, 0x5e, 0x00, 0x09, 0xca, 0xb0, 0x94, 0x86, 0x1b, 0xd0, 0x1b, 0x4c, 0x80,
	0x27, 0x4a, 0x43, 0xea, 0x9f, 0x4a, 0xa5, 0x64, 0x0e, 0x61, 0x1b, 0x8d, 0xeb, 0xe7, 0xb0, 0xd1,
	0xa2, 0xaa, 0x40, 0x9b, 0x0e, 0xf6, 0x27, 0xdb, 0x7b, 0xc2, 0x02, 0x0c, 0x89, 0xa2, 0xfa, 0x0d,
	0xac, 0x24, 0x52, 0x56, 0xc7, 0x3c, 0x51, 0x45, 0x58, 0x34, 0x48, 0x2f, 0xaa, 0x09, 0xa5, 0x5a,
	0xb4, 0xcb, 0xc5, 0x46, 0xe4, 0x98, 0x0a, 0x52, 0xda, 0x84, 0xc3, 0xd8, 0x71, 0xd3, 0x77, 0x36,
	0x5e, 0xf7, 0xa2, 0x37, 0x25, 0x21, 0xbd, 0xb9, 0x73, 0x66, 0x63, 0xea, 0x59, 0x81, 0x35, 0x73,
	0x96, 0x27, 0xbc, 0x2b, 0xe6, 0x7d, 0x31, 0xbf, 0x23, 0x8d, 0xa5, 0x7c, 0x14, 0x79, 0x0d, 0x91,
	0x8d, 0xa9, 0x7b, 0xc1, 0x0e, 0x87, 0x4b, 0x9f, 0x04, 0x79, 0x76, 0xcb, 0xf9, 0x7f, 0xb8, 0xfb,
	0x5e, 0x38, 0x72, 0x86, 0xfc, 0x15, 0x4d, 0x3f, 0x2c, 0x76, 0x34, 0x08, 0x44, 0xf0, 0x5a, 0x83,
	0x21, 0xf7, 0x92, 0xed, 0x1b, 0x30, 0x06, 0x55, 0xf9, 0xaf, 0xc6, 0xc3, 0x6d, 0x49, 0xab, 0xf3,
	0x56, 0xe3, 0x7a, 0xf4, 0xf5, 0x39, 0xb1, 0x03, 0x2b, 0xea, 0x21, 0x77, 0xcd, 0x0e, 0xe0, 0xe7,
	0x16, 0x04, 0xe3, 0xd9, 0xc1, 0xde, 0xcc, 0x59, 0xce, 0xf9, 0x8e, 0xef, 0xf3, 0xad, 0x0f, 0x44,
	0x03, 0x1d, 0x8f, 0xda, 0xc2, 0xb3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x2f, 0x55, 0xdb,
	0xce, 0x01, 0x00, 0x00,
}
