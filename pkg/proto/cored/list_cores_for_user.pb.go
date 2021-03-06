// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_cores_for_user.proto

package cored

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import state "github.com/nayotta/metathings/pkg/proto/common/state"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListCoresForUserRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ProjectId            *wrappers.StringValue `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	State                state.CoreState       `protobuf:"varint,4,opt,name=state,enum=ai.metathings.state.CoreState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListCoresForUserRequest) Reset()         { *m = ListCoresForUserRequest{} }
func (m *ListCoresForUserRequest) String() string { return proto.CompactTextString(m) }
func (*ListCoresForUserRequest) ProtoMessage()    {}
func (*ListCoresForUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_cores_for_user_5283205ebc97f0e2, []int{0}
}
func (m *ListCoresForUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCoresForUserRequest.Unmarshal(m, b)
}
func (m *ListCoresForUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCoresForUserRequest.Marshal(b, m, deterministic)
}
func (dst *ListCoresForUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCoresForUserRequest.Merge(dst, src)
}
func (m *ListCoresForUserRequest) XXX_Size() int {
	return xxx_messageInfo_ListCoresForUserRequest.Size(m)
}
func (m *ListCoresForUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCoresForUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCoresForUserRequest proto.InternalMessageInfo

func (m *ListCoresForUserRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListCoresForUserRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *ListCoresForUserRequest) GetState() state.CoreState {
	if m != nil {
		return m.State
	}
	return state.CoreState_CORE_STATE_UNKNOWN
}

type ListCoresForUserResponse struct {
	Cores                []*Core  `protobuf:"bytes,1,rep,name=cores" json:"cores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCoresForUserResponse) Reset()         { *m = ListCoresForUserResponse{} }
func (m *ListCoresForUserResponse) String() string { return proto.CompactTextString(m) }
func (*ListCoresForUserResponse) ProtoMessage()    {}
func (*ListCoresForUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_cores_for_user_5283205ebc97f0e2, []int{1}
}
func (m *ListCoresForUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCoresForUserResponse.Unmarshal(m, b)
}
func (m *ListCoresForUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCoresForUserResponse.Marshal(b, m, deterministic)
}
func (dst *ListCoresForUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCoresForUserResponse.Merge(dst, src)
}
func (m *ListCoresForUserResponse) XXX_Size() int {
	return xxx_messageInfo_ListCoresForUserResponse.Size(m)
}
func (m *ListCoresForUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCoresForUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCoresForUserResponse proto.InternalMessageInfo

func (m *ListCoresForUserResponse) GetCores() []*Core {
	if m != nil {
		return m.Cores
	}
	return nil
}

func init() {
	proto.RegisterType((*ListCoresForUserRequest)(nil), "ai.metathings.service.cored.ListCoresForUserRequest")
	proto.RegisterType((*ListCoresForUserResponse)(nil), "ai.metathings.service.cored.ListCoresForUserResponse")
}

func init() {
	proto.RegisterFile("list_cores_for_user.proto", fileDescriptor_list_cores_for_user_5283205ebc97f0e2)
}

var fileDescriptor_list_cores_for_user_5283205ebc97f0e2 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0xb6, 0x82, 0x5b, 0xf0, 0x90, 0x8b, 0xb1, 0x4a, 0xa9, 0x3d, 0xf5, 0xd2, 0x5d,
	0xa9, 0xa2, 0x07, 0x8f, 0xa2, 0x20, 0x78, 0x4a, 0xd1, 0x6b, 0xd8, 0x26, 0xd3, 0xed, 0xda, 0x24,
	0x13, 0x77, 0x26, 0x2d, 0xfe, 0x39, 0x7f, 0x9b, 0x64, 0x53, 0x3f, 0x50, 0x10, 0x6f, 0x2f, 0xe4,
	0x99, 0x37, 0xcf, 0xec, 0x88, 0xa3, 0xdc, 0x12, 0x27, 0x29, 0x3a, 0xa0, 0x64, 0x81, 0x2e, 0xa9,
	0x09, 0x9c, 0xac, 0x1c, 0x32, 0x86, 0xc7, 0xda, 0xca, 0x02, 0x58, 0xf3, 0xd2, 0x96, 0x86, 0x24,
	0x81, 0x5b, 0xdb, 0x14, 0x64, 0xc3, 0x66, 0xfd, 0x81, 0x41, 0x34, 0x39, 0x28, 0x8f, 0xce, 0xeb,
	0x85, 0xda, 0x38, 0x5d, 0x55, 0xe0, 0xa8, 0x1d, 0xee, 0x5f, 0x1a, 0xcb, 0xcb, 0x7a, 0x2e, 0x53,
	0x2c, 0x54, 0xb1, 0xb1, 0xbc, 0xc2, 0x8d, 0x32, 0x38, 0xf1, 0x1f, 0x27, 0x6b, 0x9d, 0xdb, 0x4c,
	0x33, 0x3a, 0x52, 0x9f, 0x71, 0x3b, 0x27, 0x9a, 0xfa, 0x6d, 0xbe, 0xfd, 0xd6, 0x51, 0xea, 0x57,
	0x64, 0xd6, 0xea, 0x4b, 0x48, 0x55, 0x2b, 0xd3, 0xfe, 0x5e, 0xa5, 0x58, 0x14, 0x58, 0x2a, 0x62,
	0xcd, 0xa0, 0x9a, 0x82, 0xc4, 0xc7, 0xb6, 0x66, 0xf4, 0x16, 0x88, 0xc3, 0x07, 0x4b, 0x7c, 0xd3,
	0x2c, 0x79, 0x87, 0xee, 0x91, 0xc0, 0xc5, 0xf0, 0x52, 0x03, 0x71, 0x78, 0x26, 0x3a, 0xa5, 0x2e,
	0x20, 0x0a, 0x86, 0xc1, 0xb8, 0x37, 0x3d, 0x91, 0xed, 0x56, 0xf2, 0x63, 0x2b, 0x39, 0x63, 0x67,
	0x4b, 0xf3, 0xa4, 0xf3, 0x1a, 0x62, 0x4f, 0x86, 0xd7, 0x42, 0x54, 0x0e, 0x9f, 0x21, 0xe5, 0xc4,
	0x66, 0xd1, 0xce, 0x3f, 0xe6, 0xf6, 0xb7, 0xfc, 0x7d, 0x16, 0x5e, 0x88, 0xae, 0x37, 0x8b, 0x3a,
	0xc3, 0x60, 0x7c, 0x30, 0x1d, 0xc8, 0x1f, 0x4f, 0xec, 0xad, 0x1b, 0xcf, 0x59, 0x93, 0xe2, 0x16,
	0x1e, 0xcd, 0x44, 0xf4, 0xdb, 0x9f, 0x2a, 0x2c, 0x09, 0xc2, 0x2b, 0xd1, 0xf5, 0xc7, 0x8b, 0x82,
	0xe1, 0xee, 0xb8, 0x37, 0x3d, 0x95, 0x7f, 0x1c, 0xcd, 0x37, 0xc7, 0x2d, 0x3f, 0xdf, 0xf3, 0xae,
	0xe7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x65, 0xe7, 0x9c, 0x01, 0x02, 0x00, 0x00,
}
