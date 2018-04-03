// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_group.proto

package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PatchGroupRequest struct {
	GroupId     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	Name        *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *PatchGroupRequest) Reset()                    { *m = PatchGroupRequest{} }
func (m *PatchGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*PatchGroupRequest) ProtoMessage()               {}
func (*PatchGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

func (m *PatchGroupRequest) GetGroupId() *google_protobuf.StringValue {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *PatchGroupRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PatchGroupRequest) GetDescription() *google_protobuf.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

type PatchGroupResponse struct {
	Group *Group `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
}

func (m *PatchGroupResponse) Reset()                    { *m = PatchGroupResponse{} }
func (m *PatchGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*PatchGroupResponse) ProtoMessage()               {}
func (*PatchGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{1} }

func (m *PatchGroupResponse) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchGroupRequest)(nil), "ai.metathings.service.identity.PatchGroupRequest")
	proto.RegisterType((*PatchGroupResponse)(nil), "ai.metathings.service.identity.PatchGroupResponse")
}

func init() { proto.RegisterFile("patch_group.proto", fileDescriptor30) }

var fileDescriptor30 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0xd9, 0x7e, 0x9f, 0x55, 0xd3, 0x53, 0x73, 0x2a, 0x45, 0x6a, 0x29, 0x08, 0x5e, 0x9a,
	0x15, 0x05, 0x2f, 0x82, 0x82, 0x17, 0xf1, 0xa6, 0x2b, 0x78, 0x2d, 0xe9, 0x66, 0x4c, 0x07, 0xbb,
	0x99, 0x98, 0xcc, 0xb6, 0xf8, 0x33, 0xfd, 0x05, 0x82, 0xbf, 0x44, 0x9a, 0xad, 0xb2, 0x27, 0xf5,
	0x36, 0xcc, 0xbc, 0x0f, 0x3c, 0xef, 0x88, 0xbe, 0xd7, 0x5c, 0x2e, 0x66, 0x36, 0x50, 0xed, 0x95,
	0x0f, 0xc4, 0x24, 0x47, 0x1a, 0x55, 0x05, 0xac, 0x79, 0x81, 0xce, 0x46, 0x15, 0x21, 0xac, 0xb0,
	0x04, 0x85, 0x06, 0x1c, 0x23, 0xbf, 0x0e, 0x47, 0x96, 0xc8, 0x2e, 0x21, 0x4f, 0xe9, 0x79, 0xfd,
	0x94, 0xaf, 0x83, 0xf6, 0x1e, 0x42, 0x6c, 0xf8, 0xe1, 0xb9, 0x45, 0x5e, 0xd4, 0x73, 0x55, 0x52,
	0x95, 0x57, 0x6b, 0xe4, 0x67, 0x5a, 0xe7, 0x96, 0xa6, 0xe9, 0x38, 0x5d, 0xe9, 0x25, 0x1a, 0xcd,
	0x14, 0x62, 0xfe, 0x3d, 0x6e, 0xb9, 0x5e, 0x4b, 0x62, 0xf2, 0x96, 0x89, 0xfe, 0xdd, 0x46, 0xed,
	0x66, 0xb3, 0x2c, 0xe0, 0xa5, 0x86, 0xc8, 0xf2, 0x4a, 0xec, 0xa5, 0xd0, 0x0c, 0xcd, 0x20, 0x1b,
	0x67, 0xc7, 0xbd, 0xd3, 0x03, 0xd5, 0xd8, 0xa8, 0x2f, 0x1b, 0xf5, 0xc0, 0x01, 0x9d, 0x7d, 0xd4,
	0xcb, 0x1a, 0xae, 0xbb, 0x1f, 0xef, 0x87, 0x9d, 0x71, 0x56, 0xec, 0x26, 0xea, 0xd6, 0xc8, 0x13,
	0xf1, 0xdf, 0xe9, 0x0a, 0x06, 0x9d, 0xdf, 0xe1, 0x22, 0x25, 0xe5, 0xa5, 0xe8, 0x19, 0x88, 0x65,
	0x40, 0xcf, 0x48, 0x6e, 0xf0, 0xef, 0x0f, 0x60, 0x1b, 0x28, 0xf6, 0x0d, 0x55, 0x1a, 0xdd, 0x0c,
	0xcd, 0xe4, 0x5e, 0xc8, 0x76, 0xa5, 0xe8, 0xc9, 0x45, 0x90, 0x17, 0x62, 0x27, 0xd9, 0x6d, 0x0b,
	0x1d, 0xa9, 0x9f, 0xdf, 0xaf, 0x1a, 0xba, 0x61, 0xe6, 0xdd, 0x24, 0x70, 0xf6, 0x19, 0x00, 0x00,
	0xff, 0xff, 0xc0, 0x0c, 0x2a, 0xe2, 0xc7, 0x01, 0x00, 0x00,
}
