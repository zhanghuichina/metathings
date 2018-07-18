// Code generated by protoc-gen-go. DO NOT EDIT.
// source: publish.proto

package sensord

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import sensor "github.com/nayotta/metathings/pkg/proto/sensor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PublishDataRequest struct {
	Data                 *sensor.SensorData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PublishDataRequest) Reset()         { *m = PublishDataRequest{} }
func (m *PublishDataRequest) String() string { return proto.CompactTextString(m) }
func (*PublishDataRequest) ProtoMessage()    {}
func (*PublishDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_40f882461a3960c2, []int{0}
}
func (m *PublishDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishDataRequest.Unmarshal(m, b)
}
func (m *PublishDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishDataRequest.Marshal(b, m, deterministic)
}
func (dst *PublishDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishDataRequest.Merge(dst, src)
}
func (m *PublishDataRequest) XXX_Size() int {
	return xxx_messageInfo_PublishDataRequest.Size(m)
}
func (m *PublishDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishDataRequest proto.InternalMessageInfo

func (m *PublishDataRequest) GetData() *sensor.SensorData {
	if m != nil {
		return m.Data
	}
	return nil
}

type PublishRequest struct {
	Session *wrappers.UInt64Value `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*PublishRequest_Data
	Payload              isPublishRequest_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_40f882461a3960c2, []int{1}
}
func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (dst *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(dst, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

type isPublishRequest_Payload interface {
	isPublishRequest_Payload()
}

type PublishRequest_Data struct {
	Data *PublishDataRequest `protobuf:"bytes,2,opt,name=data,oneof"`
}

func (*PublishRequest_Data) isPublishRequest_Payload() {}

func (m *PublishRequest) GetPayload() isPublishRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PublishRequest) GetSession() *wrappers.UInt64Value {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *PublishRequest) GetData() *PublishDataRequest {
	if x, ok := m.GetPayload().(*PublishRequest_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublishRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublishRequest_OneofMarshaler, _PublishRequest_OneofUnmarshaler, _PublishRequest_OneofSizer, []interface{}{
		(*PublishRequest_Data)(nil),
	}
}

func _PublishRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublishRequest)
	// payload
	switch x := m.Payload.(type) {
	case *PublishRequest_Data:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PublishRequest.Payload has unexpected type %T", x)
	}
	return nil
}

func _PublishRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublishRequest)
	switch tag {
	case 2: // payload.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PublishDataRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &PublishRequest_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PublishRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublishRequest)
	// payload
	switch x := m.Payload.(type) {
	case *PublishRequest_Data:
		s := proto.Size(x.Data)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PublishRequests struct {
	Requests             []*PublishRequest `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PublishRequests) Reset()         { *m = PublishRequests{} }
func (m *PublishRequests) String() string { return proto.CompactTextString(m) }
func (*PublishRequests) ProtoMessage()    {}
func (*PublishRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_40f882461a3960c2, []int{2}
}
func (m *PublishRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequests.Unmarshal(m, b)
}
func (m *PublishRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequests.Marshal(b, m, deterministic)
}
func (dst *PublishRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequests.Merge(dst, src)
}
func (m *PublishRequests) XXX_Size() int {
	return xxx_messageInfo_PublishRequests.Size(m)
}
func (m *PublishRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequests.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequests proto.InternalMessageInfo

func (m *PublishRequests) GetRequests() []*PublishRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

type PublishResponse struct {
	Session              uint64   `protobuf:"varint,1,opt,name=session" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_40f882461a3960c2, []int{3}
}
func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponse.Unmarshal(m, b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
}
func (dst *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(dst, src)
}
func (m *PublishResponse) XXX_Size() int {
	return xxx_messageInfo_PublishResponse.Size(m)
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

func (m *PublishResponse) GetSession() uint64 {
	if m != nil {
		return m.Session
	}
	return 0
}

type PublishResponses struct {
	Responses            []*PublishResponse `protobuf:"bytes,1,rep,name=responses" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PublishResponses) Reset()         { *m = PublishResponses{} }
func (m *PublishResponses) String() string { return proto.CompactTextString(m) }
func (*PublishResponses) ProtoMessage()    {}
func (*PublishResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_publish_40f882461a3960c2, []int{4}
}
func (m *PublishResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponses.Unmarshal(m, b)
}
func (m *PublishResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponses.Marshal(b, m, deterministic)
}
func (dst *PublishResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponses.Merge(dst, src)
}
func (m *PublishResponses) XXX_Size() int {
	return xxx_messageInfo_PublishResponses.Size(m)
}
func (m *PublishResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponses.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponses proto.InternalMessageInfo

func (m *PublishResponses) GetResponses() []*PublishResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*PublishDataRequest)(nil), "ai.metathings.service.sensord.PublishDataRequest")
	proto.RegisterType((*PublishRequest)(nil), "ai.metathings.service.sensord.PublishRequest")
	proto.RegisterType((*PublishRequests)(nil), "ai.metathings.service.sensord.PublishRequests")
	proto.RegisterType((*PublishResponse)(nil), "ai.metathings.service.sensord.PublishResponse")
	proto.RegisterType((*PublishResponses)(nil), "ai.metathings.service.sensord.PublishResponses")
}

