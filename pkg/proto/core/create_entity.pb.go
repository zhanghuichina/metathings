// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_entity.proto

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

type CreateEntityRequest struct {
	CoreId      *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=core_id,json=coreId" json:"core_id,omitempty"`
	Name        *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ServiceName *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	Endpoint    *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *CreateEntityRequest) Reset()                    { *m = CreateEntityRequest{} }
func (m *CreateEntityRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateEntityRequest) ProtoMessage()               {}
func (*CreateEntityRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CreateEntityRequest) GetCoreId() *google_protobuf.StringValue {
	if m != nil {
		return m.CoreId
	}
	return nil
}

func (m *CreateEntityRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateEntityRequest) GetServiceName() *google_protobuf.StringValue {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *CreateEntityRequest) GetEndpoint() *google_protobuf.StringValue {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type CreateEntityResponse struct {
	Entity *Entity `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
}

func (m *CreateEntityResponse) Reset()                    { *m = CreateEntityResponse{} }
func (m *CreateEntityResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateEntityResponse) ProtoMessage()               {}
func (*CreateEntityResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CreateEntityResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateEntityRequest)(nil), "ai.metathings.service.core.CreateEntityRequest")
	proto.RegisterType((*CreateEntityResponse)(nil), "ai.metathings.service.core.CreateEntityResponse")
}

func init() { proto.RegisterFile("create_entity.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0x69, 0x3d, 0xaa, 0xe4, 0x6e, 0xea, 0x39, 0x94, 0x22, 0x7a, 0x74, 0x72, 0xb9, 0x54,
	0x14, 0x1d, 0x9c, 0x44, 0x11, 0x71, 0x71, 0xa8, 0xe0, 0x5a, 0xd2, 0xf6, 0x33, 0x17, 0x6c, 0x93,
	0x9a, 0x7c, 0xbd, 0xe2, 0x4f, 0xf0, 0x57, 0x0a, 0xfe, 0x12, 0x69, 0x1a, 0x0f, 0x1d, 0x44, 0xdd,
	0x42, 0xf2, 0x3e, 0x5f, 0x9e, 0xbc, 0x21, 0xf3, 0x52, 0x03, 0x43, 0xc8, 0x41, 0xa2, 0xc0, 0x17,
	0xda, 0x6a, 0x85, 0x2a, 0x8c, 0x99, 0xa0, 0x0d, 0x20, 0xc3, 0x95, 0x90, 0xdc, 0x50, 0x03, 0x7a,
	0x2d, 0x4a, 0xa0, 0xa5, 0xd2, 0x10, 0xef, 0x73, 0xa5, 0x78, 0x0d, 0xa9, 0x4d, 0x16, 0xdd, 0x63,
	0xda, 0x6b, 0xd6, 0xb6, 0xa0, 0xcd, 0xc8, 0xc6, 0x67, 0x5c, 0xe0, 0xaa, 0x2b, 0x68, 0xa9, 0x9a,
	0xb4, 0xe9, 0x05, 0x3e, 0xa9, 0x3e, 0xe5, 0x6a, 0x69, 0x0f, 0x97, 0x6b, 0x56, 0x8b, 0x8a, 0xa1,
	0xd2, 0x26, 0xdd, 0x2c, 0x1d, 0x37, 0xfb, 0x6a, 0x90, 0xbc, 0xfa, 0x64, 0x7e, 0x65, 0xcd, 0xae,
	0xed, 0x76, 0x06, 0xcf, 0x1d, 0x18, 0x0c, 0x4f, 0xc9, 0xf6, 0x60, 0x91, 0x8b, 0x2a, 0xf2, 0x16,
	0xde, 0xe1, 0xf4, 0x78, 0x8f, 0x8e, 0x3e, 0xf4, 0xd3, 0x87, 0xde, 0xa3, 0x16, 0x92, 0x3f, 0xb0,
	0xba, 0x83, 0x2c, 0x18, 0xc2, 0xb7, 0x55, 0x78, 0x44, 0x26, 0x92, 0x35, 0x10, 0xf9, 0x7f, 0x60,
	0x6c, 0x32, 0xbc, 0x21, 0x33, 0xf7, 0xec, 0xdc, 0x92, 0x5b, 0xbf, 0x93, 0x97, 0xc1, 0xfb, 0xdb,
	0x81, 0xbf, 0xf0, 0xb2, 0xa9, 0x23, 0xef, 0x86, 0x41, 0x17, 0x64, 0x07, 0x64, 0xd5, 0x2a, 0x21,
	0x31, 0x9a, 0xfc, 0x63, 0xc8, 0x86, 0x4a, 0x32, 0xb2, 0xfb, 0xbd, 0x0a, 0xd3, 0x2a, 0x69, 0x20,
	0x3c, 0x27, 0xc1, 0xd8, 0x99, 0xab, 0x22, 0xa1, 0x3f, 0x7f, 0x1b, 0x75, 0xac, 0x23, 0x8a, 0xc0,
	0xde, 0x7d, 0xf2, 0x11, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xef, 0xd0, 0xec, 0xff, 0x01, 0x00, 0x00,
}
