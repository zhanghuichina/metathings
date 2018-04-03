// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_domain.proto

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

type DeleteDomainRequest struct {
	DomainId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
}

func (m *DeleteDomainRequest) Reset()                    { *m = DeleteDomainRequest{} }
func (m *DeleteDomainRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDomainRequest) ProtoMessage()               {}
func (*DeleteDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *DeleteDomainRequest) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteDomainRequest)(nil), "ai.metathings.service.identity.DeleteDomainRequest")
}

func init() { proto.RegisterFile("delete_domain.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xce, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x71, 0xd6, 0x43, 0xd1, 0xf5, 0xd6, 0x5e, 0xa4, 0x48, 0x2d, 0x9e, 0xbc, 0x74, 0x02,
	0x0a, 0xde, 0x95, 0x5e, 0xbc, 0x56, 0x10, 0x6f, 0x25, 0xbb, 0x19, 0xb3, 0x83, 0x49, 0x66, 0x4d,
	0x26, 0xbb, 0xf8, 0xb4, 0x82, 0x4f, 0x22, 0x24, 0xd8, 0xdb, 0xc0, 0x7c, 0x7f, 0xf8, 0xb5, 0x2b,
	0x83, 0x0e, 0x05, 0x8f, 0x86, 0xbd, 0xa6, 0x00, 0x63, 0x64, 0xe1, 0xe5, 0x46, 0x13, 0x78, 0x14,
	0x2d, 0x03, 0x05, 0x9b, 0x20, 0x61, 0x9c, 0xa8, 0x47, 0x20, 0x83, 0x41, 0x48, 0xbe, 0xd7, 0x1b,
	0xcb, 0x6c, 0x1d, 0xaa, 0xb2, 0xee, 0xf2, 0x87, 0x9a, 0xa3, 0x1e, 0x47, 0x8c, 0xa9, 0xf6, 0xeb,
	0x47, 0x4b, 0x32, 0xe4, 0x0e, 0x7a, 0xf6, 0xca, 0xcf, 0x24, 0x9f, 0x3c, 0x2b, 0xcb, 0xbb, 0xf2,
	0xdc, 0x4d, 0xda, 0x91, 0xd1, 0xc2, 0x31, 0xa9, 0xd3, 0x59, 0xbb, 0xdb, 0xf7, 0x76, 0xb5, 0x2f,
	0x9c, 0x7d, 0xd1, 0x1c, 0xf0, 0x2b, 0x63, 0x92, 0xe5, 0x53, 0x7b, 0x51, 0x79, 0x47, 0x32, 0x57,
	0xcd, 0xb6, 0xb9, 0xbb, 0xbc, 0xbf, 0x86, 0x4a, 0x80, 0x7f, 0x02, 0xbc, 0x4a, 0xa4, 0x60, 0xdf,
	0xb4, 0xcb, 0xf8, 0xbc, 0xf8, 0xfd, 0xb9, 0x39, 0xdb, 0x36, 0x87, 0xf3, 0x9a, 0xbd, 0x98, 0x6e,
	0x51, 0x76, 0x0f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xac, 0xfd, 0xe9, 0xe0, 0xef, 0x00, 0x00,
	0x00,
}
