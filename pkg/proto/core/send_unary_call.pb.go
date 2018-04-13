// Code generated by protoc-gen-go. DO NOT EDIT.
// source: send_unary_call.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SendUnaryCallRequest struct {
	CoreId  *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=core_id,json=coreId" json:"core_id,omitempty"`
	Payload *UnaryCallPayload            `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
}

func (m *SendUnaryCallRequest) Reset()                    { *m = SendUnaryCallRequest{} }
func (m *SendUnaryCallRequest) String() string            { return proto.CompactTextString(m) }
func (*SendUnaryCallRequest) ProtoMessage()               {}
func (*SendUnaryCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *SendUnaryCallRequest) GetCoreId() *google_protobuf.StringValue {
	if m != nil {
		return m.CoreId
	}
	return nil
}

func (m *SendUnaryCallRequest) GetPayload() *UnaryCallPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type SendUnaryCallResponse struct {
	Payload *UnaryCallPayload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *SendUnaryCallResponse) Reset()                    { *m = SendUnaryCallResponse{} }
func (m *SendUnaryCallResponse) String() string            { return proto.CompactTextString(m) }
func (*SendUnaryCallResponse) ProtoMessage()               {}
func (*SendUnaryCallResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *SendUnaryCallResponse) GetPayload() *UnaryCallPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*SendUnaryCallRequest)(nil), "ai.metathings.service.core.SendUnaryCallRequest")
	proto.RegisterType((*SendUnaryCallResponse)(nil), "ai.metathings.service.core.SendUnaryCallResponse")
}

func init() { proto.RegisterFile("send_unary_call.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0x80, 0x49, 0x0f, 0x29, 0xe4, 0xbf, 0xfc, 0x04, 0x0b, 0x25, 0x88, 0x96, 0x9e, 0x3c, 0xd8,
	0x09, 0x28, 0xfa, 0x00, 0x0a, 0x82, 0x17, 0x91, 0x14, 0xbd, 0x86, 0x49, 0x76, 0xdc, 0x2e, 0x6e,
	0x76, 0xe2, 0xee, 0xa6, 0xa1, 0x0f, 0xe3, 0xb3, 0x09, 0x3e, 0x89, 0x24, 0xb1, 0x11, 0x0a, 0x1e,
	0xbc, 0x2d, 0xcc, 0x7c, 0xfb, 0x7d, 0x13, 0xcd, 0x1c, 0x19, 0x91, 0x37, 0x06, 0xed, 0x2e, 0x2f,
	0x51, 0x6b, 0xa8, 0x2d, 0x7b, 0x8e, 0x13, 0x54, 0x50, 0x91, 0x47, 0xbf, 0x51, 0x46, 0x3a, 0x70,
	0x64, 0xb7, 0xaa, 0x24, 0x28, 0xd9, 0x52, 0x72, 0x22, 0x99, 0xa5, 0xa6, 0xb4, 0xdf, 0x2c, 0x9a,
	0x97, 0xb4, 0xb5, 0x58, 0xd7, 0x64, 0xdd, 0xc0, 0x26, 0xd7, 0x52, 0xf9, 0x4d, 0x53, 0x40, 0xc9,
	0x55, 0x5a, 0xb5, 0xca, 0xbf, 0x72, 0x9b, 0x4a, 0x5e, 0xf5, 0xc3, 0xd5, 0x16, 0xb5, 0x12, 0xe8,
	0xd9, 0xba, 0x74, 0x7c, 0x7e, 0x73, 0xff, 0x0f, 0x2b, 0x96, 0xef, 0x41, 0x74, 0xb4, 0x26, 0x23,
	0x9e, 0xba, 0xc1, 0x2d, 0x6a, 0x9d, 0xd1, 0x5b, 0x43, 0xce, 0xc7, 0x57, 0xd1, 0xb4, 0x4b, 0xc9,
	0x95, 0x98, 0x07, 0x8b, 0xe0, 0xec, 0xdf, 0xc5, 0x31, 0x0c, 0x51, 0xb0, 0x8f, 0x82, 0xb5, 0xb7,
	0xca, 0xc8, 0x67, 0xd4, 0x0d, 0x65, 0x61, 0xb7, 0x7c, 0x2f, 0xe2, 0x87, 0x68, 0x5a, 0xe3, 0x4e,
	0x33, 0x8a, 0xf9, 0xa4, 0xc7, 0xce, 0xe1, 0xf7, 0x3b, 0x61, 0xb4, 0x3e, 0x0e, 0xcc, 0x4d, 0xf8,
	0xf9, 0x71, 0x3a, 0x59, 0x04, 0xd9, 0xfe, 0x93, 0x65, 0x1e, 0xcd, 0x0e, 0xf2, 0x5c, 0xcd, 0xc6,
	0x51, 0x7c, 0xf7, 0x23, 0x0a, 0xfe, 0x2e, 0x1a, 0x05, 0x45, 0xd8, 0x9f, 0x73, 0xf9, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x1f, 0xdb, 0xd8, 0xcd, 0xa6, 0x01, 0x00, 0x00,
}