// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token.proto

package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Token struct {
	Methods               []string                      `protobuf:"bytes,1,rep,name=methods" json:"methods,omitempty"`
	ExipresAt             *timestamp.Timestamp          `protobuf:"bytes,2,opt,name=exipres_at,json=exipresAt" json:"exipres_at,omitempty"`
	IssuedAt              *timestamp.Timestamp          `protobuf:"bytes,3,opt,name=issued_at,json=issuedAt" json:"issued_at,omitempty"`
	IsDomain              bool                          `protobuf:"varint,4,opt,name=is_domain,json=isDomain" json:"is_domain,omitempty"`
	User                  *Token__User                  `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	Roles                 []*Token__Role                `protobuf:"bytes,6,rep,name=roles" json:"roles,omitempty"`
	Project               *Token__Project               `protobuf:"bytes,7,opt,name=project" json:"project,omitempty"`
	ApplicationCredential *Token__ApplicationCredential `protobuf:"bytes,8,opt,name=application_credential,json=applicationCredential" json:"application_credential,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                      `json:"-"`
	XXX_unrecognized      []byte                        `json:"-"`
	XXX_sizecache         int32                         `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_eac893d68c44a9c7, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *Token) GetExipresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExipresAt
	}
	return nil
}

func (m *Token) GetIssuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.IssuedAt
	}
	return nil
}

func (m *Token) GetIsDomain() bool {
	if m != nil {
		return m.IsDomain
	}
	return false
}

func (m *Token) GetUser() *Token__User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Token) GetRoles() []*Token__Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Token) GetProject() *Token__Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *Token) GetApplicationCredential() *Token__ApplicationCredential {
	if m != nil {
		return m.ApplicationCredential
	}
	return nil
}

type Token__Domain struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token__Domain) Reset()         { *m = Token__Domain{} }
func (m *Token__Domain) String() string { return proto.CompactTextString(m) }
func (*Token__Domain) ProtoMessage()    {}
func (*Token__Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_eac893d68c44a9c7, []int{0, 0}
}
func (m *Token__Domain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token__Domain.Unmarshal(m, b)
}
func (m *Token__Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token__Domain.Marshal(b, m, deterministic)
}
func (dst *Token__Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token__Domain.Merge(dst, src)
}
func (m *Token__Domain) XXX_Size() int {
	return xxx_messageInfo_Token__Domain.Size(m)
}
func (m *Token__Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Token__Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Token__Domain proto.InternalMessageInfo

func (m *Token__Domain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token__Domain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Token__Project struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Domain               *Token__Domain `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Token__Project) Reset()         { *m = Token__Project{} }
func (m *Token__Project) String() string { return proto.CompactTextString(m) }
func (*Token__Project) ProtoMessage()    {}
func (*Token__Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_eac893d68c44a9c7, []int{0, 1}
}
func (m *Token__Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token__Project.Unmarshal(m, b)
}
func (m *Token__Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token__Project.Marshal(b, m, deterministic)
}
func (dst *Token__Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token__Project.Merge(dst, src)
}
func (m *Token__Project) XXX_Size() int {
	return xxx_messageInfo_Token__Project.Size(m)
}
func (m *Token__Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Token__Project.DiscardUnknown(m)
}

var xxx_messageInfo_Token__Project proto.InternalMessageInfo

func (m *Token__Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token__Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Token__Project) GetDomain() *Token__Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

