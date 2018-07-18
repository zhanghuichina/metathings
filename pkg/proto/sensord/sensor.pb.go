// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sensor.proto

package sensord

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"
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

type Sensor struct {
	Id                      string                         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                    string                         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ApplicationCredentialId string                         `protobuf:"bytes,3,opt,name=application_credential_id,json=applicationCredentialId" json:"application_credential_id,omitempty"`
	CoreId                  string                         `protobuf:"bytes,4,opt,name=core_id,json=coreId" json:"core_id,omitempty"`
	EntityName              string                         `protobuf:"bytes,5,opt,name=entity_name,json=entityName" json:"entity_name,omitempty"`
	OwnerId                 string                         `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	Tags                    []string                       `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	State                   sensor.SensorState             `protobuf:"varint,8,opt,name=state,enum=ai.metathings.service.sensor.SensorState" json:"state,omitempty"`
	Config                  map[string]*sensor.SensorValue `protobuf:"bytes,9,rep,name=config" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral    struct{}                       `json:"-"`
	XXX_unrecognized        []byte                         `json:"-"`
	XXX_sizecache           int32                          `json:"-"`
}

func (m *Sensor) Reset()         { *m = Sensor{} }
func (m *Sensor) String() string { return proto.CompactTextString(m) }
func (*Sensor) ProtoMessage()    {}
func (*Sensor) Descriptor() ([]byte, []int) {
	return fileDescriptor_sensor_972ad8ba72db13c6, []int{0}
}
func (m *Sensor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sensor.Unmarshal(m, b)
}
func (m *Sensor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sensor.Marshal(b, m, deterministic)
}
func (dst *Sensor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sensor.Merge(dst, src)
}
func (m *Sensor) XXX_Size() int {
	return xxx_messageInfo_Sensor.Size(m)
}
func (m *Sensor) XXX_DiscardUnknown() {
	xxx_messageInfo_Sensor.DiscardUnknown(m)
}

var xxx_messageInfo_Sensor proto.InternalMessageInfo

func (m *Sensor) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Sensor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sensor) GetApplicationCredentialId() string {
	if m != nil {
		return m.ApplicationCredentialId
	}
	return ""
}

func (m *Sensor) GetCoreId() string {
	if m != nil {
		return m.CoreId
	}
	return ""
}

func (m *Sensor) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func (m *Sensor) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Sensor) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Sensor) GetState() sensor.SensorState {
	if m != nil {
		return m.State
	}
	return sensor.SensorState_SENSOR_STATE_UNKNOWN
}

func (m *Sensor) GetConfig() map[string]*sensor.SensorValue {
	if m != nil {
		return m.Config
	}
	return nil
}

type SensorData struct {
	CreatedAt            *timestamp.Timestamp           `protobuf:"bytes,1,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	ArrivedAt            *timestamp.Timestamp           `protobuf:"bytes,2,opt,name=arrived_at,json=arrivedAt" json:"arrived_at,omitempty"`
	Data                 map[string]*sensor.SensorValue `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SensorData) Reset()         { *m = SensorData{} }
func (m *SensorData) String() string { return proto.CompactTextString(m) }
func (*SensorData) ProtoMessage()    {}
func (*SensorData) Descriptor() ([]byte, []int) {
	return fileDescriptor_sensor_972ad8ba72db13c6, []int{1}
}
func (m *SensorData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorData.Unmarshal(m, b)
}
func (m *SensorData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorData.Marshal(b, m, deterministic)
}
func (dst *SensorData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorData.Merge(dst, src)
}
func (m *SensorData) XXX_Size() int {
	return xxx_messageInfo_SensorData.Size(m)
}
func (m *SensorData) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorData.DiscardUnknown(m)
}

var xxx_messageInfo_SensorData proto.InternalMessageInfo

func (m *SensorData) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *SensorData) GetArrivedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ArrivedAt
	}
	return nil
}

func (m *SensorData) GetData() map[string]*sensor.SensorValue {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Sensor)(nil), "ai.metathings.service.sensord.Sensor")
	proto.RegisterMapType((map[string]*sensor.SensorValue)(nil), "ai.metathings.service.sensord.Sensor.ConfigEntry")
	proto.RegisterType((*SensorData)(nil), "ai.metathings.service.sensord.SensorData")
	proto.RegisterMapType((map[string]*sensor.SensorValue)(nil), "ai.metathings.service.sensord.SensorData.DataEntry")
}

func init() { proto.RegisterFile("sensor.proto", fileDescriptor_sensor_972ad8ba72db13c6) }

var fileDescriptor_sensor_972ad8ba72db13c6 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0x6e, 0xdb, 0x30,
	0x10, 0x86, 0x61, 0xc9, 0x91, 0xe3, 0x53, 0x11, 0x14, 0x5c, 0xa2, 0x08, 0x68, 0x63, 0x64, 0x72,
	0x17, 0x0a, 0x75, 0x96, 0x36, 0x1d, 0x8a, 0x20, 0x2d, 0x0a, 0x2f, 0x1d, 0x94, 0xa2, 0xab, 0x71,
	0x16, 0x19, 0x85, 0x88, 0x45, 0x0a, 0xd4, 0xd9, 0x81, 0x9e, 0xaa, 0x0f, 0xd1, 0x17, 0x2b, 0x48,
	0x2a, 0xa9, 0xd1, 0x21, 0xcd, 0x92, 0x45, 0x20, 0xef, 0xfe, 0xef, 0xee, 0xf8, 0xeb, 0xe0, 0x55,
	0x27, 0x75, 0x67, 0x2c, 0x6f, 0xad, 0x21, 0xc3, 0xde, 0xa0, 0xe2, 0x8d, 0x24, 0xa4, 0x5b, 0xa5,
	0xeb, 0x8e, 0x77, 0xd2, 0xee, 0x54, 0x25, 0x79, 0xd0, 0x88, 0xfc, 0x6d, 0x6d, 0x4c, 0xbd, 0x91,
	0x85, 0x17, 0xaf, 0xb7, 0x37, 0xc5, 0xbd, 0xc5, 0xb6, 0x95, 0xb6, 0x0b, 0x78, 0x7e, 0xfa, 0x6f,
	0x9e, 0x54, 0x23, 0x3b, 0xc2, 0xa6, 0x1d, 0x04, 0x9f, 0x6a, 0x45, 0xb7, 0xdb, 0x35, 0xaf, 0x4c,
	0x53, 0x68, 0xec, 0x0d, 0x11, 0x16, 0x7f, 0xfb, 0x15, 0xed, 0x5d, 0x1d, 0xd8, 0x22, 0x74, 0x2c,
	0xf6, 0x87, 0x3b, 0xfb, 0x1d, 0x43, 0x72, 0xed, 0x03, 0xec, 0x08, 0x22, 0x25, 0xb2, 0xd1, 0x6c,
	0x34, 0x9f, 0x96, 0x91, 0x12, 0x8c, 0xc1, 0x58, 0x63, 0x23, 0xb3, 0xc8, 0x47, 0xfc, 0x99, 0x5d,
	0xc0, 0x09, 0xb6, 0xed, 0x46, 0x55, 0x48, 0xca, 0xe8, 0x55, 0x65, 0xa5, 0x90, 0x9a, 0x14, 0x6e,
	0x56, 0x4a, 0x64, 0xb1, 0x17, 0x1e, 0xef, 0x09, 0xae, 0x1e, 0xf3, 0x4b, 0xc1, 0x8e, 0x61, 0x52,
	0x19, 0x2b, 0x9d, 0x72, 0xec, 0x95, 0x89, 0xbb, 0x2e, 0x05, 0x3b, 0x85, 0xd4, 0x89, 0xa8, 0x5f,
	0xf9, 0x7e, 0x07, 0x3e, 0x09, 0x21, 0xf4, 0xdd, 0x75, 0x3d, 0x81, 0x43, 0x73, 0xaf, 0xa5, 0x75,
	0x68, 0xe2, 0xb3, 0x13, 0x7f, 0x5f, 0xfa, 0x21, 0x09, 0xeb, 0x2e, 0x9b, 0xcc, 0x62, 0x37, 0xa4,
	0x3b, 0xb3, 0xcf, 0x70, 0xd0, 0x11, 0x92, 0xcc, 0x0e, 0x67, 0xa3, 0xf9, 0xd1, 0xe2, 0x1d, 0x7f,
	0xea, 0x07, 0xf0, 0xf0, 0xfa, 0x6b, 0x07, 0x94, 0x81, 0x63, 0x4b, 0x48, 0x2a, 0xa3, 0x6f, 0x54,
	0x9d, 0x4d, 0x67, 0xf1, 0x3c, 0x5d, 0xbc, 0x7f, 0xb2, 0x82, 0x18, 0x4a, 0xf0, 0x2b, 0xcf, 0x7c,
	0xd5, 0x64, 0xfb, 0x72, 0x28, 0x90, 0x0b, 0x48, 0xf7, 0xc2, 0xec, 0x35, 0xc4, 0x77, 0xb2, 0x1f,
	0x4c, 0x76, 0x47, 0x37, 0xec, 0x0e, 0x37, 0xdb, 0x60, 0x73, 0xfa, 0xbc, 0x61, 0x7f, 0x3a, 0xa0,
	0x0c, 0xdc, 0x45, 0xf4, 0x61, 0x74, 0xf6, 0x2b, 0x02, 0x08, 0xa9, 0x2f, 0x48, 0xc8, 0x3e, 0x02,
	0x54, 0x56, 0x22, 0x49, 0xb1, 0x42, 0xf2, 0xcd, 0xd2, 0x45, 0xce, 0xc3, 0x1e, 0xf1, 0x87, 0x3d,
	0xe2, 0x3f, 0x1e, 0xf6, 0xa8, 0x9c, 0x0e, 0xea, 0x4b, 0x72, 0x28, 0x5a, 0xab, 0x76, 0x01, 0x8d,
	0xfe, 0x8f, 0x0e, 0xea, 0x4b, 0x62, 0xdf, 0x60, 0x2c, 0x90, 0x30, 0x8b, 0xbd, 0x67, 0xe7, 0xcf,
	0xf2, 0xcc, 0x8d, 0xcb, 0xdd, 0x27, 0xb8, 0xe6, 0x0b, 0xe4, 0x6b, 0x98, 0x3e, 0x86, 0x5e, 0xc8,
	0xb1, 0x75, 0xe2, 0xdf, 0x72, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xda, 0x02, 0x21, 0x88, 0xab,
	0x03, 0x00, 0x00,
}