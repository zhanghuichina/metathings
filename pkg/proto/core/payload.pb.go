// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payload.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamErrorResponsePayload struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	MethodName           string   `protobuf:"bytes,3,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	Context              string   `protobuf:"bytes,4,opt,name=context" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamErrorResponsePayload) Reset()         { *m = StreamErrorResponsePayload{} }
func (m *StreamErrorResponsePayload) String() string { return proto.CompactTextString(m) }
func (*StreamErrorResponsePayload) ProtoMessage()    {}
func (*StreamErrorResponsePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_bad64dd5ad3c7973, []int{0}
}
func (m *StreamErrorResponsePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamErrorResponsePayload.Unmarshal(m, b)
}
func (m *StreamErrorResponsePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamErrorResponsePayload.Marshal(b, m, deterministic)
}
func (dst *StreamErrorResponsePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamErrorResponsePayload.Merge(dst, src)
}
func (m *StreamErrorResponsePayload) XXX_Size() int {
	return xxx_messageInfo_StreamErrorResponsePayload.Size(m)
}
func (m *StreamErrorResponsePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamErrorResponsePayload.DiscardUnknown(m)
}

var xxx_messageInfo_StreamErrorResponsePayload proto.InternalMessageInfo

func (m *StreamErrorResponsePayload) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamErrorResponsePayload) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *StreamErrorResponsePayload) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *StreamErrorResponsePayload) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

type UnaryCallRequestPayload struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ServiceName          *wrappers.StringValue `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	MethodName           *wrappers.StringValue `protobuf:"bytes,3,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	Value                *any.Any              `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UnaryCallRequestPayload) Reset()         { *m = UnaryCallRequestPayload{} }
func (m *UnaryCallRequestPayload) String() string { return proto.CompactTextString(m) }
func (*UnaryCallRequestPayload) ProtoMessage()    {}
func (*UnaryCallRequestPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_bad64dd5ad3c7973, []int{1}
}
func (m *UnaryCallRequestPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnaryCallRequestPayload.Unmarshal(m, b)
}
func (m *UnaryCallRequestPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnaryCallRequestPayload.Marshal(b, m, deterministic)
}
func (dst *UnaryCallRequestPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnaryCallRequestPayload.Merge(dst, src)
}
func (m *UnaryCallRequestPayload) XXX_Size() int {
	return xxx_messageInfo_UnaryCallRequestPayload.Size(m)
}
func (m *UnaryCallRequestPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_UnaryCallRequestPayload.DiscardUnknown(m)
}

var xxx_messageInfo_UnaryCallRequestPayload proto.InternalMessageInfo

func (m *UnaryCallRequestPayload) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UnaryCallRequestPayload) GetServiceName() *wrappers.StringValue {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *UnaryCallRequestPayload) GetMethodName() *wrappers.StringValue {
	if m != nil {
		return m.MethodName
	}
	return nil
}

func (m *UnaryCallRequestPayload) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type UnaryCallResponsePayload struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	MethodName           string   `protobuf:"bytes,3,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	Value                *any.Any `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnaryCallResponsePayload) Reset()         { *m = UnaryCallResponsePayload{} }
func (m *UnaryCallResponsePayload) String() string { return proto.CompactTextString(m) }
func (*UnaryCallResponsePayload) ProtoMessage()    {}
func (*UnaryCallResponsePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_bad64dd5ad3c7973, []int{2}
}
func (m *UnaryCallResponsePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnaryCallResponsePayload.Unmarshal(m, b)
}
func (m *UnaryCallResponsePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnaryCallResponsePayload.Marshal(b, m, deterministic)
}
func (dst *UnaryCallResponsePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnaryCallResponsePayload.Merge(dst, src)
}
func (m *UnaryCallResponsePayload) XXX_Size() int {
	return xxx_messageInfo_UnaryCallResponsePayload.Size(m)
}
func (m *UnaryCallResponsePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_UnaryCallResponsePayload.DiscardUnknown(m)
}

var xxx_messageInfo_UnaryCallResponsePayload proto.InternalMessageInfo

func (m *UnaryCallResponsePayload) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

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

func (m *UnaryCallResponsePayload) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type StreamCallRequestPayload struct {
	// Types that are valid to be assigned to Payload:
	//	*StreamCallRequestPayload_Config
	//	*StreamCallRequestPayload_Data
	Payload              isStreamCallRequestPayload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *StreamCallRequestPayload) Reset()         { *m = StreamCallRequestPayload{} }
