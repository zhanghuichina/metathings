// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_domain.proto

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

type GetDomainRequest struct {
	DomainId             *wrappers.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetDomainRequest) Reset()         { *m = GetDomainRequest{} }
func (m *GetDomainRequest) String() string { return proto.CompactTextString(m) }
func (*GetDomainRequest) ProtoMessage()    {}
func (*GetDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_domain_4e5e89cb5e9d145e, []int{0}
}
func (m *GetDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDomainRequest.Unmarshal(m, b)
}
func (m *GetDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDomainRequest.Marshal(b, m, deterministic)
}
func (dst *GetDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDomainRequest.Merge(dst, src)
}
func (m *GetDomainRequest) XXX_Size() int {
	return xxx_messageInfo_GetDomainRequest.Size(m)
}
func (m *GetDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDomainRequest proto.InternalMessageInfo

func (m *GetDomainRequest) GetDomainId() *wrappers.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

type GetDomainResponse struct {
	Domain               *Domain  `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDomainResponse) Reset()         { *m = GetDomainResponse{} }
func (m *GetDomainResponse) String() string { return proto.CompactTextString(m) }
func (*GetDomainResponse) ProtoMessage()    {}
func (*GetDomainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_domain_4e5e89cb5e9d145e, []int{1}
}
func (m *GetDomainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDomainResponse.Unmarshal(m, b)
}
func (m *GetDomainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDomainResponse.Marshal(b, m, deterministic)
}
func (dst *GetDomainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDomainResponse.Merge(dst, src)
}
func (m *GetDomainResponse) XXX_Size() int {
	return xxx_messageInfo_GetDomainResponse.Size(m)
}
func (m *GetDomainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDomainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDomainResponse proto.InternalMessageInfo

func (m *GetDomainResponse) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDomainRequest)(nil), "ai.metathings.service.identityd.GetDomainRequest")
	proto.RegisterType((*GetDomainResponse)(nil), "ai.metathings.service.identityd.GetDomainResponse")
}

func init() { proto.RegisterFile("get_domain.proto", fileDescriptor_get_domain_4e5e89cb5e9d145e) }

var fileDescriptor_get_domain_4e5e89cb5e9d145e = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x87, 0xa0, 0xab, 0x87, 0x9a, 0x93, 0x14, 0xb1, 0xa5, 0x17, 0xbd, 0x74, 0x16,
	0x14, 0xbc, 0x8a, 0x22, 0x88, 0xd7, 0xfa, 0x71, 0x2d, 0x9b, 0xee, 0xb8, 0x1d, 0x4c, 0x76, 0xe2,
	0xee, 0xa4, 0xc1, 0x5f, 0x2b, 0xf8, 0x4b, 0x84, 0x6c, 0x94, 0xdc, 0xbc, 0x2d, 0x3b, 0xef, 0xfb,
	0xcc, 0xc3, 0xa8, 0x89, 0x43, 0x59, 0x5b, 0xae, 0x0d, 0x79, 0x68, 0x02, 0x0b, 0x17, 0x33, 0x43,
	0x50, 0xa3, 0x18, 0xd9, 0x92, 0x77, 0x11, 0x22, 0x86, 0x1d, 0x6d, 0x10, 0xc8, 0xa2, 0x17, 0x92,
	0x4f, 0x3b, 0x3d, 0x73, 0xcc, 0xae, 0x42, 0xdd, 0xc7, 0xcb, 0xf6, 0x4d, 0x77, 0xc1, 0x34, 0x0d,
	0x86, 0x98, 0x00, 0xd3, 0x6b, 0x47, 0xb2, 0x6d, 0x4b, 0xd8, 0x70, 0xad, 0xeb, 0x8e, 0xe4, 0x9d,
	0x3b, 0xed, 0x78, 0xd9, 0x0f, 0x97, 0x3b, 0x53, 0x91, 0x35, 0xc2, 0x21, 0xea, 0xbf, 0xe7, 0xd0,
	0x3b, 0x1a, 0x6b, 0x2c, 0x5e, 0xd4, 0xe4, 0x01, 0xe5, 0xbe, 0xff, 0x5a, 0xe1, 0x47, 0x8b, 0x51,
	0x8a, 0x5b, 0x75, 0x90, 0x32, 0x6b, 0xb2, 0x27, 0xd9, 0x3c, 0xbb, 0x38, 0xbc, 0x3c, 0x85, 0x64,
	0x03, 0xbf, 0x36, 0xf0, 0x24, 0x81, 0xbc, 0x7b, 0x35, 0x55, 0x8b, 0x77, 0xf9, 0xf7, 0xd7, 0x6c,
	0x6f, 0x9e, 0xad, 0xf6, 0x53, 0xed, 0xd1, 0x2e, 0x9e, 0xd5, 0xf1, 0x08, 0x1b, 0x1b, 0xf6, 0x11,
	0x8b, 0x1b, 0x95, 0xa7, 0xc0, 0x00, 0x3d, 0x87, 0x7f, 0x6e, 0x00, 0x03, 0x60, 0xa8, 0x95, 0x79,
	0xbf, 0xfd, 0xea, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x3d, 0xb3, 0xd5, 0x4e, 0x01, 0x00, 0x00,
}
