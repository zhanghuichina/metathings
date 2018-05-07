// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_domain.proto

package identity

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

type PatchDomainRequest struct {
	DomainId             *wrappers.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description          *wrappers.StringValue `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Enabled              *wrappers.BoolValue   `protobuf:"bytes,4,opt,name=enabled" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PatchDomainRequest) Reset()         { *m = PatchDomainRequest{} }
func (m *PatchDomainRequest) String() string { return proto.CompactTextString(m) }
func (*PatchDomainRequest) ProtoMessage()    {}
func (*PatchDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_patch_domain_3c8fad9d4314d7e8, []int{0}
}
func (m *PatchDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchDomainRequest.Unmarshal(m, b)
}
func (m *PatchDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchDomainRequest.Marshal(b, m, deterministic)
}
func (dst *PatchDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchDomainRequest.Merge(dst, src)
}
func (m *PatchDomainRequest) XXX_Size() int {
	return xxx_messageInfo_PatchDomainRequest.Size(m)
}
func (m *PatchDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchDomainRequest proto.InternalMessageInfo

func (m *PatchDomainRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *PatchDomainRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PatchDomainRequest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *PatchDomainRequest) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

type PatchDomainResponse struct {
	Domain               *Domain  `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchDomainResponse) Reset()         { *m = PatchDomainResponse{} }
func (m *PatchDomainResponse) String() string { return proto.CompactTextString(m) }
func (*PatchDomainResponse) ProtoMessage()    {}
func (*PatchDomainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_patch_domain_3c8fad9d4314d7e8, []int{1}
}
func (m *PatchDomainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchDomainResponse.Unmarshal(m, b)
}
func (m *PatchDomainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchDomainResponse.Marshal(b, m, deterministic)
}
func (dst *PatchDomainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchDomainResponse.Merge(dst, src)
}
func (m *PatchDomainResponse) XXX_Size() int {
	return xxx_messageInfo_PatchDomainResponse.Size(m)
}
func (m *PatchDomainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchDomainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PatchDomainResponse proto.InternalMessageInfo

func (m *PatchDomainResponse) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchDomainRequest)(nil), "ai.metathings.service.identity.PatchDomainRequest")
	proto.RegisterType((*PatchDomainResponse)(nil), "ai.metathings.service.identity.PatchDomainResponse")
}

func init() { proto.RegisterFile("patch_domain.proto", fileDescriptor_patch_domain_3c8fad9d4314d7e8) }

var fileDescriptor_patch_domain_3c8fad9d4314d7e8 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0x07, 0x70, 0xba, 0xdf, 0xd8, 0x4f, 0x33, 0x4f, 0xf1, 0x52, 0x86, 0xcc, 0xb1, 0x83, 0x78,
	0x59, 0x2a, 0x2a, 0x1e, 0x05, 0x87, 0x17, 0x6f, 0x32, 0xd1, 0xeb, 0x48, 0x9b, 0x67, 0xf7, 0xb0,
	0xcd, 0xab, 0xc9, 0xeb, 0x86, 0x7f, 0xad, 0xe0, 0x3f, 0xa2, 0xac, 0xe9, 0x64, 0x22, 0xa8, 0xb7,
	0x90, 0x7c, 0x3f, 0xe4, 0xfb, 0x12, 0x21, 0x2b, 0xcd, 0xd9, 0x62, 0x6e, 0xa8, 0xd4, 0x68, 0x55,
	0xe5, 0x88, 0x49, 0x0e, 0x35, 0xaa, 0x12, 0x58, 0xf3, 0x02, 0x6d, 0xee, 0x95, 0x07, 0xb7, 0xc4,
	0x0c, 0x14, 0x1a, 0xb0, 0x8c, 0xfc, 0x32, 0x18, 0xe6, 0x44, 0x79, 0x01, 0x49, 0x93, 0x4e, 0xeb,
	0xc7, 0x64, 0xe5, 0x74, 0x55, 0x81, 0xf3, 0xc1, 0x0f, 0x2e, 0x72, 0xe4, 0x45, 0x9d, 0xaa, 0x8c,
	0xca, 0xa4, 0x5c, 0x21, 0x3f, 0xd1, 0x2a, 0xc9, 0x69, 0xd2, 0x1c, 0x4e, 0x96, 0xba, 0x40, 0xa3,
	0x99, 0x9c, 0x4f, 0x3e, 0x97, 0xad, 0xdb, 0xdb, 0x6e, 0x31, 0x7e, 0x8f, 0x84, 0xbc, 0x5d, 0x97,
	0xbb, 0x6e, 0x76, 0x67, 0xf0, 0x5c, 0x83, 0x67, 0x79, 0x25, 0x76, 0x43, 0x6c, 0x8e, 0x26, 0x8e,
	0x46, 0xd1, 0x71, 0xff, 0xf4, 0x40, 0x85, 0x42, 0x6a, 0x53, 0x48, 0xdd, 0xb1, 0x43, 0x9b, 0x3f,
	0xe8, 0xa2, 0x86, 0x69, 0xef, 0xed, 0xf5, 0xb0, 0x33, 0x8a, 0x66, 0x3b, 0x81, 0xdd, 0x18, 0x79,
	0x22, 0xba, 0x56, 0x97, 0x10, 0x77, 0x7e, 0xd7, 0xb3, 0x26, 0x29, 0x2f, 0x45, 0xdf, 0x80, 0xcf,
	0x1c, 0x56, 0x8c, 0x64, 0xe3, 0x7f, 0x7f, 0x80, 0xdb, 0x40, 0x9e, 0x8b, 0xff, 0x60, 0x75, 0x5a,
	0x80, 0x89, 0xbb, 0x8d, 0x1d, 0x7c, 0xb3, 0x53, 0xa2, 0x22, 0xc8, 0x4d, 0x74, 0x7c, 0x2f, 0xf6,
	0xbf, 0x3c, 0x80, 0xaf, 0xc8, 0xfa, 0x75, 0x99, 0x5e, 0x18, 0xa5, 0x1d, 0xff, 0x48, 0xfd, 0xfc,
	0x5f, 0xaa, 0xf5, 0xad, 0x4a, 0x7b, 0xcd, 0x9d, 0x67, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8c,
	0xe8, 0xb8, 0xcb, 0xfb, 0x01, 0x00, 0x00,
}
