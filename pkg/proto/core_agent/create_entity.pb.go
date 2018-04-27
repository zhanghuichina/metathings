// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_entity.proto

/*
Package core_agent is a generated protocol buffer package.

It is generated from these files:
	create_entity.proto
	create_or_get_entity.proto
	delete_entity.proto
	entity.proto
	get_entity.proto
	heartbeat.proto
	list_entities.proto
	patch_entity.proto
	service.proto

It has these top-level messages:
	CreateEntityRequest
	CreateEntityResponse
	CreateOrGetEntityRequest
	CreateOrGetEntityResponse
	DeleteEntityRequest
	Entity
	GetEntityRequest
	GetEntityResponse
	HeartbeatRequest
	ListEntitiesRequest
	ListEntitiesResponse
	PatchEntityRequest
	PatchEntityResponse
*/
package core_agent

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
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

type CreateEntityRequest struct {
	Name        *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ServiceName *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	Endpoint    *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *CreateEntityRequest) Reset()                    { *m = CreateEntityRequest{} }
func (m *CreateEntityRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateEntityRequest) ProtoMessage()               {}
func (*CreateEntityRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (*CreateEntityResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateEntityResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateEntityRequest)(nil), "ai.metathings.service.core_agent.CreateEntityRequest")
	proto.RegisterType((*CreateEntityResponse)(nil), "ai.metathings.service.core_agent.CreateEntityResponse")
}

func init() { proto.RegisterFile("create_entity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xe9, 0x94, 0x21, 0xd9, 0x4e, 0x9d, 0x87, 0x31, 0x44, 0xcb, 0x4e, 0xbb, 0xec, 0xab,
	0x28, 0x78, 0x16, 0x45, 0xbc, 0x79, 0xa8, 0x20, 0xde, 0x4a, 0xda, 0x7d, 0x66, 0xc1, 0x36, 0x5f,
	0x4d, 0xbe, 0xae, 0xf8, 0x37, 0xfd, 0x03, 0x82, 0xbf, 0x44, 0x96, 0xc6, 0xa1, 0xa7, 0xe1, 0x2d,
	0x24, 0xef, 0xf3, 0xf2, 0xbc, 0x44, 0x4c, 0x4a, 0x8b, 0x92, 0x31, 0x47, 0xc3, 0x9a, 0xdf, 0xa1,
	0xb1, 0xc4, 0x14, 0x27, 0x52, 0x43, 0x8d, 0x2c, 0x79, 0xad, 0x8d, 0x72, 0xe0, 0xd0, 0x6e, 0x74,
	0x89, 0x50, 0x92, 0xc5, 0x5c, 0x2a, 0x34, 0x3c, 0x3b, 0x55, 0x44, 0xaa, 0xc2, 0xd4, 0xe7, 0x8b,
	0xf6, 0x25, 0xed, 0xac, 0x6c, 0x1a, 0xb4, 0xae, 0x6f, 0x98, 0x5d, 0x29, 0xcd, 0xeb, 0xb6, 0x80,
	0x92, 0xea, 0xb4, 0xee, 0x34, 0xbf, 0x52, 0x97, 0x2a, 0x5a, 0xfa, 0xc7, 0xe5, 0x46, 0x56, 0x7a,
	0x25, 0x99, 0xac, 0x4b, 0x77, 0xc7, 0xc0, 0x8d, 0x7f, 0x7b, 0xcc, 0x3f, 0x22, 0x31, 0xb9, 0xf5,
	0x7e, 0x77, 0xfe, 0x3a, 0xc3, 0xb7, 0x16, 0x1d, 0xc7, 0xe7, 0xe2, 0xd0, 0xc8, 0x1a, 0xa7, 0x51,
	0x12, 0x2d, 0x46, 0x17, 0x27, 0xd0, 0xcb, 0xc0, 0x8f, 0x0c, 0x3c, 0xb2, 0xd5, 0x46, 0x3d, 0xc9,
	0xaa, 0xc5, 0xcc, 0x27, 0xe3, 0x7b, 0x31, 0x0e, 0x2b, 0x72, 0x4f, 0x0e, 0xf6, 0x93, 0x37, 0xc3,
	0xaf, 0xcf, 0xb3, 0x41, 0x12, 0x65, 0xa3, 0x40, 0x3e, 0x6c, 0x8b, 0xae, 0xc5, 0x11, 0x9a, 0x55,
	0x43, 0xda, 0xf0, 0xf4, 0xe0, 0x1f, 0x25, 0x3b, 0x6a, 0xfe, 0x2c, 0x8e, 0xff, 0x6e, 0x72, 0x0d,
	0x19, 0xb7, 0x6d, 0x1e, 0xf6, 0xe3, 0xc3, 0xac, 0x05, 0xec, 0xfb, 0x05, 0x08, 0x0d, 0x81, 0x2b,
	0x86, 0xde, 0xe0, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x0d, 0x7f, 0x59, 0xd4, 0x01, 0x00,
	0x00,
}
