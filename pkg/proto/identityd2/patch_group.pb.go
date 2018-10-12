// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_group.proto

package identityd2

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

type PatchGroupRequest struct {
	Id                   *wrappers.StringValue            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Alias                *wrappers.StringValue            `protobuf:"bytes,6,opt,name=alias" json:"alias,omitempty"`
	Description          *wrappers.StringValue            `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	Extra                map[string]*wrappers.StringValue `protobuf:"bytes,8,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PatchGroupRequest) Reset()         { *m = PatchGroupRequest{} }
func (m *PatchGroupRequest) String() string { return proto.CompactTextString(m) }
func (*PatchGroupRequest) ProtoMessage()    {}
func (*PatchGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_patch_group_1e560c53dca72d30, []int{0}
}
func (m *PatchGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchGroupRequest.Unmarshal(m, b)
}
func (m *PatchGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchGroupRequest.Marshal(b, m, deterministic)
}
func (dst *PatchGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchGroupRequest.Merge(dst, src)
}
func (m *PatchGroupRequest) XXX_Size() int {
	return xxx_messageInfo_PatchGroupRequest.Size(m)
}
func (m *PatchGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchGroupRequest proto.InternalMessageInfo

func (m *PatchGroupRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PatchGroupRequest) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *PatchGroupRequest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *PatchGroupRequest) GetExtra() map[string]*wrappers.StringValue {
	if m != nil {
		return m.Extra
	}
	return nil
}

type PatchGroupResponse struct {
	Group                *Group   `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchGroupResponse) Reset()         { *m = PatchGroupResponse{} }
func (m *PatchGroupResponse) String() string { return proto.CompactTextString(m) }
func (*PatchGroupResponse) ProtoMessage()    {}
func (*PatchGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_patch_group_1e560c53dca72d30, []int{1}
}
func (m *PatchGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchGroupResponse.Unmarshal(m, b)
}
func (m *PatchGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchGroupResponse.Marshal(b, m, deterministic)
}
func (dst *PatchGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchGroupResponse.Merge(dst, src)
}
func (m *PatchGroupResponse) XXX_Size() int {
	return xxx_messageInfo_PatchGroupResponse.Size(m)
}
func (m *PatchGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatchGroupResponse proto.InternalMessageInfo

func (m *PatchGroupResponse) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchGroupRequest)(nil), "ai.metathings.service.identityd2.PatchGroupRequest")
	proto.RegisterMapType((map[string]*wrappers.StringValue)(nil), "ai.metathings.service.identityd2.PatchGroupRequest.ExtraEntry")
	proto.RegisterType((*PatchGroupResponse)(nil), "ai.metathings.service.identityd2.PatchGroupResponse")
}

func init() { proto.RegisterFile("patch_group.proto", fileDescriptor_patch_group_1e560c53dca72d30) }

var fileDescriptor_patch_group_1e560c53dca72d30 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xb1, 0x33, 0x7b, 0x99, 0x7c, 0x59, 0x74, 0x32, 0x61, 0x6c, 0x26, 0x97, 0xe5, 0x12,
	0x19, 0xbc, 0x31, 0xc6, 0x60, 0x39, 0x0c, 0xc2, 0xae, 0x45, 0x29, 0xb9, 0x16, 0xc5, 0x7a, 0x75,
	0x1e, 0xb1, 0x25, 0x57, 0x92, 0x93, 0xe6, 0x2b, 0xf4, 0x4b, 0x16, 0xfa, 0x49, 0x8a, 0xed, 0xb4,
	0x0d, 0xf4, 0x90, 0xdc, 0x9e, 0xad, 0xf7, 0x7b, 0xfc, 0xfe, 0x4f, 0x22, 0xa3, 0x5a, 0xb8, 0x7c,
	0x73, 0x53, 0x18, 0xdd, 0xd4, 0xac, 0x36, 0xda, 0x69, 0x9a, 0x08, 0x64, 0x15, 0x38, 0xe1, 0x36,
	0xa8, 0x0a, 0xcb, 0x2c, 0x98, 0x1d, 0xe6, 0xc0, 0x50, 0x82, 0x72, 0xe8, 0x0e, 0x32, 0x1b, 0x7f,
	0x2d, 0xb4, 0x2e, 0x4a, 0x48, 0xbb, 0xfe, 0x75, 0x73, 0x9b, 0xee, 0x8d, 0xa8, 0x6b, 0x30, 0xb6,
	0x9f, 0x30, 0xfe, 0x55, 0xa0, 0xdb, 0x34, 0x6b, 0x96, 0xeb, 0x2a, 0xad, 0xf6, 0xe8, 0xb6, 0x7a,
	0x9f, 0x16, 0x7a, 0xd6, 0x1d, 0xce, 0x76, 0xa2, 0x44, 0x29, 0x9c, 0x36, 0x36, 0x7d, 0x2d, 0x8f,
	0x5c, 0x54, 0x69, 0x09, 0x65, 0xff, 0x31, 0x79, 0x18, 0x90, 0xd1, 0x55, 0x2b, 0xf7, 0xbf, 0x75,
	0xe3, 0x70, 0xd7, 0x80, 0x75, 0xf4, 0x27, 0xf1, 0x51, 0xc6, 0x5e, 0xe2, 0x4d, 0xa3, 0xec, 0x0b,
	0xeb, 0x3d, 0xd8, 0x8b, 0x07, 0x5b, 0x3a, 0x83, 0xaa, 0x58, 0x89, 0xb2, 0x81, 0x7f, 0xe1, 0xd3,
	0xe3, 0x37, 0x3f, 0xf1, 0xb8, 0x8f, 0x92, 0x66, 0x24, 0x10, 0x25, 0x0a, 0x1b, 0x87, 0xe7, 0x41,
	0xde, 0xb7, 0xd2, 0x39, 0x89, 0x24, 0xd8, 0xdc, 0x60, 0xed, 0x50, 0xab, 0xf8, 0xe3, 0x05, 0xe4,
	0x29, 0x40, 0xaf, 0x49, 0x00, 0xf7, 0xce, 0x88, 0x78, 0x98, 0x0c, 0xa6, 0x51, 0x36, 0x67, 0xe7,
	0xd6, 0xca, 0xde, 0xa5, 0x65, 0x8b, 0x76, 0xc0, 0x42, 0x39, 0x73, 0xe0, 0xfd, 0xb0, 0xf1, 0x8a,
	0x90, 0xb7, 0x9f, 0xf4, 0x33, 0x19, 0x6c, 0xe1, 0xd0, 0xad, 0xe3, 0x13, 0x6f, 0xcb, 0x36, 0xe9,
	0xae, 0x75, 0x89, 0xfd, 0x4b, 0x92, 0x76, 0xad, 0x7f, 0xfc, 0xdf, 0x1e, 0x0f, 0xa5, 0xae, 0x04,
	0x2a, 0x1e, 0x18, 0x5d, 0x82, 0xe5, 0xc3, 0x4e, 0x0a, 0xc1, 0xf2, 0x0f, 0x4a, 0x54, 0x30, 0x59,
	0x12, 0x7a, 0x6a, 0x67, 0x6b, 0xad, 0x2c, 0xd0, 0xbf, 0x24, 0xe8, 0x1e, 0xce, 0xf1, 0x3e, 0xbe,
	0x9f, 0x8f, 0xd8, 0xf3, 0x3d, 0xb5, 0x0e, 0x3b, 0xa9, 0x1f, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0x95, 0xe8, 0xff, 0x84, 0x02, 0x00, 0x00,
}