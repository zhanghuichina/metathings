// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delete_user.proto

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

type DeleteUserRequest struct {
	UserId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *DeleteUserRequest) Reset()                    { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()               {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

func (m *DeleteUserRequest) GetUserId() *google_protobuf.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteUserRequest)(nil), "ai.metathings.service.identity.DeleteUserRequest")
}

func init() { proto.RegisterFile("delete_user.proto", fileDescriptor20) }

var fileDescriptor20 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x46, 0x89, 0x45, 0x84, 0x58, 0xdd, 0x55, 0x72, 0xc8, 0x79, 0x58, 0xd9, 0xdc, 0x2c, 0x28,
	0xd8, 0xd9, 0x88, 0x8d, 0x6d, 0x44, 0xdb, 0x63, 0x73, 0x3b, 0x6e, 0x06, 0x37, 0x99, 0x38, 0x3b,
	0x9b, 0xe0, 0xaf, 0x15, 0xfc, 0x25, 0x92, 0x0d, 0x5e, 0x37, 0x30, 0xdf, 0xe3, 0xbd, 0x6a, 0xe5,
	0x30, 0xa0, 0xe2, 0x21, 0x45, 0x14, 0x18, 0x84, 0x95, 0xd7, 0x5b, 0x4b, 0xd0, 0xa1, 0x5a, 0x6d,
	0xa9, 0xf7, 0x11, 0x22, 0xca, 0x48, 0x47, 0x04, 0x72, 0xd8, 0x2b, 0xe9, 0xf7, 0x66, 0xeb, 0x99,
	0x7d, 0x40, 0x93, 0xd7, 0x4d, 0xfa, 0x30, 0x93, 0xd8, 0x61, 0x40, 0x89, 0x0b, 0xbf, 0x79, 0xf0,
	0xa4, 0x6d, 0x6a, 0xe0, 0xc8, 0x9d, 0xe9, 0x26, 0xd2, 0x4f, 0x9e, 0x8c, 0xe7, 0x7d, 0x7e, 0xee,
	0x47, 0x1b, 0xc8, 0x59, 0x65, 0x89, 0xe6, 0x74, 0x2e, 0xdc, 0x4d, 0x5d, 0xad, 0x9e, 0x73, 0xcc,
	0x5b, 0x44, 0xa9, 0xf1, 0x2b, 0x61, 0xd4, 0xf5, 0x63, 0x75, 0x3e, 0xa7, 0x1d, 0xc8, 0x5d, 0x16,
	0xbb, 0xe2, 0xf6, 0xe2, 0xee, 0x0a, 0x16, 0x3d, 0xfc, 0xeb, 0xe1, 0x55, 0x85, 0x7a, 0xff, 0x6e,
	0x43, 0xc2, 0xa7, 0xf2, 0xf7, 0xe7, 0xfa, 0x6c, 0x57, 0xd4, 0xe5, 0x0c, 0xbd, 0xb8, 0xa6, 0xcc,
	0xab, 0xfb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x23, 0x98, 0x80, 0x78, 0xe7, 0x00, 0x00, 0x00,
}