func (m *StreamCallRequestPayload) String() string { return proto.CompactTextString(m) }
func (*StreamCallRequestPayload) ProtoMessage()    {}
func (*StreamCallRequestPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_bad64dd5ad3c7973, []int{3}
}
func (m *StreamCallRequestPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallRequestPayload.Unmarshal(m, b)
}
func (m *StreamCallRequestPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallRequestPayload.Marshal(b, m, deterministic)
}
func (dst *StreamCallRequestPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallRequestPayload.Merge(dst, src)
}
func (m *StreamCallRequestPayload) XXX_Size() int {
	return xxx_messageInfo_StreamCallRequestPayload.Size(m)
}
func (m *StreamCallRequestPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallRequestPayload.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallRequestPayload proto.InternalMessageInfo

type isStreamCallRequestPayload_Payload interface {
	isStreamCallRequestPayload_Payload()
}

type StreamCallRequestPayload_Config struct {
	Config *StreamCallConfigRequest `protobuf:"bytes,1,opt,name=config,oneof"`
}
type StreamCallRequestPayload_Data struct {
	Data *StreamCallDataRequest `protobuf:"bytes,2,opt,name=data,oneof"`
}

func (*StreamCallRequestPayload_Config) isStreamCallRequestPayload_Payload() {}
func (*StreamCallRequestPayload_Data) isStreamCallRequestPayload_Payload()   {}

func (m *StreamCallRequestPayload) GetPayload() isStreamCallRequestPayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *StreamCallRequestPayload) GetConfig() *StreamCallConfigRequest {
	if x, ok := m.GetPayload().(*StreamCallRequestPayload_Config); ok {
		return x.Config
	}
	return nil
}

func (m *StreamCallRequestPayload) GetData() *StreamCallDataRequest {
	if x, ok := m.GetPayload().(*StreamCallRequestPayload_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamCallRequestPayload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamCallRequestPayload_OneofMarshaler, _StreamCallRequestPayload_OneofUnmarshaler, _StreamCallRequestPayload_OneofSizer, []interface{}{
		(*StreamCallRequestPayload_Config)(nil),
		(*StreamCallRequestPayload_Data)(nil),
	}
}

func _StreamCallRequestPayload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamCallRequestPayload)
	// payload
	switch x := m.Payload.(type) {
	case *StreamCallRequestPayload_Config:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *StreamCallRequestPayload_Data:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamCallRequestPayload.Payload has unexpected type %T", x)
	}
	return nil
}

