// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patch_core.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import ai_metathings_state "github.com/bigdatagz/metathings/pkg/proto/common/state"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PatchCoreRequest struct {
	Id    *google_protobuf.StringValue  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  *google_protobuf.StringValue  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	State ai_metathings_state.CoreState `protobuf:"varint,3,opt,name=state,enum=ai.metathings.state.CoreState" json:"state,omitempty"`
}

func (m *PatchCoreRequest) Reset()                    { *m = PatchCoreRequest{} }
func (m *PatchCoreRequest) String() string            { return proto.CompactTextString(m) }
func (*PatchCoreRequest) ProtoMessage()               {}
func (*PatchCoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *PatchCoreRequest) GetId() *google_protobuf.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PatchCoreRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PatchCoreRequest) GetState() ai_metathings_state.CoreState {
	if m != nil {
		return m.State
	}
	return ai_metathings_state.CoreState_CORE_STATE_UNKNOWN
}

type PatchCoreResponse struct {
	Core *Core `protobuf:"bytes,1,opt,name=core" json:"core,omitempty"`
}

func (m *PatchCoreResponse) Reset()                    { *m = PatchCoreResponse{} }
func (m *PatchCoreResponse) String() string            { return proto.CompactTextString(m) }
func (*PatchCoreResponse) ProtoMessage()               {}
func (*PatchCoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *PatchCoreResponse) GetCore() *Core {
	if m != nil {
		return m.Core
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchCoreRequest)(nil), "ai.metathings.service.core.PatchCoreRequest")
	proto.RegisterType((*PatchCoreResponse)(nil), "ai.metathings.service.core.PatchCoreResponse")
}

func init() { proto.RegisterFile("patch_core.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x69, 0x9d, 0x3b, 0x44, 0x90, 0xd9, 0xd3, 0x18, 0x32, 0xcb, 0x4e, 0xbb, 0x2c, 0x91,
	0x39, 0xfc, 0x00, 0x7a, 0x10, 0x6f, 0xd2, 0x81, 0xd7, 0x91, 0xb6, 0x7f, 0xb3, 0xb0, 0xb5, 0xff,
	0x98, 0xfc, 0xbb, 0x81, 0x9f, 0xca, 0x6f, 0x24, 0xf8, 0x49, 0x24, 0x69, 0x75, 0x43, 0x10, 0x6f,
	0x0f, 0xf2, 0xde, 0xcb, 0xef, 0x25, 0x6c, 0x60, 0x24, 0x15, 0xeb, 0x55, 0x81, 0x16, 0xb8, 0xb1,
	0x48, 0x98, 0x8c, 0xa4, 0xe6, 0x15, 0x90, 0xa4, 0xb5, 0xae, 0x95, 0xe3, 0x0e, 0xec, 0x4e, 0x17,
	0xc0, 0xbd, 0x63, 0x34, 0x56, 0x88, 0x6a, 0x0b, 0x22, 0x38, 0xf3, 0xe6, 0x45, 0xec, 0xad, 0x34,
	0x06, 0xac, 0x6b, 0xb3, 0xa3, 0x5b, 0xa5, 0x69, 0xdd, 0xe4, 0xbc, 0xc0, 0x4a, 0x54, 0x7b, 0x4d,
	0x1b, 0xdc, 0x0b, 0x85, 0xb3, 0x70, 0x38, 0xdb, 0xc9, 0xad, 0x2e, 0x25, 0xa1, 0x75, 0xe2, 0x47,
	0x76, 0xb9, 0x87, 0xa3, 0x5c, 0xae, 0x55, 0x29, 0x49, 0xaa, 0x37, 0x71, 0xa0, 0x10, 0x66, 0xa3,
	0xda, 0x4b, 0x45, 0x81, 0x55, 0x85, 0xb5, 0x70, 0x24, 0x09, 0x84, 0x87, 0x5a, 0x05, 0xd9, 0x15,
	0xb1, 0xc3, 0x90, 0xc9, 0x7b, 0xc4, 0x06, 0x4f, 0x7e, 0xdd, 0x3d, 0x5a, 0xc8, 0xe0, 0xb5, 0x01,
	0x47, 0xc9, 0x82, 0xc5, 0xba, 0x1c, 0x46, 0x69, 0x34, 0x3d, 0x9b, 0x5f, 0xf2, 0x76, 0x0e, 0xff,
	0x9e, 0xc3, 0x97, 0x64, 0x75, 0xad, 0x9e, 0xe5, 0xb6, 0x81, 0xbb, 0xfe, 0xe7, 0xc7, 0x55, 0x9c,
	0x46, 0x59, 0xac, 0xcb, 0xe4, 0x9a, 0xf5, 0x6a, 0x59, 0xc1, 0x30, 0xfe, 0x3f, 0x97, 0x05, 0x67,
	0xb2, 0x60, 0xa7, 0x81, 0x6b, 0x78, 0x92, 0x46, 0xd3, 0xf3, 0xf9, 0x98, 0xff, 0x7a, 0xd5, 0xc0,
	0xec, 0xc1, 0x96, 0x5e, 0x65, 0xad, 0x79, 0xf2, 0xc8, 0x2e, 0x8e, 0x88, 0x9d, 0xc1, 0xda, 0xf9,
	0xaa, 0x9e, 0x5f, 0xd5, 0x41, 0xa7, 0xfc, 0xef, 0xff, 0x09, 0x85, 0x59, 0x70, 0xe7, 0xfd, 0x00,
	0x77, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0x69, 0xe0, 0xde, 0xc0, 0xe1, 0x01, 0x00, 0x00,
}
