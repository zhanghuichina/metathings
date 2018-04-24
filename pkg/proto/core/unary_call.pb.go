// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unary_call.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf2 "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UnaryCallRequestPayload struct {
	ServiceName *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	MethodName  *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	Payload     *google_protobuf2.Any        `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
}

func (m *UnaryCallRequestPayload) Reset()                    { *m = UnaryCallRequestPayload{} }
func (m *UnaryCallRequestPayload) String() string            { return proto.CompactTextString(m) }
func (*UnaryCallRequestPayload) ProtoMessage()               {}
func (*UnaryCallRequestPayload) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *UnaryCallRequestPayload) GetServiceName() *google_protobuf.StringValue {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *UnaryCallRequestPayload) GetMethodName() *google_protobuf.StringValue {
	if m != nil {
		return m.MethodName
	}
	return nil
}

func (m *UnaryCallRequestPayload) GetPayload() *google_protobuf2.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type UnaryCallResponsePayload struct {
	ServiceName string                `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	MethodName  string                `protobuf:"bytes,2,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	Payload     *google_protobuf2.Any `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
}

func (m *UnaryCallResponsePayload) Reset()                    { *m = UnaryCallResponsePayload{} }
func (m *UnaryCallResponsePayload) String() string            { return proto.CompactTextString(m) }
func (*UnaryCallResponsePayload) ProtoMessage()               {}
func (*UnaryCallResponsePayload) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *UnaryCallResponsePayload) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *UnaryCallResponsePayload) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *UnaryCallResponsePayload) GetPayload() *google_protobuf2.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*UnaryCallRequestPayload)(nil), "ai.metathings.service.core.UnaryCallRequestPayload")
	proto.RegisterType((*UnaryCallResponsePayload)(nil), "ai.metathings.service.core.UnaryCallResponsePayload")
}

func init() { proto.RegisterFile("unary_call.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x8f, 0xcb, 0x4a, 0xc4, 0x40,
	0x10, 0x45, 0x89, 0x82, 0x32, 0x1d, 0x17, 0x12, 0x04, 0x63, 0x10, 0x1f, 0xb3, 0x72, 0x55, 0x03,
	0xba, 0x16, 0x11, 0xf7, 0x22, 0x11, 0xdd, 0x0e, 0x35, 0x99, 0x32, 0x13, 0xe8, 0x74, 0xb5, 0xfd,
	0x50, 0xfa, 0x27, 0xfc, 0x2f, 0xff, 0x4a, 0xec, 0x9e, 0x80, 0x1a, 0x17, 0x82, 0xdb, 0xaa, 0x3a,
	0x75, 0xcf, 0x15, 0xbb, 0x5e, 0xa1, 0x09, 0xf3, 0x06, 0xa5, 0x04, 0x6d, 0xd8, 0x71, 0x51, 0x61,
	0x07, 0x3d, 0x39, 0x74, 0xab, 0x4e, 0xb5, 0x16, 0x2c, 0x99, 0x97, 0xae, 0x21, 0x68, 0xd8, 0x50,
	0x75, 0xd4, 0x32, 0xb7, 0x92, 0x66, 0xf1, 0x72, 0xe1, 0x9f, 0x66, 0xaf, 0x06, 0xb5, 0x26, 0x63,
	0x13, 0x5b, 0x1d, 0xfc, 0xdc, 0xa3, 0x0a, 0x69, 0x35, 0x7d, 0xcf, 0xc4, 0xfe, 0xc3, 0x67, 0xd6,
	0x0d, 0x4a, 0x59, 0xd3, 0xb3, 0x27, 0xeb, 0xee, 0x30, 0x48, 0xc6, 0x65, 0x71, 0x25, 0x76, 0xd6,
	0x31, 0x73, 0x85, 0x3d, 0x95, 0xd9, 0x49, 0x76, 0x96, 0x9f, 0x1f, 0x42, 0xfa, 0x06, 0xc3, 0x37,
	0xb8, 0x77, 0xa6, 0x53, 0xed, 0x23, 0x4a, 0x4f, 0x75, 0xbe, 0x26, 0x6e, 0xb1, 0xa7, 0xe2, 0x52,
	0xe4, 0x3d, 0xb9, 0x15, 0x2f, 0x13, 0xbf, 0xf1, 0x07, 0x5e, 0x24, 0x20, 0xe2, 0x20, 0xb6, 0x75,
	0x52, 0x29, 0x37, 0x23, 0xba, 0x37, 0x42, 0xaf, 0x55, 0xa8, 0x87, 0xa3, 0xe9, 0x5b, 0x26, 0xca,
	0x2f, 0x5d, 0xac, 0x66, 0x65, 0x69, 0x28, 0x73, 0xfa, 0x4b, 0x99, 0xc9, 0x77, 0xdd, 0xe3, 0xb1,
	0xee, 0xe4, 0x3f, 0x42, 0x8b, 0xad, 0x38, 0xbe, 0xf8, 0x08, 0x00, 0x00, 0xff, 0xff, 0x35, 0x8d,
	0xb1, 0xa5, 0xce, 0x01, 0x00, 0x00,
}