func init() { proto.RegisterFile("publish.proto", fileDescriptor_publish_40f882461a3960c2) }

var fileDescriptor_publish_40f882461a3960c2 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x7f, 0xdd, 0x6f, 0x6c, 0x2e, 0xf3, 0x1f, 0x3d, 0x8d, 0xe1, 0x9f, 0xd1, 0xd3, 0x40,
	0x96, 0xe0, 0x94, 0x5d, 0x04, 0x0f, 0x63, 0xa0, 0x03, 0x0f, 0x52, 0xd1, 0x83, 0x78, 0xf0, 0xe9,
	0x1a, 0xbb, 0xb0, 0xae, 0xa9, 0x79, 0xd2, 0x8d, 0xbd, 0x1f, 0xdf, 0x97, 0xe0, 0x2b, 0x11, 0xdb,
	0xb4, 0xdb, 0x14, 0xc6, 0x4e, 0x49, 0x48, 0x3e, 0xdf, 0x7c, 0x9e, 0x2f, 0xd9, 0x8b, 0x13, 0x2f,
	0x14, 0x38, 0xa6, 0xb1, 0x92, 0x5a, 0xda, 0xc7, 0x20, 0xe8, 0x94, 0x6b, 0xd0, 0x63, 0x11, 0x05,
	0x48, 0x91, 0xab, 0x99, 0x18, 0x71, 0x8a, 0x3c, 0x42, 0xa9, 0xfc, 0xe6, 0x49, 0x20, 0x65, 0x10,
	0x72, 0x96, 0x3e, 0xf6, 0x92, 0x37, 0x36, 0x57, 0x10, 0xc7, 0x5c, 0x61, 0x86, 0x37, 0x7b, 0x81,
	0xd0, 0xe3, 0xc4, 0xa3, 0x23, 0x39, 0x65, 0xd3, 0xb9, 0xd0, 0x13, 0x39, 0x67, 0x81, 0xec, 0xa4,
	0x97, 0x9d, 0x19, 0x84, 0xc2, 0x07, 0x2d, 0x15, 0xb2, 0x62, 0x6b, 0xb8, 0xab, 0x15, 0x2e, 0x82,
	0x85, 0xd4, 0x1a, 0xd8, 0x52, 0x83, 0xc5, 0x93, 0x20, 0xfb, 0x92, 0x65, 0x22, 0x66, 0x31, 0xf0,
	0xee, 0xea, 0xc9, 0x79, 0x26, 0xf6, 0x7d, 0x36, 0xd2, 0x00, 0x34, 0xb8, 0xfc, 0x3d, 0xe1, 0xa8,
	0xed, 0x01, 0x29, 0xfb, 0xa0, 0xa1, 0x61, 0xb5, 0xac, 0x76, 0xbd, 0xdb, 0xa6, 0x9b, 0xc6, 0xa4,
	0x0f, 0xe9, 0xf2, 0x83, 0xf7, 0x2b, 0x5f, 0x9f, 0xa7, 0xa5, 0x96, 0xe5, 0xa6, 0xb4, 0xf3, 0x61,
	0x91, 0x7d, 0x13, 0x9e, 0x07, 0x5f, 0x93, 0x2a, 0x72, 0x44, 0x21, 0x23, 0x93, 0x7d, 0x44, 0xb3,
	0x8e, 0x68, 0xde, 0x11, 0x7d, 0x1c, 0x46, 0xba, 0x77, 0xf9, 0x04, 0x61, 0xc2, 0x8b, 0xbc, 0x1c,
	0xb2, 0x6f, 0x8c, 0x58, 0x29, 0x85, 0xcf, 0x37, 0x8a, 0xf9, 0xf4, 0xef, 0x64, 0xb7, 0xff, 0x32,
	0xb7, 0x7e, 0x8d, 0x54, 0x63, 0x58, 0x84, 0x12, 0x7c, 0xe7, 0x85, 0x1c, 0xac, 0x5b, 0xa2, 0x3d,
	0x24, 0x3b, 0xca, 0xec, 0x1b, 0x56, 0xeb, 0x7f, 0xbb, 0xde, 0xed, 0x6c, 0xf7, 0x95, 0x49, 0x70,
	0x0b, 0xdc, 0x39, 0x5b, 0x49, 0xc7, 0x58, 0x46, 0xc8, 0xed, 0xc6, 0x7a, 0x09, 0xe5, 0x62, 0x3c,
	0xe7, 0x95, 0x1c, 0xfe, 0x7a, 0x8c, 0xf6, 0x1d, 0xa9, 0xa9, 0xfc, 0x60, 0x64, 0xe8, 0xb6, 0x32,
	0x19, 0xe6, 0x2e, 0x03, 0xbc, 0x4a, 0xda, 0xf3, 0xc5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8,
	0xb0, 0x6c, 0x6f, 0xc9, 0x02, 0x00, 0x00,
}