func _StreamCallRequestPayload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamCallRequestPayload)
	switch tag {
	case 1: // payload.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallConfigRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &StreamCallRequestPayload_Config{msg}
		return true, err
	case 2: // payload.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallDataRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &StreamCallRequestPayload_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamCallRequestPayload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamCallRequestPayload)
	// payload
	switch x := m.Payload.(type) {
	case *StreamCallRequestPayload_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamCallRequestPayload_Data:
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

type StreamCallResponsePayload struct {
	// Types that are valid to be assigned to Payload:
	//	*StreamCallResponsePayload_Config
	//	*StreamCallResponsePayload_Data
	Payload              isStreamCallResponsePayload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *StreamCallResponsePayload) Reset()         { *m = StreamCallResponsePayload{} }
func (m *StreamCallResponsePayload) String() string { return proto.CompactTextString(m) }
func (*StreamCallResponsePayload) ProtoMessage()    {}
func (*StreamCallResponsePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_bad64dd5ad3c7973, []int{4}
}
func (m *StreamCallResponsePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallResponsePayload.Unmarshal(m, b)
}
func (m *StreamCallResponsePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallResponsePayload.Marshal(b, m, deterministic)
}
func (dst *StreamCallResponsePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallResponsePayload.Merge(dst, src)
}
func (m *StreamCallResponsePayload) XXX_Size() int {
	return xxx_messageInfo_StreamCallResponsePayload.Size(m)
}
func (m *StreamCallResponsePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallResponsePayload.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallResponsePayload proto.InternalMessageInfo

type isStreamCallResponsePayload_Payload interface {
	isStreamCallResponsePayload_Payload()
}

type StreamCallResponsePayload_Config struct {
	Config *StreamCallConfigResponse `protobuf:"bytes,1,opt,name=config,oneof"`
}
type StreamCallResponsePayload_Data struct {
	Data *StreamCallDataResponse `protobuf:"bytes,2,opt,name=data,oneof"`
}

func (*StreamCallResponsePayload_Config) isStreamCallResponsePayload_Payload() {}
func (*StreamCallResponsePayload_Data) isStreamCallResponsePayload_Payload()   {}

func (m *StreamCallResponsePayload) GetPayload() isStreamCallResponsePayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *StreamCallResponsePayload) GetConfig() *StreamCallConfigResponse {
	if x, ok := m.GetPayload().(*StreamCallResponsePayload_Config); ok {
		return x.Config
	}
	return nil
}

func (m *StreamCallResponsePayload) GetData() *StreamCallDataResponse {
	if x, ok := m.GetPayload().(*StreamCallResponsePayload_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamCallResponsePayload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamCallResponsePayload_OneofMarshaler, _StreamCallResponsePayload_OneofUnmarshaler, _StreamCallResponsePayload_OneofSizer, []interface{}{
		(*StreamCallResponsePayload_Config)(nil),
		(*StreamCallResponsePayload_Data)(nil),
	}
}

func _StreamCallResponsePayload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamCallResponsePayload)
	// payload
	switch x := m.Payload.(type) {
	case *StreamCallResponsePayload_Config:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *StreamCallResponsePayload_Data:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamCallResponsePayload.Payload has unexpected type %T", x)
	}
	return nil
}

func _StreamCallResponsePayload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamCallResponsePayload)
	switch tag {
	case 1: // payload.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallConfigResponse)
		err := b.DecodeMessage(msg)
		m.Payload = &StreamCallResponsePayload_Config{msg}
		return true, err
	case 2: // payload.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamCallDataResponse)
		err := b.DecodeMessage(msg)
		m.Payload = &StreamCallResponsePayload_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamCallResponsePayload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamCallResponsePayload)
	// payload
	switch x := m.Payload.(type) {
	case *StreamCallResponsePayload_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamCallResponsePayload_Data:
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

type StreamCallConfigRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ServiceName          *wrappers.StringValue `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	MethodName           *wrappers.StringValue `protobuf:"bytes,3,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StreamCallConfigRequest) Reset()         { *m = StreamCallConfigRequest{} }
func (m *StreamCallConfigRequest) String() string { return proto.CompactTextString(m) }
func (*StreamCallConfigRequest) ProtoMessage()    {}
func (*StreamCallConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_bad64dd5ad3c7973, []int{5}
}
func (m *StreamCallConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallConfigRequest.Unmarshal(m, b)
}
func (m *StreamCallConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallConfigRequest.Marshal(b, m, deterministic)
}
func (dst *StreamCallConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallConfigRequest.Merge(dst, src)
}
func (m *StreamCallConfigRequest) XXX_Size() int {
	return xxx_messageInfo_StreamCallConfigRequest.Size(m)
}
func (m *StreamCallConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallConfigRequest proto.InternalMessageInfo

func (m *StreamCallConfigRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *StreamCallConfigRequest) GetServiceName() *wrappers.StringValue {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *StreamCallConfigRequest) GetMethodName() *wrappers.StringValue {
	if m != nil {
		return m.MethodName
	}
	return nil
}

type StreamCallConfigResponse struct {
	Session              uint64   `protobuf:"varint,1,opt,name=session" json:"session,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ServiceName          string   `protobuf:"bytes,3,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	MethodName           string   `protobuf:"bytes,4,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamCallConfigResponse) Reset()         { *m = StreamCallConfigResponse{} }
func (m *StreamCallConfigResponse) String() string { return proto.CompactTextString(m) }
func (*StreamCallConfigResponse) ProtoMessage()    {}
func (*StreamCallConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_bad64dd5ad3c7973, []int{6}
}
func (m *StreamCallConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallConfigResponse.Unmarshal(m, b)
}
func (m *StreamCallConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallConfigResponse.Marshal(b, m, deterministic)
}
func (dst *StreamCallConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallConfigResponse.Merge(dst, src)
}
func (m *StreamCallConfigResponse) XXX_Size() int {
	return xxx_messageInfo_StreamCallConfigResponse.Size(m)
}
func (m *StreamCallConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallConfigResponse proto.InternalMessageInfo

func (m *StreamCallConfigResponse) GetSession() uint64 {
	if m != nil {
		return m.Session
	}
	return 0
}

func (m *StreamCallConfigResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamCallConfigResponse) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *StreamCallConfigResponse) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

type StreamCallDataRequest struct {
	Value                *any.Any `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamCallDataRequest) Reset()         { *m = StreamCallDataRequest{} }
func (m *StreamCallDataRequest) String() string { return proto.CompactTextString(m) }
func (*StreamCallDataRequest) ProtoMessage()    {}
func (*StreamCallDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_bad64dd5ad3c7973, []int{7}
}
func (m *StreamCallDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallDataRequest.Unmarshal(m, b)
}
func (m *StreamCallDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallDataRequest.Marshal(b, m, deterministic)
}
func (dst *StreamCallDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallDataRequest.Merge(dst, src)
}
func (m *StreamCallDataRequest) XXX_Size() int {
	return xxx_messageInfo_StreamCallDataRequest.Size(m)
}
func (m *StreamCallDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallDataRequest proto.InternalMessageInfo

func (m *StreamCallDataRequest) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type StreamCallDataResponse struct {
	Value                *any.Any `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamCallDataResponse) Reset()         { *m = StreamCallDataResponse{} }
func (m *StreamCallDataResponse) String() string { return proto.CompactTextString(m) }
func (*StreamCallDataResponse) ProtoMessage()    {}
func (*StreamCallDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payload_bad64dd5ad3c7973, []int{8}
}
func (m *StreamCallDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamCallDataResponse.Unmarshal(m, b)
}
func (m *StreamCallDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamCallDataResponse.Marshal(b, m, deterministic)
}
func (dst *StreamCallDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamCallDataResponse.Merge(dst, src)
}
func (m *StreamCallDataResponse) XXX_Size() int {
	return xxx_messageInfo_StreamCallDataResponse.Size(m)
}
func (m *StreamCallDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamCallDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamCallDataResponse proto.InternalMessageInfo

func (m *StreamCallDataResponse) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamErrorResponsePayload)(nil), "ai.metathings.service.core.StreamErrorResponsePayload")
	proto.RegisterType((*UnaryCallRequestPayload)(nil), "ai.metathings.service.core.UnaryCallRequestPayload")
	proto.RegisterType((*UnaryCallResponsePayload)(nil), "ai.metathings.service.core.UnaryCallResponsePayload")
	proto.RegisterType((*StreamCallRequestPayload)(nil), "ai.metathings.service.core.StreamCallRequestPayload")
	proto.RegisterType((*StreamCallResponsePayload)(nil), "ai.metathings.service.core.StreamCallResponsePayload")
	proto.RegisterType((*StreamCallConfigRequest)(nil), "ai.metathings.service.core.StreamCallConfigRequest")
	proto.RegisterType((*StreamCallConfigResponse)(nil), "ai.metathings.service.core.StreamCallConfigResponse")
	proto.RegisterType((*StreamCallDataRequest)(nil), "ai.metathings.service.core.StreamCallDataRequest")
	proto.RegisterType((*StreamCallDataResponse)(nil), "ai.metathings.service.core.StreamCallDataResponse")
}

