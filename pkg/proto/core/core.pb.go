// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	core.proto
	create_core.proto
	delete_core.proto
	get_core.proto
	heartbeat.proto
	list_cores.proto
	list_cores_for_user.proto
	patch_core.proto
	pipeline.proto
	send_unary_call.proto
	service.proto
	unary_call.proto

It has these top-level messages:
	Core
	CreateCoreRequest
	CreateCoreResponse
	DeleteCoreRequest
	GetCoreRequest
	GetCoreResponse
	HeartbeatRequest
	ListCoresRequest
	ListCoresResponse
	ListCoresForUserRequest
	ListCoresForUserResponse
	PatchCoreRequest
	PatchCoreResponse
	PipelineOutStream
	PipelineInStream
	SendUnaryCallRequest
	SendUnaryCallResponse
	UnaryCallPayload
*/
package core

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

type CoreState int32

const (
	CoreState_CORE_STATE_UNKNOWN CoreState = 0
	CoreState_CORE_STATE_ONLINE  CoreState = 1
	CoreState_CORE_STATE_OFFLINE CoreState = 2
)

var CoreState_name = map[int32]string{
	0: "CORE_STATE_UNKNOWN",
	1: "CORE_STATE_ONLINE",
	2: "CORE_STATE_OFFLINE",
}
var CoreState_value = map[string]int32{
	"CORE_STATE_UNKNOWN": 0,
	"CORE_STATE_ONLINE":  1,
	"CORE_STATE_OFFLINE": 2,
}

func (x CoreState) String() string {
	return proto.EnumName(CoreState_name, int32(x))
}
func (CoreState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Core struct {
	Id        string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name      string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ProjectId string    `protobuf:"bytes,3,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	OwnerId   string    `protobuf:"bytes,4,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	State     CoreState `protobuf:"varint,5,opt,name=state,enum=ai.metathings.service.core.CoreState" json:"state,omitempty"`
}

func (m *Core) Reset()                    { *m = Core{} }
func (m *Core) String() string            { return proto.CompactTextString(m) }
func (*Core) ProtoMessage()               {}
func (*Core) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Core) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Core) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Core) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Core) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Core) GetState() CoreState {
	if m != nil {
		return m.State
	}
	return CoreState_CORE_STATE_UNKNOWN
}

func init() {
	proto.RegisterType((*Core)(nil), "ai.metathings.service.core.Core")
	proto.RegisterEnum("ai.metathings.service.core.CoreState", CoreState_name, CoreState_value)
}

func init() { proto.RegisterFile("core.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x2f, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4a, 0xcc, 0xd4, 0xcb, 0x4d, 0x2d, 0x49, 0x2c,
	0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0xa9,
	0x50, 0x5a, 0xc8, 0xc8, 0xc5, 0xe2, 0x9c, 0x5f, 0x94, 0x2a, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98,
	0x9b, 0x2a, 0xc1, 0x04, 0x16, 0x01, 0xb3, 0x85, 0x64, 0xb9, 0xb8, 0x0a, 0x8a, 0xf2, 0xb3, 0x52,
	0x93, 0x4b, 0xe2, 0x33, 0x53, 0x24, 0x98, 0xc1, 0x32, 0x9c, 0x50, 0x11, 0xcf, 0x14, 0x21, 0x49,
	0x2e, 0x8e, 0xfc, 0xf2, 0xbc, 0xd4, 0x22, 0x90, 0x24, 0x0b, 0x58, 0x92, 0x1d, 0xcc, 0xf7, 0x4c,
	0x11, 0xb2, 0xe6, 0x62, 0x2d, 0x2e, 0x49, 0x2c, 0x49, 0x95, 0x60, 0x55, 0x60, 0xd4, 0xe0, 0x33,
	0x52, 0xd5, 0xc3, 0xed, 0x24, 0x3d, 0x90, 0x73, 0x82, 0x41, 0x8a, 0x83, 0x20, 0x7a, 0xb4, 0x82,
	0xb8, 0x38, 0xe1, 0x62, 0x42, 0x62, 0x5c, 0x42, 0xce, 0xfe, 0x41, 0xae, 0xf1, 0xc1, 0x21, 0x8e,
	0x21, 0xae, 0xf1, 0xa1, 0x7e, 0xde, 0x7e, 0xfe, 0xe1, 0x7e, 0x02, 0x0c, 0x42, 0xa2, 0x5c, 0x82,
	0x48, 0xe2, 0xfe, 0x7e, 0x3e, 0x9e, 0x7e, 0xae, 0x02, 0x8c, 0x68, 0xca, 0xfd, 0xdd, 0xdc, 0xc0,
	0xe2, 0x4c, 0x49, 0x6c, 0xe0, 0xa0, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x42, 0x41, 0x85,
	0x78, 0x28, 0x01, 0x00, 0x00,
}