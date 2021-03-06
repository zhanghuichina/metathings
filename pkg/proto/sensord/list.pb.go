// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list.proto

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

type ListRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CoreId               *wrappers.StringValue `protobuf:"bytes,2,opt,name=core_id,json=coreId" json:"core_id,omitempty"`
	EntityName           *wrappers.StringValue `protobuf:"bytes,3,opt,name=entity_name,json=entityName" json:"entity_name,omitempty"`
	OwnerId              *wrappers.StringValue `protobuf:"bytes,4,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	State                sensor.SensorState    `protobuf:"varint,5,opt,name=state,enum=ai.metathings.service.sensor.SensorState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_899572bc594d25bb, []int{0}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListRequest) GetCoreId() *wrappers.StringValue {
	if m != nil {
		return m.CoreId
	}
	return nil
}

func (m *ListRequest) GetEntityName() *wrappers.StringValue {
	if m != nil {
		return m.EntityName
	}
	return nil
}

func (m *ListRequest) GetOwnerId() *wrappers.StringValue {
	if m != nil {
		return m.OwnerId
	}
	return nil
}

func (m *ListRequest) GetState() sensor.SensorState {
	if m != nil {
		return m.State
	}
	return sensor.SensorState_SENSOR_STATE_UNKNOWN
}

type ListResponse struct {
	Sensors              []*Sensor `protobuf:"bytes,1,rep,name=sensors" json:"sensors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_899572bc594d25bb, []int{1}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetSensors() []*Sensor {
	if m != nil {
		return m.Sensors
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "ai.metathings.service.sensord.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "ai.metathings.service.sensord.ListResponse")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_list_899572bc594d25bb) }

var fileDescriptor_list_899572bc594d25bb = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x6b, 0xe3, 0x30,
	0x10, 0xc5, 0x71, 0xfe, 0x2e, 0x72, 0xd8, 0x83, 0x4f, 0x26, 0xec, 0x96, 0x10, 0x28, 0xa4, 0x87,
	0x48, 0x25, 0xa5, 0xed, 0xa1, 0x94, 0x9c, 0x03, 0xa5, 0x05, 0x07, 0x7a, 0x0d, 0x4a, 0x3c, 0x55,
	0x44, 0x6c, 0xc9, 0x95, 0xc6, 0x31, 0xf9, 0x54, 0xfd, 0x8a, 0xc5, 0x92, 0xd3, 0xe6, 0x14, 0x72,
	0x1a, 0x09, 0xbd, 0xdf, 0xbc, 0x37, 0x1a, 0x42, 0x32, 0x69, 0x91, 0x16, 0x46, 0xa3, 0x8e, 0xfe,
	0x73, 0x49, 0x73, 0x40, 0x8e, 0x5b, 0xa9, 0x84, 0xa5, 0x16, 0xcc, 0x5e, 0x6e, 0x80, 0x5a, 0x50,
	0x56, 0x9b, 0x74, 0x78, 0x25, 0xb4, 0x16, 0x19, 0x30, 0x27, 0x5e, 0x97, 0x1f, 0xac, 0x32, 0xbc,
	0x28, 0xc0, 0x58, 0x8f, 0x0f, 0x1f, 0x84, 0xc4, 0x6d, 0xb9, 0xa6, 0x1b, 0x9d, 0xb3, 0xbc, 0x92,
	0xb8, 0xd3, 0x15, 0x13, 0x7a, 0xea, 0x1e, 0xa7, 0x7b, 0x9e, 0xc9, 0x94, 0xa3, 0x36, 0x96, 0xfd,
	0x1c, 0x1b, 0xee, 0xe9, 0x84, 0x53, 0xfc, 0xa0, 0x11, 0x39, 0xfb, 0x8d, 0xc1, 0x8a, 0x9d, 0xf0,
	0x96, 0xcc, 0x07, 0x69, 0x4a, 0x03, 0x0f, 0x4e, 0x6f, 0xe3, 0xaf, 0x16, 0x09, 0x5f, 0xa4, 0xc5,
	0x04, 0x3e, 0x4b, 0xb0, 0x18, 0xdd, 0x92, 0x8e, 0xe2, 0x39, 0xc4, 0xc1, 0x28, 0x98, 0x84, 0xb3,
	0x7f, 0xd4, 0x4f, 0x40, 0x8f, 0x13, 0xd0, 0x25, 0x1a, 0xa9, 0xc4, 0x3b, 0xcf, 0x4a, 0x48, 0x9c,
	0x32, 0xba, 0x27, 0xfd, 0x8d, 0x36, 0xb0, 0x92, 0x69, 0xdc, 0xba, 0x00, 0xea, 0xd5, 0xe2, 0x45,
	0x1a, 0x3d, 0x93, 0x10, 0x14, 0x4a, 0x3c, 0xac, 0x9c, 0x5f, 0xfb, 0x02, 0x94, 0x78, 0xe0, 0xb5,
	0x76, 0x7d, 0x24, 0x7f, 0x74, 0xa5, 0xc0, 0xd4, 0xb6, 0x9d, 0x0b, 0xd8, 0xbe, 0x53, 0x2f, 0xd2,
	0x68, 0x4e, 0xba, 0x16, 0x39, 0x42, 0xdc, 0x1d, 0x05, 0x93, 0xbf, 0xb3, 0x1b, 0x7a, 0x6e, 0x85,
	0x74, 0xe9, 0xca, 0xb2, 0x06, 0x12, 0xcf, 0x8d, 0xdf, 0xc8, 0xc0, 0x7f, 0x98, 0x2d, 0xb4, 0xb2,
	0x10, 0xcd, 0x49, 0xdf, 0x8b, 0x6d, 0x1c, 0x8c, 0xda, 0x93, 0x70, 0x76, 0x7d, 0xb6, 0x65, 0xda,
	0xf4, 0x4c, 0x8e, 0xd4, 0xba, 0xe7, 0x02, 0xdf, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x8b,
	0xc9, 0xe5, 0x59, 0x02, 0x00, 0x00,
}
