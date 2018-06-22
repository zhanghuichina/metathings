// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list.proto

package camerad

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import camera "github.com/nayotta/metathings/pkg/proto/camera"

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
	State                camera.CameraState    `protobuf:"varint,4,opt,name=state,enum=ai.metathings.service.camera.CameraState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_099dba37251eb15f, []int{0}
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

func (m *ListRequest) GetState() camera.CameraState {
	if m != nil {
		return m.State
	}
	return camera.CameraState_CAMERA_STATE_UNKNOWN
}

type ListResponse struct {
	Cameras              []*Camera `protobuf:"bytes,1,rep,name=cameras" json:"cameras,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_099dba37251eb15f, []int{1}
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

func (m *ListResponse) GetCameras() []*Camera {
	if m != nil {
		return m.Cameras
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "ai.metathings.service.camerad.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "ai.metathings.service.camerad.ListResponse")
}

func init() { proto.RegisterFile("list.proto", fileDescriptor_list_099dba37251eb15f) }

var fileDescriptor_list_099dba37251eb15f = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0xf3, 0x30,
	0x18, 0xc6, 0xe9, 0xb7, 0x7d, 0x1b, 0xa4, 0xc3, 0x43, 0x4f, 0x65, 0xa8, 0x8c, 0x81, 0x30, 0x0f,
	0x4b, 0x64, 0xa2, 0x17, 0x11, 0x0f, 0x9e, 0x04, 0x51, 0xe8, 0xc0, 0xeb, 0x78, 0xd7, 0xbe, 0x66,
	0x61, 0x6d, 0x53, 0x93, 0xb7, 0x1b, 0xfb, 0xc7, 0x3d, 0x4b, 0x93, 0x4e, 0x77, 0x1a, 0x3b, 0xbd,
	0x09, 0x79, 0x7e, 0xef, 0xf3, 0x3c, 0x84, 0xb1, 0x5c, 0x59, 0xe2, 0x95, 0xd1, 0xa4, 0xa3, 0x0b,
	0x50, 0xbc, 0x40, 0x02, 0x5a, 0xa9, 0x52, 0x5a, 0x6e, 0xd1, 0x6c, 0x54, 0x8a, 0x3c, 0x85, 0x02,
	0x0d, 0x64, 0xc3, 0x4b, 0xa9, 0xb5, 0xcc, 0x51, 0x38, 0xf1, 0xb2, 0xfe, 0x14, 0x5b, 0x03, 0x55,
	0x85, 0xc6, 0x7a, 0x7c, 0x78, 0x2f, 0x15, 0xad, 0xea, 0x25, 0x4f, 0x75, 0x21, 0x8a, 0xad, 0xa2,
	0xb5, 0xde, 0x0a, 0xa9, 0xa7, 0xee, 0x71, 0xba, 0x81, 0x5c, 0x65, 0x40, 0xda, 0x58, 0xf1, 0x7b,
	0x6c, 0xb9, 0x87, 0x03, 0xae, 0x84, 0x9d, 0x26, 0x02, 0xf1, 0x17, 0x43, 0x54, 0x6b, 0xe9, 0x2d,
	0x85, 0x0f, 0xd2, 0x8e, 0x16, 0x1e, 0x1c, 0xde, 0xc6, 0xdf, 0x01, 0x0b, 0x5f, 0x95, 0xa5, 0x04,
	0xbf, 0x6a, 0xb4, 0x14, 0xdd, 0xb0, 0x6e, 0x09, 0x05, 0xc6, 0xc1, 0x28, 0x98, 0x84, 0xb3, 0x73,
	0xee, 0x1b, 0xf0, 0x7d, 0x03, 0x3e, 0x27, 0xa3, 0x4a, 0xf9, 0x01, 0x79, 0x8d, 0x89, 0x53, 0x46,
	0x77, 0xac, 0x9f, 0x6a, 0x83, 0x0b, 0x95, 0xc5, 0xff, 0x4e, 0x80, 0x7a, 0x8d, 0xf8, 0x25, 0x8b,
	0x1e, 0x59, 0x88, 0x25, 0x29, 0xda, 0x2d, 0x9c, 0x5f, 0xe7, 0x04, 0x94, 0x79, 0xe0, 0xad, 0x71,
	0x7d, 0x62, 0xff, 0x2d, 0x01, 0x61, 0xdc, 0x1d, 0x05, 0x93, 0xb3, 0xd9, 0x35, 0x3f, 0xf6, 0x13,
	0xfc, 0xd9, 0x8d, 0x79, 0x03, 0x24, 0x9e, 0x1b, 0xbf, 0xb3, 0x81, 0xef, 0x6d, 0x2b, 0x5d, 0xda,
	0x66, 0x61, 0xdf, 0x8b, 0x6d, 0x1c, 0x8c, 0x3a, 0x93, 0x70, 0x76, 0x75, 0x74, 0x65, 0xd6, 0xee,
	0x4c, 0xf6, 0xd4, 0xb2, 0xe7, 0x32, 0xdf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x02, 0xc2, 0x23,
	0x46, 0x20, 0x02, 0x00, 0x00,
}
