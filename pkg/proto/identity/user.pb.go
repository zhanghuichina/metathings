// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	Id               string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name             string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	DefaultProjectId string            `protobuf:"bytes,3,opt,name=default_project_id,json=defaultProjectId" json:"default_project_id,omitempty"`
	DomainId         string            `protobuf:"bytes,4,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	Enabled          bool              `protobuf:"varint,5,opt,name=enabled" json:"enabled,omitempty"`
	Extra            map[string]string `protobuf:"bytes,6,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor67, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDefaultProjectId() string {
	if m != nil {
		return m.DefaultProjectId
	}
	return ""
}

func (m *User) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *User) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *User) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "ai.metathings.service.identity.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor67) }

var fileDescriptor67 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x26, 0x69, 0x52, 0xdb, 0x11, 0xa4, 0x8c, 0x1e, 0x16, 0x05, 0x09, 0x9e, 0x72, 0x90, 0x15,
	0xf4, 0x52, 0xbc, 0xe7, 0xd0, 0x9b, 0x04, 0x3c, 0x87, 0x6d, 0x77, 0xd4, 0xd5, 0x64, 0x13, 0x76,
	0x27, 0xb5, 0x79, 0x20, 0xdf, 0x53, 0xb2, 0x69, 0xf1, 0xe6, 0xed, 0xfb, 0x85, 0xf9, 0x06, 0xa0,
	0xf7, 0xe4, 0x64, 0xe7, 0x5a, 0x6e, 0xf1, 0x56, 0x19, 0xd9, 0x10, 0x2b, 0xfe, 0x30, 0xf6, 0xdd,
	0x4b, 0x4f, 0x6e, 0x6f, 0x76, 0x24, 0x8d, 0x26, 0xcb, 0x86, 0x87, 0xbb, 0x9f, 0x18, 0x92, 0x57,
	0x4f, 0x0e, 0x2f, 0x20, 0x36, 0x5a, 0x44, 0x59, 0x94, 0x2f, 0xcb, 0xd8, 0x68, 0x44, 0x48, 0xac,
	0x6a, 0x48, 0xc4, 0x41, 0x09, 0x18, 0xef, 0x01, 0x35, 0xbd, 0xa9, 0xbe, 0xe6, 0xaa, 0x73, 0xed,
	0x27, 0xed, 0xb8, 0x32, 0x5a, 0xcc, 0x42, 0x62, 0x75, 0x74, 0x5e, 0x26, 0x63, 0xa3, 0xf1, 0x06,
	0x96, 0xba, 0x6d, 0x94, 0xb1, 0x63, 0x28, 0x09, 0xa1, 0xc5, 0x24, 0x6c, 0x34, 0x0a, 0x38, 0x23,
	0xab, 0xb6, 0x35, 0x69, 0x91, 0x66, 0x51, 0xbe, 0x28, 0x4f, 0x14, 0x0b, 0x48, 0xe9, 0xc0, 0x4e,
	0x89, 0x79, 0x36, 0xcb, 0xcf, 0x1f, 0x1f, 0xe4, 0xff, 0x0b, 0xe4, 0x78, 0xbd, 0x2c, 0xc6, 0x46,
	0x61, 0xd9, 0x0d, 0xe5, 0xd4, 0xbe, 0x5e, 0x03, 0xfc, 0x89, 0xb8, 0x82, 0xd9, 0x17, 0x0d, 0xc7,
	0x79, 0x23, 0xc4, 0x2b, 0x48, 0xf7, 0xaa, 0xee, 0x4f, 0x03, 0x27, 0xf2, 0x1c, 0xaf, 0xa3, 0xf2,
	0xb2, 0x53, 0xde, 0x7f, 0xb7, 0x4e, 0x57, 0x74, 0xe8, 0x8c, 0x23, 0x5f, 0x29, 0xde, 0xce, 0xc3,
	0x3b, 0x9f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xf1, 0xb9, 0x78, 0x5c, 0x01, 0x00, 0x00,
}