// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_project.proto

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

type CreateProjectRequest struct {
	DomainId    *google_protobuf.StringValue   `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	ParentId    *google_protobuf.StringValue   `protobuf:"bytes,2,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	Enabled     *google_protobuf.BoolValue     `protobuf:"bytes,3,opt,name=enabled" json:"enabled,omitempty"`
	Name        *google_protobuf.StringValue   `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Description *google_protobuf.StringValue   `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Tags        []*google_protobuf.StringValue `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
}

func (m *CreateProjectRequest) Reset()                    { *m = CreateProjectRequest{} }
func (m *CreateProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateProjectRequest) ProtoMessage()               {}
func (*CreateProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *CreateProjectRequest) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *CreateProjectRequest) GetParentId() *google_protobuf.StringValue {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (m *CreateProjectRequest) GetEnabled() *google_protobuf.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *CreateProjectRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateProjectRequest) GetDescription() *google_protobuf.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CreateProjectRequest) GetTags() []*google_protobuf.StringValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreateProjectResponse struct {
	Project *Project `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
}

func (m *CreateProjectResponse) Reset()                    { *m = CreateProjectResponse{} }
func (m *CreateProjectResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateProjectResponse) ProtoMessage()               {}
func (*CreateProjectResponse) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *CreateProjectResponse) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateProjectRequest)(nil), "ai.metathings.service.identity.CreateProjectRequest")
	proto.RegisterType((*CreateProjectResponse)(nil), "ai.metathings.service.identity.CreateProjectResponse")
}

func init() { proto.RegisterFile("create_project.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd2, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0x07, 0x70, 0xf6, 0xe7, 0xb7, 0xfd, 0x96, 0xe1, 0xa5, 0x4c, 0x28, 0x43, 0xe6, 0xd8, 0x45,
	0x2f, 0x4b, 0x45, 0x45, 0xf4, 0x22, 0x38, 0x4f, 0xbb, 0x49, 0x05, 0x0f, 0x5e, 0x46, 0xda, 0x3c,
	0x76, 0xd1, 0x36, 0x4f, 0x4d, 0x9e, 0x6e, 0xf8, 0x6a, 0x05, 0xdf, 0x85, 0x37, 0x59, 0xb3, 0x8a,
	0x7f, 0xc0, 0xed, 0x56, 0x9a, 0xef, 0x27, 0xf9, 0xf2, 0x24, 0xac, 0x17, 0x1b, 0x10, 0x04, 0xb3,
	0xdc, 0xe0, 0x23, 0xc4, 0xc4, 0x73, 0x83, 0x84, 0xde, 0x40, 0x28, 0x9e, 0x01, 0x09, 0x9a, 0x2b,
	0x9d, 0x58, 0x6e, 0xc1, 0x2c, 0x54, 0x0c, 0x5c, 0x49, 0xd0, 0xa4, 0xe8, 0xa5, 0x3f, 0x48, 0x10,
	0x93, 0x14, 0x82, 0x32, 0x1d, 0x15, 0x0f, 0xc1, 0xd2, 0x88, 0x3c, 0x07, 0x63, 0x9d, 0xef, 0x9f,
	0x25, 0x8a, 0xe6, 0x45, 0xc4, 0x63, 0xcc, 0x82, 0x6c, 0xa9, 0xe8, 0x09, 0x97, 0x41, 0x82, 0xe3,
	0x72, 0x71, 0xbc, 0x10, 0xa9, 0x92, 0x82, 0xd0, 0xd8, 0xe0, 0xf3, 0x73, 0xed, 0x76, 0xbe, 0xd5,
	0x18, 0xbd, 0xd7, 0x59, 0xef, 0xba, 0xec, 0x77, 0xe3, 0xfe, 0x87, 0xf0, 0x5c, 0x80, 0x25, 0xef,
	0x82, 0x75, 0x24, 0x66, 0x42, 0xe9, 0x99, 0x92, 0x7e, 0x6d, 0x58, 0x3b, 0xec, 0x1e, 0xef, 0x71,
	0xd7, 0x89, 0x57, 0x9d, 0xf8, 0x2d, 0x19, 0xa5, 0x93, 0x3b, 0x91, 0x16, 0x10, 0xfe, 0x77, 0xf1,
	0xa9, 0x5c, 0xd1, 0x5c, 0x18, 0xd0, 0xb4, 0xa2, 0xf5, 0x6d, 0xa8, 0x8b, 0x4f, 0xa5, 0x77, 0xca,
	0xda, 0xa0, 0x45, 0x94, 0x82, 0xf4, 0x1b, 0x25, 0xec, 0xff, 0x82, 0x13, 0xc4, 0xd4, 0xb1, 0x2a,
	0xea, 0x9d, 0xb3, 0xa6, 0x16, 0x19, 0xf8, 0xcd, 0xcd, 0x67, 0x4d, 0x5a, 0x6f, 0xaf, 0xfb, 0xf5,
	0x61, 0x2d, 0x2c, 0x85, 0x77, 0xc9, 0xba, 0x12, 0x6c, 0x6c, 0x54, 0x4e, 0x0a, 0xb5, 0xff, 0x6f,
	0x8b, 0xb2, 0x5f, 0x81, 0x77, 0xc4, 0x9a, 0x24, 0x12, 0xeb, 0xb7, 0x86, 0x8d, 0x8d, 0xb0, 0x4c,
	0x86, 0x1d, 0x65, 0x67, 0x6e, 0x52, 0xa3, 0x7b, 0xb6, 0xfb, 0x63, 0xf4, 0x36, 0x47, 0x6d, 0xc1,
	0xbb, 0x62, 0xed, 0xf5, 0x2d, 0xad, 0x27, 0x7f, 0xc0, 0xff, 0x7e, 0x2d, 0xbc, 0xda, 0xa1, 0x72,
	0x51, 0xab, 0xac, 0x70, 0xf2, 0x11, 0x00, 0x00, 0xff, 0xff, 0xca, 0x0b, 0xa3, 0xca, 0x7d, 0x02,
	0x00, 0x00,
}