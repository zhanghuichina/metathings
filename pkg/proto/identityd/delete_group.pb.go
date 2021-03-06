// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_group.proto

package identityd

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

type DeleteGroupRequest struct {
	GroupId              *wrappers.StringValue `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteGroupRequest) Reset()         { *m = DeleteGroupRequest{} }
func (m *DeleteGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupRequest) ProtoMessage()    {}
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_delete_group_a9157c4808291add, []int{0}
}
func (m *DeleteGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupRequest.Unmarshal(m, b)
}
func (m *DeleteGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupRequest.Merge(dst, src)
}
func (m *DeleteGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupRequest.Size(m)
}
func (m *DeleteGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupRequest proto.InternalMessageInfo

func (m *DeleteGroupRequest) GetGroupId() *wrappers.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteGroupRequest)(nil), "ai.metathings.service.identityd.DeleteGroupRequest")
}

func init() { proto.RegisterFile("delete_group.proto", fileDescriptor_delete_group_a9157c4808291add) }

var fileDescriptor_delete_group_a9157c4808291add = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xce, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x05, 0x60, 0xea, 0x61, 0x95, 0x7a, 0xcb, 0x49, 0x16, 0x71, 0x17, 0x4f, 0x5e, 0x76, 0x02,
	0x0a, 0x5e, 0x05, 0x11, 0xc4, 0x6b, 0x45, 0xaf, 0x25, 0x6d, 0xc6, 0x74, 0x30, 0xed, 0xc4, 0x64,
	0xd2, 0xe2, 0xaf, 0x15, 0xfc, 0x25, 0x62, 0x8a, 0xde, 0x06, 0xe6, 0xbd, 0xc7, 0x57, 0x2b, 0x8b,
	0x1e, 0x05, 0x5b, 0x17, 0x39, 0x07, 0x08, 0x91, 0x85, 0xd5, 0xce, 0x10, 0x8c, 0x28, 0x46, 0x06,
	0x9a, 0x5c, 0x82, 0x84, 0x71, 0xa6, 0x1e, 0x81, 0x2c, 0x4e, 0x42, 0xf2, 0x69, 0xb7, 0x17, 0x8e,
	0xd9, 0x79, 0xd4, 0x25, 0xde, 0xe5, 0x37, 0xbd, 0x44, 0x13, 0x02, 0xc6, 0xb4, 0x0e, 0x6c, 0x6f,
	0x1d, 0xc9, 0x90, 0x3b, 0xe8, 0x79, 0xd4, 0xe3, 0x42, 0xf2, 0xce, 0x8b, 0x76, 0x7c, 0x28, 0xcf,
	0xc3, 0x6c, 0x3c, 0x59, 0x23, 0x1c, 0x93, 0xfe, 0x3f, 0xd7, 0xde, 0xe5, 0x4b, 0xad, 0x1e, 0x0a,
	0xe7, 0xf1, 0x57, 0xd3, 0xe0, 0x47, 0xc6, 0x24, 0xea, 0xae, 0x3e, 0x29, 0xba, 0x96, 0xec, 0x59,
	0xb5, 0xaf, 0xae, 0x4e, 0xaf, 0xcf, 0x61, 0x05, 0xc0, 0x1f, 0x00, 0x9e, 0x25, 0xd2, 0xe4, 0x5e,
	0x8d, 0xcf, 0x78, 0xbf, 0xf9, 0xfe, 0xda, 0x1d, 0xed, 0xab, 0xe6, 0xb8, 0xb4, 0x9e, 0x6c, 0xb7,
	0x29, 0xb1, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x21, 0xe0, 0x08, 0xec, 0x00, 0x00,
	0x00,
}
