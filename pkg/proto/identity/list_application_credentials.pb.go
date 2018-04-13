// Code generated by protoc-gen-go. DO NOT EDIT.
// source: list_application_credentials.proto

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

type ListApplicationCredentialsRequest struct {
	UserId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *ListApplicationCredentialsRequest) Reset()         { *m = ListApplicationCredentialsRequest{} }
func (m *ListApplicationCredentialsRequest) String() string { return proto.CompactTextString(m) }
func (*ListApplicationCredentialsRequest) ProtoMessage()    {}
func (*ListApplicationCredentialsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor36, []int{0}
}

func (m *ListApplicationCredentialsRequest) GetUserId() *google_protobuf.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

type ListApplicationCredentialsResponse struct {
	ApplicationCredentials []*ApplicationCredential `protobuf:"bytes,1,rep,name=application_credentials,json=applicationCredentials" json:"application_credentials,omitempty"`
}

func (m *ListApplicationCredentialsResponse) Reset()         { *m = ListApplicationCredentialsResponse{} }
func (m *ListApplicationCredentialsResponse) String() string { return proto.CompactTextString(m) }
func (*ListApplicationCredentialsResponse) ProtoMessage()    {}
func (*ListApplicationCredentialsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor36, []int{1}
}

func (m *ListApplicationCredentialsResponse) GetApplicationCredentials() []*ApplicationCredential {
	if m != nil {
		return m.ApplicationCredentials
	}
	return nil
}

func init() {
	proto.RegisterType((*ListApplicationCredentialsRequest)(nil), "ai.metathings.service.identity.ListApplicationCredentialsRequest")
	proto.RegisterType((*ListApplicationCredentialsResponse)(nil), "ai.metathings.service.identity.ListApplicationCredentialsResponse")
}

func init() { proto.RegisterFile("list_application_credentials.proto", fileDescriptor36) }

var fileDescriptor36 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x09, 0x42, 0x85, 0xf4, 0xd6, 0x83, 0x96, 0x52, 0x4a, 0xcd, 0xa9, 0x97, 0x6e, 0xa0,
	0x52, 0xef, 0xe2, 0x49, 0xf0, 0x14, 0xc1, 0x83, 0x97, 0xb0, 0x49, 0xc6, 0xed, 0xe0, 0x66, 0x67,
	0xdd, 0x99, 0x6d, 0xf0, 0x7f, 0xf8, 0x83, 0xc5, 0xb4, 0xda, 0x4b, 0xec, 0x6d, 0x61, 0xbe, 0xb7,
	0xef, 0xe3, 0xa5, 0x99, 0x45, 0x96, 0x52, 0x7b, 0x6f, 0xb1, 0xd6, 0x82, 0xe4, 0xca, 0x3a, 0x40,
	0x03, 0x4e, 0x50, 0x5b, 0x56, 0x3e, 0x90, 0xd0, 0x64, 0xa1, 0x51, 0xb5, 0x20, 0x5a, 0x76, 0xe8,
	0x0c, 0x2b, 0x86, 0xb0, 0xc7, 0x1a, 0x14, 0xf6, 0x98, 0x7c, 0xce, 0x16, 0x86, 0xc8, 0x58, 0xc8,
	0x7b, 0xba, 0x8a, 0x6f, 0x79, 0x17, 0xb4, 0xf7, 0x10, 0x8e, 0xf9, 0xd9, 0x9d, 0x41, 0xd9, 0xc5,
	0x4a, 0xd5, 0xd4, 0xe6, 0x6d, 0x87, 0xf2, 0x4e, 0x5d, 0x6e, 0x68, 0xdd, 0x1f, 0xd7, 0x7b, 0x6d,
	0xb1, 0xd1, 0x42, 0x81, 0xf3, 0xbf, 0xe7, 0x31, 0x37, 0x1f, 0xd6, 0x3a, 0x5c, 0xb3, 0xd7, 0xf4,
	0xe6, 0x09, 0x59, 0xee, 0x4f, 0xcc, 0xc3, 0xc9, 0xbc, 0x80, 0x8f, 0x08, 0x2c, 0x93, 0x6d, 0x7a,
	0x19, 0x19, 0x42, 0x89, 0xcd, 0x34, 0x59, 0x26, 0xab, 0xf1, 0x66, 0xae, 0x0e, 0xb2, 0xea, 0x57,
	0x56, 0x3d, 0x4b, 0x40, 0x67, 0x5e, 0xb4, 0x8d, 0x50, 0x8c, 0x7e, 0xe0, 0xc7, 0x26, 0xfb, 0x4a,
	0xd2, 0xec, 0xdc, 0xe7, 0xec, 0xc9, 0x31, 0x4c, 0x5c, 0x7a, 0xfd, 0xcf, 0x72, 0xd3, 0x64, 0x79,
	0xb1, 0x1a, 0x6f, 0xb6, 0xea, 0xfc, 0x74, 0x6a, 0xb0, 0xa0, 0xb8, 0xd2, 0x83, 0xbd, 0xd5, 0xa8,
	0x97, 0xbe, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xff, 0x1d, 0x13, 0xb5, 0x01, 0x00, 0x00,
}