type Token__User struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Domain               *Token__Domain `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Token__User) Reset()         { *m = Token__User{} }
func (m *Token__User) String() string { return proto.CompactTextString(m) }
func (*Token__User) ProtoMessage()    {}
func (*Token__User) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_eac893d68c44a9c7, []int{0, 2}
}
func (m *Token__User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token__User.Unmarshal(m, b)
}
func (m *Token__User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token__User.Marshal(b, m, deterministic)
}
func (dst *Token__User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token__User.Merge(dst, src)
}
func (m *Token__User) XXX_Size() int {
	return xxx_messageInfo_Token__User.Size(m)
}
func (m *Token__User) XXX_DiscardUnknown() {
	xxx_messageInfo_Token__User.DiscardUnknown(m)
}

var xxx_messageInfo_Token__User proto.InternalMessageInfo

func (m *Token__User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token__User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Token__User) GetDomain() *Token__Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

type Token__Role struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token__Role) Reset()         { *m = Token__Role{} }
func (m *Token__Role) String() string { return proto.CompactTextString(m) }
func (*Token__Role) ProtoMessage()    {}
func (*Token__Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_eac893d68c44a9c7, []int{0, 3}
}
func (m *Token__Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token__Role.Unmarshal(m, b)
}
func (m *Token__Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token__Role.Marshal(b, m, deterministic)
}
func (dst *Token__Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token__Role.Merge(dst, src)
}
func (m *Token__Role) XXX_Size() int {
	return xxx_messageInfo_Token__Role.Size(m)
}
func (m *Token__Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Token__Role.DiscardUnknown(m)
}

var xxx_messageInfo_Token__Role proto.InternalMessageInfo

func (m *Token__Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token__Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Token__ApplicationCredential struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Restricted           bool     `protobuf:"varint,3,opt,name=restricted" json:"restricted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token__ApplicationCredential) Reset()         { *m = Token__ApplicationCredential{} }
func (m *Token__ApplicationCredential) String() string { return proto.CompactTextString(m) }
func (*Token__ApplicationCredential) ProtoMessage()    {}
func (*Token__ApplicationCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_token_eac893d68c44a9c7, []int{0, 4}
}
func (m *Token__ApplicationCredential) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token__ApplicationCredential.Unmarshal(m, b)
}
func (m *Token__ApplicationCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token__ApplicationCredential.Marshal(b, m, deterministic)
}
func (dst *Token__ApplicationCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token__ApplicationCredential.Merge(dst, src)
}
func (m *Token__ApplicationCredential) XXX_Size() int {
	return xxx_messageInfo_Token__ApplicationCredential.Size(m)
}
func (m *Token__ApplicationCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_Token__ApplicationCredential.DiscardUnknown(m)
}

var xxx_messageInfo_Token__ApplicationCredential proto.InternalMessageInfo

func (m *Token__ApplicationCredential) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token__ApplicationCredential) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Token__ApplicationCredential) GetRestricted() bool {
	if m != nil {
		return m.Restricted
	}
	return false
}

func init() {
	proto.RegisterType((*Token)(nil), "ai.metathings.service.identity.Token")
	proto.RegisterType((*Token__Domain)(nil), "ai.metathings.service.identity.Token._Domain")
	proto.RegisterType((*Token__Project)(nil), "ai.metathings.service.identity.Token._Project")
	proto.RegisterType((*Token__User)(nil), "ai.metathings.service.identity.Token._User")
	proto.RegisterType((*Token__Role)(nil), "ai.metathings.service.identity.Token._Role")
	proto.RegisterType((*Token__ApplicationCredential)(nil), "ai.metathings.service.identity.Token._ApplicationCredential")
}

func init() { proto.RegisterFile("token.proto", fileDescriptor_token_eac893d68c44a9c7) }

