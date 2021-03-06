// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_domains.proto

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

type ListDomainsRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Enabled              *wrappers.BoolValue   `protobuf:"bytes,2,opt,name=enabled" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListDomainsRequest) Reset()         { *m = ListDomainsRequest{} }
func (m *ListDomainsRequest) String() string { return proto.CompactTextString(m) }
func (*ListDomainsRequest) ProtoMessage()    {}
func (*ListDomainsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_domains_9e3a7ab1bf655566, []int{0}
}
func (m *ListDomainsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDomainsRequest.Unmarshal(m, b)
}
func (m *ListDomainsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDomainsRequest.Marshal(b, m, deterministic)
}
func (dst *ListDomainsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDomainsRequest.Merge(dst, src)
}
func (m *ListDomainsRequest) XXX_Size() int {
	return xxx_messageInfo_ListDomainsRequest.Size(m)
}
func (m *ListDomainsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDomainsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDomainsRequest proto.InternalMessageInfo

func (m *ListDomainsRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListDomainsRequest) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

type ListDomainsResponse struct {
	Domains              []*Domain `protobuf:"bytes,1,rep,name=domains" json:"domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListDomainsResponse) Reset()         { *m = ListDomainsResponse{} }
func (m *ListDomainsResponse) String() string { return proto.CompactTextString(m) }
func (*ListDomainsResponse) ProtoMessage()    {}
func (*ListDomainsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_list_domains_9e3a7ab1bf655566, []int{1}
}
func (m *ListDomainsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDomainsResponse.Unmarshal(m, b)
}
func (m *ListDomainsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDomainsResponse.Marshal(b, m, deterministic)
}
func (dst *ListDomainsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDomainsResponse.Merge(dst, src)
}
func (m *ListDomainsResponse) XXX_Size() int {
	return xxx_messageInfo_ListDomainsResponse.Size(m)
}
func (m *ListDomainsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDomainsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDomainsResponse proto.InternalMessageInfo

func (m *ListDomainsResponse) GetDomains() []*Domain {
	if m != nil {
		return m.Domains
	}
	return nil
}

func init() {
	proto.RegisterType((*ListDomainsRequest)(nil), "ai.metathings.service.identityd.ListDomainsRequest")
	proto.RegisterType((*ListDomainsResponse)(nil), "ai.metathings.service.identityd.ListDomainsResponse")
}

func init() { proto.RegisterFile("list_domains.proto", fileDescriptor_list_domains_9e3a7ab1bf655566) }

var fileDescriptor_list_domains_9e3a7ab1bf655566 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0x8a, 0x0b, 0x59, 0x4f, 0xf1, 0x52, 0x8a, 0xe8, 0xb2, 0x17, 0xf7, 0xb2, 0xa9,
	0xac, 0xe2, 0x5d, 0xf1, 0xe8, 0xa9, 0x82, 0x78, 0x93, 0x74, 0x3b, 0x66, 0x07, 0xd3, 0x4c, 0x4d,
	0xa6, 0x5b, 0x04, 0x7f, 0xbc, 0x90, 0xb4, 0x82, 0x78, 0xf0, 0x96, 0x30, 0xef, 0x9b, 0xf9, 0x78,
	0x42, 0x5a, 0x0c, 0xfc, 0xda, 0x50, 0xab, 0xd1, 0x05, 0xd5, 0x79, 0x62, 0x92, 0x17, 0x1a, 0x55,
	0x0b, 0xac, 0x79, 0x87, 0xce, 0x04, 0x15, 0xc0, 0xef, 0x71, 0x0b, 0x0a, 0x1b, 0x70, 0x8c, 0xfc,
	0xd9, 0x14, 0xe7, 0x86, 0xc8, 0x58, 0x28, 0x63, 0xbc, 0xee, 0xdf, 0xca, 0xc1, 0xeb, 0xae, 0x03,
	0x3f, 0x2e, 0x28, 0x6e, 0x0d, 0xf2, 0xae, 0xaf, 0xd5, 0x96, 0xda, 0xb2, 0x1d, 0x90, 0xdf, 0x69,
	0x28, 0x0d, 0xad, 0xe3, 0x70, 0xbd, 0xd7, 0x16, 0x1b, 0xcd, 0xe4, 0x43, 0xf9, 0xf3, 0x1c, 0xb9,
	0x93, 0xe4, 0x91, 0x7e, 0xcb, 0x2f, 0x21, 0x1f, 0x31, 0xf0, 0x43, 0x72, 0xab, 0xe0, 0xa3, 0x87,
	0xc0, 0xf2, 0x4a, 0x1c, 0x39, 0xdd, 0x42, 0x9e, 0x2d, 0xb2, 0xd5, 0x7c, 0x73, 0xa6, 0x92, 0x8a,
	0x9a, 0x54, 0xd4, 0x13, 0x7b, 0x74, 0xe6, 0x59, 0xdb, 0x1e, 0xaa, 0x98, 0x94, 0x37, 0x62, 0x06,
	0x4e, 0xd7, 0x16, 0x9a, 0xfc, 0x20, 0x42, 0xc5, 0x1f, 0xe8, 0x9e, 0xc8, 0x26, 0x64, 0x8a, 0x2e,
	0x5f, 0xc4, 0xe9, 0xaf, 0xeb, 0xa1, 0x23, 0x17, 0x40, 0xde, 0x89, 0xd9, 0x58, 0x56, 0x9e, 0x2d,
	0x0e, 0x57, 0xf3, 0xcd, 0xa5, 0xfa, 0xa7, 0x2d, 0x95, 0x56, 0x54, 0x13, 0x57, 0x1f, 0xc7, 0xb3,
	0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x32, 0x52, 0xc4, 0x7b, 0x01, 0x00, 0x00,
}
