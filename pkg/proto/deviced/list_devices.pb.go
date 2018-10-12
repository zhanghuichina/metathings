// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_devices.proto

package deviced

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import state "github.com/nayotta/metathings/pkg/proto/constant/state"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListDevicesRequest struct {
	Name                 *wrappers.StringValue   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ProjectId            *wrappers.StringValue   `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	Owners               []*wrappers.StringValue `protobuf:"bytes,3,rep,name=owners" json:"owners,omitempty"`
	Groups               []*wrappers.StringValue `protobuf:"bytes,4,rep,name=groups" json:"groups,omitempty"`
	State                state.DeviceState       `protobuf:"varint,5,opt,name=state,enum=ai.metathings.constant.state.DeviceState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ListDevicesRequest) Reset()         { *m = ListDevicesRequest{} }
func (m *ListDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*ListDevicesRequest) ProtoMessage()    {}
func (*ListDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_devices_847511c5a6bf64d0, []int{0}
}
func (m *ListDevicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDevicesRequest.Unmarshal(m, b)
}
func (m *ListDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDevicesRequest.Marshal(b, m, deterministic)
}
func (dst *ListDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDevicesRequest.Merge(dst, src)
}
func (m *ListDevicesRequest) XXX_Size() int {
	return xxx_messageInfo_ListDevicesRequest.Size(m)
}
func (m *ListDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDevicesRequest proto.InternalMessageInfo

func (m *ListDevicesRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListDevicesRequest) GetProjectId() *wrappers.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *ListDevicesRequest) GetOwners() []*wrappers.StringValue {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *ListDevicesRequest) GetGroups() []*wrappers.StringValue {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ListDevicesRequest) GetState() state.DeviceState {
	if m != nil {
		return m.State
	}
	return state.DeviceState_DEVICE_STATE_UNKNOWN
}

type ListDevicesResponse struct {
	Devices              []*Device `protobuf:"bytes,1,rep,name=devices" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListDevicesResponse) Reset()         { *m = ListDevicesResponse{} }
func (m *ListDevicesResponse) String() string { return proto.CompactTextString(m) }
func (*ListDevicesResponse) ProtoMessage()    {}
func (*ListDevicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_devices_847511c5a6bf64d0, []int{1}
}
func (m *ListDevicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDevicesResponse.Unmarshal(m, b)
}
func (m *ListDevicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDevicesResponse.Marshal(b, m, deterministic)
}
func (dst *ListDevicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDevicesResponse.Merge(dst, src)
}
func (m *ListDevicesResponse) XXX_Size() int {
	return xxx_messageInfo_ListDevicesResponse.Size(m)
}
func (m *ListDevicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDevicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDevicesResponse proto.InternalMessageInfo

func (m *ListDevicesResponse) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func init() {
	proto.RegisterType((*ListDevicesRequest)(nil), "ai.metathings.service.deviced.ListDevicesRequest")
	proto.RegisterType((*ListDevicesResponse)(nil), "ai.metathings.service.deviced.ListDevicesResponse")
}

func init() { proto.RegisterFile("list_devices.proto", fileDescriptor_list_devices_847511c5a6bf64d0) }

var fileDescriptor_list_devices_847511c5a6bf64d0 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0x95, 0xfe, 0x7b, 0xf5, 0xfa, 0x7d, 0xc5, 0x60, 0x96, 0xa8, 0x02, 0x54, 0x55, 0x42,
	0x2a, 0x43, 0x6d, 0x54, 0x10, 0x0b, 0x43, 0x07, 0x58, 0x90, 0x98, 0x52, 0xa9, 0x6b, 0xe5, 0x26,
	0xc6, 0x35, 0x4d, 0x7c, 0xc1, 0xbe, 0x34, 0xe2, 0x5b, 0xf1, 0x11, 0x51, 0x63, 0x17, 0xda, 0x05,
	0xba, 0x5d, 0x74, 0xcf, 0xef, 0xf2, 0xdc, 0xe3, 0x23, 0x34, 0xd7, 0x0e, 0x17, 0x99, 0xdc, 0xe8,
	0x54, 0x3a, 0x56, 0x5a, 0x40, 0xa0, 0xe7, 0x42, 0xb3, 0x42, 0xa2, 0xc0, 0x95, 0x36, 0xca, 0x31,
	0x27, 0xed, 0xb6, 0xcb, 0xbc, 0x28, 0xeb, 0x5f, 0x28, 0x00, 0x95, 0x4b, 0xde, 0x88, 0x97, 0xd5,
	0x0b, 0xaf, 0xad, 0x28, 0x4b, 0x69, 0x03, 0xde, 0xbf, 0x53, 0x1a, 0x57, 0xd5, 0x92, 0xa5, 0x50,
	0xf0, 0xa2, 0xd6, 0xb8, 0x86, 0x9a, 0x2b, 0x18, 0x37, 0xcd, 0xf1, 0x46, 0xe4, 0x3a, 0x13, 0x08,
	0xd6, 0xf1, 0xaf, 0x32, 0x70, 0x0f, 0x7b, 0x9c, 0x11, 0xef, 0x80, 0x28, 0xf8, 0xb7, 0x0d, 0x5e,
	0xae, 0x95, 0xff, 0x25, 0x4f, 0xc1, 0x38, 0x14, 0x06, 0xb9, 0x43, 0x81, 0x92, 0x7b, 0x5f, 0x61,
	0xc8, 0xff, 0xfd, 0xaf, 0xe1, 0x47, 0x8b, 0xd0, 0x67, 0xed, 0xf0, 0xd1, 0xef, 0x97, 0xc8, 0xb7,
	0x4a, 0x3a, 0xa4, 0xd7, 0xa4, 0x63, 0x44, 0x21, 0xe3, 0x68, 0x10, 0x8d, 0xfe, 0x4d, 0xce, 0x98,
	0x5f, 0x88, 0xed, 0x16, 0x62, 0x33, 0xb4, 0xda, 0xa8, 0xb9, 0xc8, 0x2b, 0x99, 0x34, 0x4a, 0x7a,
	0x4f, 0x48, 0x69, 0xe1, 0x55, 0xa6, 0xb8, 0xd0, 0x59, 0xdc, 0x3a, 0x82, 0xfb, 0x1b, 0xf4, 0x4f,
	0x19, 0xbd, 0x25, 0x3d, 0xa8, 0x8d, 0xb4, 0x2e, 0x6e, 0x0f, 0xda, 0xbf, 0x82, 0x41, 0xbb, 0xa5,
	0x94, 0x85, 0xaa, 0x74, 0x71, 0xe7, 0x18, 0xca, 0x6b, 0xe9, 0x94, 0x74, 0x9b, 0x54, 0xe2, 0xee,
	0x20, 0x1a, 0x9d, 0x4c, 0xae, 0xd8, 0xe1, 0x5b, 0xee, 0xa2, 0x63, 0x8d, 0x88, 0xf9, 0x5c, 0x66,
	0xdb, 0x3a, 0xf1, 0xdc, 0x70, 0x4e, 0x4e, 0x0f, 0x12, 0x73, 0x25, 0x18, 0x27, 0xe9, 0x94, 0xfc,
	0x09, 0x47, 0x12, 0x47, 0x8d, 0x9d, 0x4b, 0xf6, 0xe3, 0x95, 0x84, 0xd1, 0xc9, 0x8e, 0x5a, 0xf6,
	0x1a, 0xdb, 0x37, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x66, 0x93, 0x1d, 0xee, 0x71, 0x02, 0x00,
	0x00,
}