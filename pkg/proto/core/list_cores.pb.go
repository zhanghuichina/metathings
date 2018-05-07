// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_cores.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import state "github.com/bigdatagz/metathings/pkg/proto/common/state"
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

type ListCoresRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ProjectId            *wrappers.StringValue `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	OwnerId              *wrappers.StringValue `protobuf:"bytes,3,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	State                state.CoreState       `protobuf:"varint,4,opt,name=state,enum=ai.metathings.state.CoreState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListCoresRequest) Reset()         { *m = ListCoresRequest{} }
func (m *ListCoresRequest) String() string { return proto.CompactTextString(m) }
func (*ListCoresRequest) ProtoMessage()    {}
func (*ListCoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_cores_c9e34e088f468c00, []int{0}
}
func (m *ListCoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCoresRequest.Unmarshal(m, b)
}
func (m *ListCoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCoresRequest.Marshal(b, m, deterministic)
}
func (dst *ListCoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCoresRequest.Merge(dst, src)
}
func (m *ListCoresRequest) XXX_Size() int {
	return xxx_messageInfo_ListCoresRequest.Size(m)
}
func (m *ListCoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCoresRequest proto.InternalMessageInfo

func (m *ListCoresRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListCoresRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *ListCoresRequest) GetOwnerId() *wrappers.StringValue {
	if m != nil {
		return m.OwnerId
	}
	return nil
}

func (m *ListCoresRequest) GetState() state.CoreState {
	if m != nil {
		return m.State
	}
	return state.CoreState_CORE_STATE_UNKNOWN
}

type ListCoresResponse struct {
	Cores                []*Core  `protobuf:"bytes,1,rep,name=cores" json:"cores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCoresResponse) Reset()         { *m = ListCoresResponse{} }
func (m *ListCoresResponse) String() string { return proto.CompactTextString(m) }
func (*ListCoresResponse) ProtoMessage()    {}
func (*ListCoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_cores_c9e34e088f468c00, []int{1}
}
func (m *ListCoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCoresResponse.Unmarshal(m, b)
}
func (m *ListCoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCoresResponse.Marshal(b, m, deterministic)
}
func (dst *ListCoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCoresResponse.Merge(dst, src)
}
func (m *ListCoresResponse) XXX_Size() int {
	return xxx_messageInfo_ListCoresResponse.Size(m)
}
func (m *ListCoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCoresResponse proto.InternalMessageInfo

func (m *ListCoresResponse) GetCores() []*Core {
	if m != nil {
		return m.Cores
	}
	return nil
}

func init() {
	proto.RegisterType((*ListCoresRequest)(nil), "ai.metathings.service.core.ListCoresRequest")
	proto.RegisterType((*ListCoresResponse)(nil), "ai.metathings.service.core.ListCoresResponse")
}

func init() { proto.RegisterFile("list_cores.proto", fileDescriptor_list_cores_c9e34e088f468c00) }

var fileDescriptor_list_cores_c9e34e088f468c00 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xc9, 0xd7, 0xf6, 0x53, 0xb7, 0x20, 0x35, 0xa7, 0x50, 0xa4, 0x84, 0x9e, 0x7a, 0xe9,
	0xae, 0x54, 0xa9, 0x07, 0x8f, 0x1e, 0xa4, 0xe8, 0x29, 0x05, 0xaf, 0x65, 0x93, 0x8c, 0xdb, 0xb5,
	0x49, 0x26, 0xee, 0x4e, 0x1a, 0xf0, 0x57, 0xfb, 0x13, 0x24, 0x9b, 0xa8, 0x45, 0x10, 0x7a, 0x7b,
	0x49, 0x9e, 0x77, 0x76, 0x1e, 0x86, 0x8d, 0x32, 0x6d, 0x69, 0x93, 0xa0, 0x01, 0xcb, 0x4b, 0x83,
	0x84, 0xfe, 0x58, 0x6a, 0x9e, 0x03, 0x49, 0xda, 0xea, 0x42, 0x59, 0x6e, 0xc1, 0xec, 0x75, 0x02,
	0xbc, 0x41, 0xc6, 0x13, 0x85, 0xa8, 0x32, 0x10, 0x8e, 0x8c, 0xab, 0x17, 0x51, 0x1b, 0x59, 0x96,
	0x60, 0xba, 0xee, 0x78, 0xa9, 0x34, 0x6d, 0xab, 0x98, 0x27, 0x98, 0x8b, 0xbc, 0xd6, 0xb4, 0xc3,
	0x5a, 0x28, 0x9c, 0xbb, 0x9f, 0xf3, 0xbd, 0xcc, 0x74, 0x2a, 0x09, 0x8d, 0x15, 0xdf, 0xb1, 0xeb,
	0x3d, 0x1c, 0xf4, 0x62, 0xad, 0x52, 0x49, 0x52, 0xbd, 0x8b, 0x9f, 0x2d, 0x44, 0xb9, 0x53, 0xed,
	0xa3, 0x22, 0xc1, 0x3c, 0xc7, 0x42, 0x58, 0x92, 0x04, 0xa2, 0x59, 0x6a, 0xe3, 0x62, 0x37, 0x88,
	0x35, 0x5f, 0xda, 0x3c, 0xfd, 0xf0, 0xd8, 0xe8, 0x49, 0x5b, 0xba, 0x6f, 0xe4, 0x22, 0x78, 0xab,
	0xc0, 0x92, 0x7f, 0xc5, 0xfa, 0x85, 0xcc, 0x21, 0xf0, 0x42, 0x6f, 0x36, 0x5c, 0x5c, 0xf2, 0x56,
	0x88, 0x7f, 0x09, 0xf1, 0x35, 0x19, 0x5d, 0xa8, 0x67, 0x99, 0x55, 0x10, 0x39, 0xd2, 0xbf, 0x63,
	0xac, 0x34, 0xf8, 0x0a, 0x09, 0x6d, 0x74, 0x1a, 0xfc, 0x3b, 0xa2, 0x77, 0xd6, 0xf1, 0xab, 0xd4,
	0xbf, 0x65, 0xa7, 0x58, 0x17, 0x60, 0x9a, 0x6a, 0xef, 0x88, 0xea, 0x89, 0xa3, 0x57, 0xa9, 0x7f,
	0xc3, 0x06, 0xce, 0x2b, 0xe8, 0x87, 0xde, 0xec, 0x7c, 0x31, 0xe1, 0xbf, 0xae, 0xe2, 0x9c, 0x1b,
	0xb3, 0x75, 0x93, 0xa2, 0x16, 0x9e, 0x3e, 0xb2, 0x8b, 0x03, 0x63, 0x5b, 0x62, 0x61, 0xc1, 0x5f,
	0xb2, 0x81, 0xbb, 0x6f, 0xe0, 0x85, 0xbd, 0xd9, 0x70, 0x11, 0xf2, 0xbf, 0x0f, 0xec, 0x26, 0x46,
	0x2d, 0x1e, 0xff, 0x77, 0x1b, 0x5e, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x35, 0xd0, 0x50, 0xe4,
	0x23, 0x02, 0x00, 0x00,
}
