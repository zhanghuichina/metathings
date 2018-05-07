// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_application_credential.proto

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

type DeleteApplicationCredentialRequest struct {
	UserId                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	ApplicationCredentialId *wrappers.StringValue `protobuf:"bytes,2,opt,name=application_credential_id,json=applicationCredentialId" json:"application_credential_id,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *DeleteApplicationCredentialRequest) Reset()         { *m = DeleteApplicationCredentialRequest{} }
func (m *DeleteApplicationCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteApplicationCredentialRequest) ProtoMessage()    {}
func (*DeleteApplicationCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_delete_application_credential_5990b7f1e4b37882, []int{0}
}
func (m *DeleteApplicationCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteApplicationCredentialRequest.Unmarshal(m, b)
}
func (m *DeleteApplicationCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteApplicationCredentialRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteApplicationCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteApplicationCredentialRequest.Merge(dst, src)
}
func (m *DeleteApplicationCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteApplicationCredentialRequest.Size(m)
}
func (m *DeleteApplicationCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteApplicationCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteApplicationCredentialRequest proto.InternalMessageInfo

func (m *DeleteApplicationCredentialRequest) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *DeleteApplicationCredentialRequest) GetApplicationCredentialId() *wrappers.StringValue {
	if m != nil {
		return m.ApplicationCredentialId
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteApplicationCredentialRequest)(nil), "ai.metathings.service.identity.DeleteApplicationCredentialRequest")
}

func init() {
	proto.RegisterFile("delete_application_credential.proto", fileDescriptor_delete_application_credential_5990b7f1e4b37882)
}

var fileDescriptor_delete_application_credential_5990b7f1e4b37882 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0x0e, 0x41, 0x0a, 0x5b, 0x16, 0xa0, 0x42, 0xa5, 0x2a, 0x0b, 0x4b, 0x6d, 0x09,
	0x24, 0x36, 0x06, 0xfe, 0x2c, 0x5d, 0x8b, 0xc4, 0x1a, 0x2e, 0xf1, 0xe1, 0x9e, 0x70, 0x62, 0x63,
	0x9f, 0x1b, 0xf1, 0x80, 0x3c, 0x07, 0x12, 0x4f, 0x82, 0xea, 0x88, 0xb2, 0x64, 0xe9, 0x76, 0xd2,
	0x7d, 0xbf, 0x4f, 0xbf, 0xaf, 0xb8, 0x54, 0x68, 0x90, 0xb1, 0x02, 0xe7, 0x0c, 0x35, 0xc0, 0x64,
	0xbb, 0xaa, 0xf1, 0xa8, 0xb0, 0x63, 0x02, 0x23, 0x9c, 0xb7, 0x6c, 0xcb, 0x19, 0x90, 0x68, 0x91,
	0x81, 0x37, 0xd4, 0xe9, 0x20, 0x02, 0xfa, 0x2d, 0x35, 0x28, 0x28, 0xa5, 0xf8, 0x73, 0x3a, 0xd3,
	0xd6, 0x6a, 0x83, 0x32, 0xa5, 0xeb, 0xf8, 0x26, 0x7b, 0x0f, 0xce, 0xa1, 0x0f, 0x03, 0x3f, 0xbd,
	0xd5, 0xc4, 0x9b, 0x58, 0x8b, 0xc6, 0xb6, 0xb2, 0xed, 0x89, 0xdf, 0x6d, 0x2f, 0xb5, 0x5d, 0xa6,
	0xe7, 0x72, 0x0b, 0x86, 0x14, 0xb0, 0xf5, 0x41, 0xee, 0xcf, 0x81, 0x5b, 0x7c, 0x65, 0xc5, 0xe2,
	0x29, 0xf9, 0xdd, 0xff, 0xeb, 0x3d, 0xee, 0xed, 0xd6, 0xf8, 0x11, 0x31, 0x70, 0x79, 0x57, 0x1c,
	0xc5, 0x80, 0xbe, 0x22, 0x75, 0x9a, 0xcd, 0xb3, 0xab, 0xe3, 0xeb, 0x73, 0x31, 0x08, 0x89, 0x3f,
	0x21, 0xf1, 0xcc, 0x9e, 0x3a, 0xfd, 0x02, 0x26, 0xe2, 0x43, 0xfe, 0xf3, 0x7d, 0x31, 0x99, 0x67,
	0xeb, 0x7c, 0x07, 0xad, 0x54, 0xf9, 0x5a, 0x9c, 0x8d, 0xaf, 0xdf, 0x15, 0x4e, 0x0e, 0x28, 0x3c,
	0x81, 0x31, 0xcb, 0x95, 0xaa, 0xf3, 0x84, 0xdd, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x2d,
	0x41, 0xea, 0x6d, 0x01, 0x00, 0x00,
}