var fileDescriptor_token_eac893d68c44a9c7 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0x86, 0xc9, 0xcc, 0x64, 0x92, 0xd4, 0x82, 0x87, 0x16, 0x97, 0x10, 0x61, 0x0d, 0x9e, 0x06,
	0x96, 0xed, 0x05, 0x3d, 0x88, 0x20, 0xc8, 0xa0, 0x82, 0x47, 0x69, 0xd6, 0x9b, 0x10, 0x7a, 0xd3,
	0xe5, 0x6c, 0x69, 0x92, 0x0e, 0xdd, 0x15, 0x5d, 0x9f, 0xc0, 0x17, 0xf3, 0xc1, 0x64, 0x3a, 0x13,
	0xf5, 0x30, 0x68, 0xbc, 0xec, 0x6d, 0xba, 0xa7, 0xbe, 0x2f, 0x55, 0xf5, 0x27, 0x70, 0xc2, 0xf6,
	0x33, 0x76, 0xb2, 0x77, 0x96, 0xad, 0x38, 0xd3, 0x24, 0x5b, 0x64, 0xcd, 0x37, 0xd4, 0xed, 0xbc,
	0xf4, 0xe8, 0xbe, 0x50, 0x8d, 0x92, 0x0c, 0x76, 0x4c, 0xfc, 0xad, 0x78, 0xb4, 0xb3, 0x76, 0xd7,
	0xe0, 0x65, 0xa8, 0xbe, 0x1e, 0x3e, 0x5e, 0x32, 0xb5, 0xe8, 0x59, 0xb7, 0xfd, 0x28, 0x78, 0xfc,
	0x23, 0x81, 0xf8, 0x6a, 0x2f, 0x14, 0x39, 0x24, 0x2d, 0xf2, 0x8d, 0x35, 0x3e, 0x8f, 0xca, 0xe5,
	0x26, 0x53, 0xd3, 0x51, 0x3c, 0x07, 0xc0, 0x5b, 0xea, 0x1d, 0xfa, 0x4a, 0x73, 0xbe, 0x28, 0xa3,
	0xcd, 0xc9, 0x93, 0x42, 0x8e, 0x66, 0x39, 0x99, 0xe5, 0xd5, 0x64, 0x56, 0xd9, 0xa1, 0x7a, 0xcb,
	0xe2, 0x19, 0x64, 0xe4, 0xfd, 0x80, 0x66, 0x4f, 0x2e, 0xff, 0x49, 0xa6, 0x63, 0xf1, 0x96, 0xc5,
	0xc3, 0x3d, 0x58, 0x19, 0xdb, 0x6a, 0xea, 0xf2, 0x55, 0x19, 0x6d, 0xd2, 0xfd, 0x9f, 0xaf, 0xc3,
	0x59, 0xbc, 0x84, 0xd5, 0xe0, 0xd1, 0xe5, 0x71, 0x10, 0x9e, 0xcb, 0xbf, 0x2f, 0x41, 0x86, 0xf9,
	0x64, 0xf5, 0xde, 0xa3, 0x53, 0x01, 0x14, 0x5b, 0x88, 0x9d, 0x6d, 0xd0, 0xe7, 0xeb, 0x72, 0xf9,
	0x1f, 0x06, 0x65, 0x1b, 0x54, 0x23, 0x29, 0xde, 0x42, 0xd2, 0x3b, 0xfb, 0x09, 0x6b, 0xce, 0x93,
	0xd0, 0x86, 0x9c, 0x29, 0x79, 0x37, 0x52, 0x6a, 0xc2, 0x85, 0x87, 0x53, 0xdd, 0xf7, 0x0d, 0xd5,
	0x9a, 0xc9, 0x76, 0x55, 0xed, 0x30, 0x10, 0xba, 0xc9, 0xd3, 0x20, 0x7e, 0x31, 0x53, 0xbc, 0xfd,
	0x2d, 0x79, 0xf5, 0xcb, 0xa1, 0x1e, 0xe8, 0x63, 0xd7, 0xc5, 0x05, 0x24, 0xd5, 0x61, 0x9b, 0xf7,
	0x60, 0x41, 0x26, 0x8f, 0xca, 0x68, 0x93, 0xa9, 0x05, 0x19, 0x21, 0x60, 0xd5, 0xe9, 0x16, 0x43,
	0xd0, 0x99, 0x0a, 0xbf, 0x8b, 0x01, 0xd2, 0xa9, 0xf1, 0x39, 0xf5, 0xe2, 0x0d, 0xac, 0x0f, 0xd9,
	0x8d, 0xa1, 0x5f, 0xcc, 0x9c, 0x61, 0x6c, 0x49, 0x1d, 0xe0, 0xe2, 0x7b, 0x04, 0x71, 0xc8, 0xed,
	0x0e, 0x1f, 0xaa, 0xee, 0xf7, 0xda, 0xfb, 0xaf, 0xd6, 0x99, 0x0a, 0x6f, 0x7b, 0x1a, 0xdf, 0xfa,
	0xe2, 0x1c, 0xe2, 0x10, 0xff, 0xac, 0x6d, 0x7d, 0x80, 0xd3, 0xe3, 0x69, 0xcc, 0x1a, 0xe3, 0x0c,
	0xc0, 0xa1, 0x67, 0x47, 0x35, 0xa3, 0x09, 0xa3, 0xa4, 0xea, 0x8f, 0x1b, 0x95, 0xd4, 0x9a, 0x75,
	0x63, 0x77, 0x2a, 0xd3, 0x83, 0x21, 0xae, 0xc8, 0xf8, 0xeb, 0x75, 0xf8, 0x98, 0x9e, 0xfe, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x0c, 0xb1, 0x6b, 0xe7, 0x1d, 0x04, 0x00, 0x00,
}
