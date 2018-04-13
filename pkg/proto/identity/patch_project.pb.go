// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_project.proto

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

type PatchProjectRequest struct {
	Id          *google_protobuf.StringValue   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Enabled     *google_protobuf.BoolValue     `protobuf:"bytes,2,opt,name=enabled" json:"enabled,omitempty"`
	Name        *google_protobuf.StringValue   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Description *google_protobuf.StringValue   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Tags        []*google_protobuf.StringValue `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
}

func (m *PatchProjectRequest) Reset()                    { *m = PatchProjectRequest{} }
func (m *PatchProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*PatchProjectRequest) ProtoMessage()               {}
func (*PatchProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor53, []int{0} }

func (m *PatchProjectRequest) GetId() *google_protobuf.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PatchProjectRequest) GetEnabled() *google_protobuf.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *PatchProjectRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PatchProjectRequest) GetDescription() *google_protobuf.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *PatchProjectRequest) GetTags() []*google_protobuf.StringValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

type PatchProjectResponse struct {
	Project *Project `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
}

func (m *PatchProjectResponse) Reset()                    { *m = PatchProjectResponse{} }
func (m *PatchProjectResponse) String() string            { return proto.CompactTextString(m) }
func (*PatchProjectResponse) ProtoMessage()               {}
func (*PatchProjectResponse) Descriptor() ([]byte, []int) { return fileDescriptor53, []int{1} }

func (m *PatchProjectResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchProjectRequest)(nil), "ai.metathings.service.identity.PatchProjectRequest")
	proto.RegisterType((*PatchProjectResponse)(nil), "ai.metathings.service.identity.PatchProjectResponse")
}

func init() { proto.RegisterFile("patch_project.proto", fileDescriptor53) }

var fileDescriptor53 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x69, 0xda, 0xdb, 0xd2, 0x29, 0x77, 0x93, 0xde, 0x45, 0x28, 0x97, 0x5a, 0xba, 0xd1,
	0x4d, 0x27, 0xa2, 0xc5, 0xa5, 0x60, 0x9f, 0xa0, 0x44, 0x10, 0x5c, 0x95, 0x49, 0xe6, 0x98, 0x1e,
	0x4d, 0xe6, 0xc4, 0x99, 0x93, 0x16, 0x9f, 0xcd, 0x87, 0x11, 0x7c, 0x12, 0x69, 0xd2, 0x88, 0x45,
	0xb0, 0xae, 0x66, 0x98, 0xf9, 0xbf, 0x39, 0x1f, 0xff, 0x88, 0x61, 0xa1, 0x38, 0x59, 0xaf, 0x0a,
	0x4b, 0x8f, 0x90, 0xb0, 0x2c, 0x2c, 0x31, 0xf9, 0x63, 0x85, 0x32, 0x07, 0x56, 0xbc, 0x46, 0x93,
	0x3a, 0xe9, 0xc0, 0x6e, 0x30, 0x01, 0x89, 0x1a, 0x0c, 0x23, 0xbf, 0x8c, 0xc6, 0x29, 0x51, 0x9a,
	0x41, 0x58, 0xa5, 0xe3, 0xf2, 0x21, 0xdc, 0x5a, 0x55, 0x14, 0x60, 0x5d, 0xcd, 0x8f, 0xae, 0x52,
	0xe4, 0x75, 0x19, 0xcb, 0x84, 0xf2, 0x30, 0xdf, 0x22, 0x3f, 0xd1, 0x36, 0x4c, 0x69, 0x56, 0x5d,
	0xce, 0x36, 0x2a, 0x43, 0xad, 0x98, 0xac, 0x0b, 0x3f, 0xb7, 0x7b, 0xee, 0xef, 0x81, 0xc6, 0xf4,
	0xd5, 0x13, 0xc3, 0xe5, 0x4e, 0x6f, 0x59, 0x1f, 0x47, 0xf0, 0x5c, 0x82, 0x63, 0x7f, 0x2e, 0x3c,
	0xd4, 0x41, 0x6b, 0xd2, 0x3a, 0x1b, 0x5c, 0xfc, 0x97, 0xb5, 0x8b, 0x6c, 0x5c, 0xe4, 0x2d, 0x5b,
	0x34, 0xe9, 0x9d, 0xca, 0x4a, 0x58, 0x74, 0xdf, 0xdf, 0x4e, 0xbc, 0x49, 0x2b, 0xf2, 0x50, 0xfb,
	0x73, 0xd1, 0x03, 0xa3, 0xe2, 0x0c, 0x74, 0xe0, 0x55, 0xe8, 0xe8, 0x1b, 0xba, 0x20, 0xca, 0x2a,
	0x30, 0x6a, 0xa2, 0xfe, 0xb9, 0xe8, 0x18, 0x95, 0x43, 0xd0, 0x3e, 0x3e, 0x2d, 0xaa, 0x92, 0xfe,
	0xb5, 0x18, 0x68, 0x70, 0x89, 0xc5, 0x82, 0x91, 0x4c, 0xd0, 0xf9, 0x05, 0xf8, 0x15, 0xd8, 0x4d,
	0x64, 0x95, 0xba, 0xe0, 0xcf, 0xa4, 0x7d, 0x7c, 0xe2, 0x2e, 0x19, 0xf5, 0xd1, 0xad, 0x34, 0xe5,
	0x0a, 0x4d, 0xd4, 0xaf, 0xd7, 0x15, 0xea, 0xe9, 0xbd, 0xf8, 0x77, 0x58, 0x9e, 0x2b, 0xc8, 0x38,
	0xf0, 0x6f, 0x44, 0x6f, 0x5f, 0xf3, 0xbe, 0xc2, 0x53, 0xf9, 0xf3, 0x77, 0xcb, 0xe6, 0x85, 0x86,
	0x8b, 0xbb, 0x95, 0xcc, 0xe5, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x03, 0x84, 0x87, 0x3d,
	0x02, 0x00, 0x00,
}