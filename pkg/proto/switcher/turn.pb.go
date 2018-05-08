// Code generated by protoc-gen-go. DO NOT EDIT.
// source: turn.proto

package switcher

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
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

type TurnRequest struct {
	State                SwitcherState `protobuf:"varint,1,opt,name=state,enum=ai.metathings.service.switcher.SwitcherState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TurnRequest) Reset()         { *m = TurnRequest{} }
func (m *TurnRequest) String() string { return proto.CompactTextString(m) }
func (*TurnRequest) ProtoMessage()    {}
func (*TurnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_turn_83a5c40d30fc9c74, []int{0}
}
func (m *TurnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TurnRequest.Unmarshal(m, b)
}
func (m *TurnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TurnRequest.Marshal(b, m, deterministic)
}
func (dst *TurnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TurnRequest.Merge(dst, src)
}
func (m *TurnRequest) XXX_Size() int {
	return xxx_messageInfo_TurnRequest.Size(m)
}
func (m *TurnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TurnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TurnRequest proto.InternalMessageInfo

func (m *TurnRequest) GetState() SwitcherState {
	if m != nil {
		return m.State
	}
	return SwitcherState_SWITCHER_STATE_UNKNOWN
}

type TurnResponse struct {
	Switcher             *Switcher `protobuf:"bytes,1,opt,name=switcher" json:"switcher,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TurnResponse) Reset()         { *m = TurnResponse{} }
func (m *TurnResponse) String() string { return proto.CompactTextString(m) }
func (*TurnResponse) ProtoMessage()    {}
func (*TurnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_turn_83a5c40d30fc9c74, []int{1}
}
func (m *TurnResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TurnResponse.Unmarshal(m, b)
}
func (m *TurnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TurnResponse.Marshal(b, m, deterministic)
}
func (dst *TurnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TurnResponse.Merge(dst, src)
}
func (m *TurnResponse) XXX_Size() int {
	return xxx_messageInfo_TurnResponse.Size(m)
}
func (m *TurnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TurnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TurnResponse proto.InternalMessageInfo

func (m *TurnResponse) GetSwitcher() *Switcher {
	if m != nil {
		return m.Switcher
	}
	return nil
}

func init() {
	proto.RegisterType((*TurnRequest)(nil), "ai.metathings.service.switcher.TurnRequest")
	proto.RegisterType((*TurnResponse)(nil), "ai.metathings.service.switcher.TurnResponse")
}

func init() { proto.RegisterFile("turn.proto", fileDescriptor_turn_83a5c40d30fc9c74) }

var fileDescriptor_turn_83a5c40d30fc9c74 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x24, 0x2a, 0xe4, 0xa2, 0x0e, 0x99, 0x50, 0x87, 0x52, 0x75, 0xea, 0x12, 0x5b,
	0x2a, 0x12, 0x0f, 0x80, 0xd8, 0xd8, 0xd2, 0x4e, 0x6c, 0x4e, 0x38, 0x1c, 0x8b, 0x26, 0x17, 0xee,
	0xce, 0xf5, 0xe3, 0x22, 0xf1, 0x24, 0x08, 0xbb, 0xcd, 0x88, 0xd8, 0xfe, 0xd3, 0xdd, 0xf7, 0xdd,
	0xaf, 0x94, 0x04, 0x1a, 0xf4, 0x48, 0x28, 0x58, 0xae, 0xac, 0xd7, 0x3d, 0x88, 0x95, 0xce, 0x0f,
	0x8e, 0x35, 0x03, 0x9d, 0x7c, 0x0b, 0x9a, 0xa3, 0x97, 0xb6, 0x03, 0x5a, 0xae, 0x1c, 0xa2, 0x3b,
	0x82, 0x49, 0xd7, 0x4d, 0x78, 0x37, 0x91, 0xec, 0x38, 0x02, 0x71, 0xe6, 0x97, 0x8f, 0xce, 0x4b,
	0x17, 0x1a, 0xdd, 0x62, 0x6f, 0xfa, 0xe8, 0xe5, 0x03, 0xa3, 0x71, 0x58, 0xa5, 0x65, 0x75, 0xb2,
	0x47, 0xff, 0x66, 0x05, 0x89, 0xcd, 0x14, 0xcf, 0xdc, 0xe2, 0xf2, 0x21, 0xcf, 0x9b, 0x57, 0x35,
	0x3f, 0x04, 0x1a, 0x6a, 0xf8, 0x0c, 0xc0, 0x52, 0xbe, 0xa8, 0x6b, 0x16, 0x2b, 0x70, 0x57, 0xac,
	0x8b, 0xed, 0x62, 0x57, 0xe9, 0xbf, 0x6b, 0xea, 0xfd, 0x39, 0xec, 0x7f, 0xa1, 0xa7, 0xd9, 0xf7,
	0xd7, 0xfd, 0xd5, 0xba, 0xa8, 0xb3, 0x63, 0x73, 0x50, 0xb7, 0xd9, 0xcd, 0x23, 0x0e, 0x0c, 0xe5,
	0xb3, 0xba, 0xb9, 0x80, 0xc9, 0x3f, 0xdf, 0x6d, 0xff, 0xeb, 0xaf, 0x27, 0xb2, 0x99, 0xa5, 0xe2,
	0x0f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x91, 0xb3, 0xcc, 0x4e, 0x01, 0x00, 0x00,
}