func init() { proto.RegisterFile("payload.proto", fileDescriptor_payload_bad64dd5ad3c7973) }

var fileDescriptor_payload_bad64dd5ad3c7973 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x71, 0x1b, 0x5a, 0x75, 0x16, 0x2e, 0x11, 0xd0, 0x34, 0x42, 0xfc, 0xc9, 0x09, 0x71,
	0x70, 0x61, 0xcb, 0x15, 0x21, 0xd8, 0x22, 0x7a, 0xa1, 0x42, 0xa9, 0xe0, 0x8a, 0xdc, 0xec, 0x34,
	0x8d, 0x94, 0xd8, 0xc1, 0xf6, 0x16, 0xf2, 0x0e, 0x88, 0x47, 0xe0, 0x4d, 0x90, 0xb8, 0xf3, 0x40,
	0x5c, 0xd1, 0xda, 0x4e, 0x37, 0xd9, 0xcd, 0x12, 0x72, 0x41, 0xea, 0x6d, 0xe3, 0x99, 0xf9, 0x3c,
	0x3f, 0xcf, 0x7c, 0x0b, 0x37, 0x4b, 0x56, 0xe5, 0x82, 0x4d, 0x69, 0x29, 0x85, 0x16, 0x7e, 0xc8,
	0x32, 0x5a, 0xa0, 0x66, 0xfa, 0x3c, 0xe3, 0xa9, 0xa2, 0x0a, 0xe5, 0x45, 0x96, 0x20, 0x4d, 0x84,
	0xc4, 0xf0, 0x5e, 0x2a, 0x44, 0x9a, 0xe3, 0xbe, 0xc9, 0x3c, 0x9d, 0x9d, 0xed, 0x7f, 0x96, 0xac,
	0x2c, 0x51, 0x2a, 0x5b, 0x1b, 0xee, 0x2d, 0xc7, 0x19, 0xaf, 0x6c, 0x28, 0xfa, 0x46, 0x20, 0x3c,
	0xd1, 0x12, 0x59, 0xf1, 0x5a, 0x4a, 0x21, 0x63, 0x54, 0xa5, 0xe0, 0x0a, 0xdf, 0xd9, 0xbb, 0x7d,
	0x1f, 0x3c, 0xce, 0x0a, 0x0c, 0xc8, 0x03, 0xf2, 0x68, 0x27, 0x36, 0xbf, 0xfd, 0x87, 0x70, 0xc3,
	0xdd, 0xfe, 0xd1, 0xc4, 0x36, 0x4c, 0x6c, 0xe4, 0xce, 0x8e, 0xe7, 0x29, 0xf7, 0x61, 0x54, 0xa0,
	0x3e, 0x17, 0x53, 0x9b, 0xb1, 0x69, 0x32, 0xc0, 0x1e, 0x99, 0x84, 0x00, 0xb6, 0x13, 0xc1, 0x35,
	0x7e, 0xd1, 0x81, 0x67, 0x82, 0xf5, 0x67, 0xf4, 0x9b, 0xc0, 0xee, 0x7b, 0xce, 0x64, 0x35, 0x61,
	0x79, 0x1e, 0xe3, 0xa7, 0x19, 0x2a, 0x5d, 0x77, 0xf3, 0xa4, 0xd1, 0xcd, 0x68, 0x7c, 0x97, 0x5a,
	0x2c, 0x5a, 0x63, 0xd1, 0x13, 0x2d, 0x33, 0x9e, 0x7e, 0x60, 0xf9, 0x0c, 0x5d, 0xaf, 0x2f, 0x3a,
	0x7a, 0xed, 0xab, 0x6c, 0x91, 0x3c, 0x5f, 0x25, 0xe9, 0xab, 0x6f, 0x72, 0x3e, 0x86, 0xeb, 0x17,
	0xf3, 0x43, 0x43, 0x39, 0x1a, 0xdf, 0x5a, 0x29, 0x7c, 0xc9, 0xab, 0xd8, 0xa6, 0x44, 0xdf, 0x09,
	0x04, 0x0d, 0xf2, 0xff, 0x33, 0x88, 0x21, 0x0d, 0xfe, 0x20, 0x10, 0xd8, 0x5d, 0xe9, 0x98, 0xcd,
	0x5b, 0xd8, 0x4a, 0x04, 0x3f, 0xcb, 0x52, 0x37, 0x9d, 0x03, 0xba, 0x7e, 0x61, 0xe9, 0x42, 0x65,
	0x62, 0x6a, 0x9c, 0xd6, 0xd1, 0xb5, 0xd8, 0x89, 0xf8, 0x6f, 0xc0, 0x9b, 0x32, 0xcd, 0xdc, 0xc0,
	0x9e, 0xfe, 0x9b, 0xd8, 0x21, 0xd3, 0x6c, 0x21, 0x65, 0x04, 0x5e, 0xed, 0xc0, 0xb6, 0x33, 0x52,
	0xf4, 0x93, 0xc0, 0x5e, 0xb3, 0xff, 0xf6, 0x0b, 0x1f, 0x2f, 0x01, 0x3c, 0x1b, 0x06, 0x60, 0xc5,
	0x1a, 0x04, 0x47, 0x2d, 0x82, 0xf1, 0x10, 0x82, 0x4b, 0xad, 0x15, 0x84, 0x5f, 0x04, 0x76, 0xd7,
	0x3c, 0xde, 0xd5, 0x73, 0x47, 0xf4, 0xb5, 0xb5, 0x50, 0xed, 0x97, 0x9c, 0xff, 0x45, 0x28, 0x54,
	0x2a, 0x13, 0xdc, 0x10, 0x79, 0x71, 0xfd, 0x79, 0xe9, 0x85, 0x8d, 0xbf, 0x78, 0x61, 0xb3, 0xd7,
	0x0b, 0xde, 0xb2, 0x17, 0xa2, 0x09, 0xdc, 0xee, 0xdc, 0xa5, 0x85, 0x49, 0x48, 0xbf, 0x49, 0x0e,
	0xe1, 0x4e, 0xf7, 0x38, 0x87, 0xa8, 0x9c, 0x6e, 0x99, 0xc3, 0x83, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xff, 0xa0, 0x8a, 0xcf, 0x05, 0x06, 0x00, 0x00,
}
