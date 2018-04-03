// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_domain.proto

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

type CreateDomainRequest struct {
	Name        *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Enabled     *google_protobuf.BoolValue   `protobuf:"bytes,3,opt,name=enabled" json:"enabled,omitempty"`
}

func (m *CreateDomainRequest) Reset()                    { *m = CreateDomainRequest{} }
func (m *CreateDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDomainRequest) ProtoMessage()               {}
func (*CreateDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *CreateDomainRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateDomainRequest) GetDescription() *google_protobuf.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CreateDomainRequest) GetEnabled() *google_protobuf.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

type CreateDomainResponse struct {
	Domain *Domain `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
}

func (m *CreateDomainResponse) Reset()                    { *m = CreateDomainResponse{} }
func (m *CreateDomainResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDomainResponse) ProtoMessage()               {}
func (*CreateDomainResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *CreateDomainResponse) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDomainRequest)(nil), "ai.metathings.service.identity.CreateDomainRequest")
	proto.RegisterType((*CreateDomainResponse)(nil), "ai.metathings.service.identity.CreateDomainResponse")
}

func init() { proto.RegisterFile("create_domain.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0x3a, 0xa5, 0x42, 0xe6, 0xa9, 0xf3, 0x50, 0x8a, 0xcc, 0xb1, 0x83, 0x78, 0x59,
	0x0a, 0x2a, 0xe2, 0xc9, 0xc3, 0xf4, 0x13, 0x54, 0xd8, 0x55, 0xd2, 0xf6, 0x99, 0x3d, 0x6c, 0xf3,
	0x6a, 0xf2, 0xba, 0xe2, 0x07, 0xf4, 0x73, 0x08, 0x7e, 0x12, 0x31, 0xed, 0x64, 0x22, 0xb8, 0x5b,
	0x48, 0xde, 0x2f, 0xf9, 0xff, 0x23, 0x26, 0x85, 0x05, 0xc5, 0xf0, 0x54, 0x52, 0xad, 0xd0, 0xc8,
	0xc6, 0x12, 0x53, 0x34, 0x55, 0x28, 0x6b, 0x60, 0xc5, 0x6b, 0x34, 0xda, 0x49, 0x07, 0x76, 0x83,
	0x05, 0x48, 0x2c, 0xc1, 0x30, 0xf2, 0x5b, 0x32, 0xd5, 0x44, 0xba, 0x82, 0xd4, 0x4f, 0xe7, 0xed,
	0x73, 0xda, 0x59, 0xd5, 0x34, 0x60, 0x5d, 0xef, 0x93, 0x1b, 0x8d, 0xbc, 0x6e, 0x73, 0x59, 0x50,
	0x9d, 0xd6, 0x1d, 0xf2, 0x0b, 0x75, 0xa9, 0xa6, 0x85, 0x3f, 0x5c, 0x6c, 0x54, 0x85, 0xa5, 0x62,
	0xb2, 0x2e, 0xfd, 0x59, 0x0e, 0xee, 0x78, 0x37, 0xc5, 0xfc, 0x3d, 0x10, 0x93, 0x7b, 0x9f, 0xee,
	0xc1, 0x6f, 0x67, 0xf0, 0xda, 0x82, 0xe3, 0xe8, 0x56, 0x1c, 0x1a, 0x55, 0x43, 0x1c, 0xcc, 0x82,
	0x8b, 0xf1, 0xe5, 0xa9, 0xec, 0xc3, 0xc8, 0x6d, 0x18, 0xf9, 0xc8, 0x16, 0x8d, 0x5e, 0xa9, 0xaa,
	0x85, 0x65, 0xf8, 0xf9, 0x71, 0x36, 0x9a, 0x05, 0x99, 0x17, 0xd1, 0x9d, 0x18, 0x97, 0xe0, 0x0a,
	0x8b, 0x0d, 0x23, 0x99, 0x78, 0xb4, 0xff, 0x82, 0x6c, 0x17, 0x44, 0xd7, 0xe2, 0x08, 0x8c, 0xca,
	0x2b, 0x28, 0xe3, 0x03, 0x6f, 0x93, 0x3f, 0x76, 0x49, 0x54, 0xf5, 0x72, 0x3b, 0x3a, 0x5f, 0x89,
	0x93, 0xdf, 0x35, 0x5c, 0x43, 0xc6, 0x7d, 0xa7, 0x09, 0xfb, 0xbe, 0x43, 0x93, 0x73, 0xf9, 0xff,
	0xb7, 0xcb, 0xc1, 0x0f, 0x2a, 0x0f, 0xfd, 0xa3, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a,
	0xbc, 0x7e, 0x2d, 0xc3, 0x01, 0x00, 0x00,
}
