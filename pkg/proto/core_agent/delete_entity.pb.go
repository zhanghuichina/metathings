// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_entity.proto

package core_agent

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type DeleteEntityRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteEntityRequest) Reset()         { *m = DeleteEntityRequest{} }
func (m *DeleteEntityRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEntityRequest) ProtoMessage()    {}
func (*DeleteEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_delete_entity_450550a9c1a15b61, []int{0}
}
func (m *DeleteEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntityRequest.Unmarshal(m, b)
}
func (m *DeleteEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntityRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntityRequest.Merge(dst, src)
}
func (m *DeleteEntityRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEntityRequest.Size(m)
}
func (m *DeleteEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntityRequest proto.InternalMessageInfo

func (m *DeleteEntityRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteEntityRequest)(nil), "ai.metathings.service.core_agent.DeleteEntityRequest")
}

func init() { proto.RegisterFile("delete_entity.proto", fileDescriptor_delete_entity_450550a9c1a15b61) }

var fileDescriptor_delete_entity_450550a9c1a15b61 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xb1, 0x4a, 0xc6, 0x30,
	0x10, 0x80, 0x69, 0x87, 0x7f, 0xa8, 0xdb, 0xff, 0x2f, 0x52, 0x44, 0x8b, 0x93, 0x4b, 0x2f, 0xa0,
	0xe2, 0x03, 0x88, 0x4e, 0x6e, 0x15, 0x5c, 0x4b, 0xda, 0x9c, 0xe9, 0x61, 0x9a, 0xab, 0xc9, 0xa5,
	0xc5, 0xa7, 0x15, 0x7c, 0x12, 0x21, 0x45, 0xb7, 0x83, 0xfb, 0xee, 0xbe, 0xaf, 0x3a, 0x19, 0x74,
	0x28, 0xd8, 0xa3, 0x17, 0x92, 0x2f, 0x58, 0x02, 0x0b, 0x1f, 0x1b, 0x4d, 0x30, 0xa3, 0x68, 0x99,
	0xc8, 0xdb, 0x08, 0x11, 0xc3, 0x4a, 0x23, 0xc2, 0xc8, 0x01, 0x7b, 0x6d, 0xd1, 0x4b, 0x7d, 0x69,
	0x99, 0xad, 0x43, 0x95, 0xf9, 0x21, 0xbd, 0xab, 0x2d, 0xe8, 0x65, 0xc1, 0x10, 0xf7, 0x0f, 0xf5,
	0x83, 0x25, 0x99, 0xd2, 0x00, 0x23, 0xcf, 0x6a, 0xde, 0x48, 0x3e, 0x78, 0x53, 0x96, 0xdb, 0xbc,
	0x6c, 0x57, 0xed, 0xc8, 0x68, 0xe1, 0x10, 0xd5, 0xff, 0xb8, 0xdf, 0x5d, 0xbf, 0x54, 0xa7, 0xa7,
	0x1c, 0xf4, 0x9c, 0x7b, 0x3a, 0xfc, 0x4c, 0x18, 0xe5, 0x78, 0x5f, 0x95, 0x64, 0xce, 0x8b, 0xa6,
	0xb8, 0x39, 0xbb, 0xbd, 0x80, 0xdd, 0x0d, 0x7f, 0x6e, 0x78, 0x95, 0x40, 0xde, 0xbe, 0x69, 0x97,
	0xf0, 0xf1, 0xf0, 0xf3, 0x7d, 0x55, 0x36, 0x45, 0x57, 0x92, 0x19, 0x0e, 0x99, 0xb8, 0xfb, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0xeb, 0x2a, 0x0f, 0xe4, 0x00, 0x00, 0x00,
}
