// Code generated by protoc-gen-go. DO NOT EDIT.
// source: issue_token.proto

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

type AUTH_METHOD int32

const (
	AUTH_METHOD_UNKNOWN                AUTH_METHOD = 0
	AUTH_METHOD_PASSWORD               AUTH_METHOD = 1
	AUTH_METHOD_TOKEN                  AUTH_METHOD = 2
	AUTH_METHOD_APPLICATION_CREDENTIAL AUTH_METHOD = 3
)

var AUTH_METHOD_name = map[int32]string{
	0: "UNKNOWN",
	1: "PASSWORD",
	2: "TOKEN",
	3: "APPLICATION_CREDENTIAL",
}
var AUTH_METHOD_value = map[string]int32{
	"UNKNOWN":                0,
	"PASSWORD":               1,
	"TOKEN":                  2,
	"APPLICATION_CREDENTIAL": 3,
}

func (x AUTH_METHOD) String() string {
	return proto.EnumName(AUTH_METHOD_name, int32(x))
}
func (AUTH_METHOD) EnumDescriptor() ([]byte, []int) { return fileDescriptor35, []int{0} }

type TokenScope struct {
	DomainId  *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	ProjectId *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
}

func (m *TokenScope) Reset()                    { *m = TokenScope{} }
func (m *TokenScope) String() string            { return proto.CompactTextString(m) }
func (*TokenScope) ProtoMessage()               {}
func (*TokenScope) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{0} }

func (m *TokenScope) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *TokenScope) GetProjectId() *google_protobuf.StringValue {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

type PasswordPayload struct {
	Id         *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username   *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password   *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	DomainId   *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
	DomainName *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=domain_name,json=domainName" json:"domain_name,omitempty"`
	Scope      *TokenScope                  `protobuf:"bytes,6,opt,name=scope" json:"scope,omitempty"`
}

func (m *PasswordPayload) Reset()                    { *m = PasswordPayload{} }
func (m *PasswordPayload) String() string            { return proto.CompactTextString(m) }
func (*PasswordPayload) ProtoMessage()               {}
func (*PasswordPayload) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{1} }

func (m *PasswordPayload) GetId() *google_protobuf.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PasswordPayload) GetUsername() *google_protobuf.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *PasswordPayload) GetPassword() *google_protobuf.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *PasswordPayload) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

func (m *PasswordPayload) GetDomainName() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainName
	}
	return nil
}

func (m *PasswordPayload) GetScope() *TokenScope {
	if m != nil {
		return m.Scope
	}
	return nil
}

type TokenPayload struct {
	TokenId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	Scope   *TokenScope                  `protobuf:"bytes,2,opt,name=scope" json:"scope,omitempty"`
}

func (m *TokenPayload) Reset()                    { *m = TokenPayload{} }
func (m *TokenPayload) String() string            { return proto.CompactTextString(m) }
func (*TokenPayload) ProtoMessage()               {}
func (*TokenPayload) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{2} }

func (m *TokenPayload) GetTokenId() *google_protobuf.StringValue {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *TokenPayload) GetScope() *TokenScope {
	if m != nil {
		return m.Scope
	}
	return nil
}

type ApplicationCredentialPayload struct {
	Id       *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Secret   *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
	UserId   *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Username *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=username" json:"username,omitempty"`
	DomainId *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=domain_id,json=domainId" json:"domain_id,omitempty"`
}

func (m *ApplicationCredentialPayload) Reset()                    { *m = ApplicationCredentialPayload{} }
func (m *ApplicationCredentialPayload) String() string            { return proto.CompactTextString(m) }
func (*ApplicationCredentialPayload) ProtoMessage()               {}
func (*ApplicationCredentialPayload) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{3} }

func (m *ApplicationCredentialPayload) GetId() *google_protobuf.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ApplicationCredentialPayload) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ApplicationCredentialPayload) GetSecret() *google_protobuf.StringValue {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *ApplicationCredentialPayload) GetUserId() *google_protobuf.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *ApplicationCredentialPayload) GetUsername() *google_protobuf.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *ApplicationCredentialPayload) GetDomainId() *google_protobuf.StringValue {
	if m != nil {
		return m.DomainId
	}
	return nil
}

type IssueTokenRequest struct {
	Method AUTH_METHOD `protobuf:"varint,1,opt,name=method,enum=ai.metathings.service.identity.AUTH_METHOD" json:"method,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*IssueTokenRequest_Password
	//	*IssueTokenRequest_Token
	//	*IssueTokenRequest_ApplicationCredential
	Payload isIssueTokenRequest_Payload `protobuf_oneof:"payload"`
}

func (m *IssueTokenRequest) Reset()                    { *m = IssueTokenRequest{} }
func (m *IssueTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*IssueTokenRequest) ProtoMessage()               {}
func (*IssueTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{4} }

type isIssueTokenRequest_Payload interface {
	isIssueTokenRequest_Payload()
}

type IssueTokenRequest_Password struct {
	Password *PasswordPayload `protobuf:"bytes,2,opt,name=password,oneof"`
}
type IssueTokenRequest_Token struct {
	Token *TokenPayload `protobuf:"bytes,3,opt,name=token,oneof"`
}
type IssueTokenRequest_ApplicationCredential struct {
	ApplicationCredential *ApplicationCredentialPayload `protobuf:"bytes,4,opt,name=application_credential,json=applicationCredential,oneof"`
}

func (*IssueTokenRequest_Password) isIssueTokenRequest_Payload()              {}
func (*IssueTokenRequest_Token) isIssueTokenRequest_Payload()                 {}
func (*IssueTokenRequest_ApplicationCredential) isIssueTokenRequest_Payload() {}

func (m *IssueTokenRequest) GetPayload() isIssueTokenRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *IssueTokenRequest) GetMethod() AUTH_METHOD {
	if m != nil {
		return m.Method
	}
	return AUTH_METHOD_UNKNOWN
}

func (m *IssueTokenRequest) GetPassword() *PasswordPayload {
	if x, ok := m.GetPayload().(*IssueTokenRequest_Password); ok {
		return x.Password
	}
	return nil
}

func (m *IssueTokenRequest) GetToken() *TokenPayload {
	if x, ok := m.GetPayload().(*IssueTokenRequest_Token); ok {
		return x.Token
	}
	return nil
}

func (m *IssueTokenRequest) GetApplicationCredential() *ApplicationCredentialPayload {
	if x, ok := m.GetPayload().(*IssueTokenRequest_ApplicationCredential); ok {
		return x.ApplicationCredential
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*IssueTokenRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _IssueTokenRequest_OneofMarshaler, _IssueTokenRequest_OneofUnmarshaler, _IssueTokenRequest_OneofSizer, []interface{}{
		(*IssueTokenRequest_Password)(nil),
		(*IssueTokenRequest_Token)(nil),
		(*IssueTokenRequest_ApplicationCredential)(nil),
	}
}

func _IssueTokenRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*IssueTokenRequest)
	// payload
	switch x := m.Payload.(type) {
	case *IssueTokenRequest_Password:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Password); err != nil {
			return err
		}
	case *IssueTokenRequest_Token:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Token); err != nil {
			return err
		}
	case *IssueTokenRequest_ApplicationCredential:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApplicationCredential); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("IssueTokenRequest.Payload has unexpected type %T", x)
	}
	return nil
}

func _IssueTokenRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*IssueTokenRequest)
	switch tag {
	case 2: // payload.password
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PasswordPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &IssueTokenRequest_Password{msg}
		return true, err
	case 3: // payload.token
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TokenPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &IssueTokenRequest_Token{msg}
		return true, err
	case 4: // payload.application_credential
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApplicationCredentialPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &IssueTokenRequest_ApplicationCredential{msg}
		return true, err
	default:
		return false, nil
	}
}

func _IssueTokenRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*IssueTokenRequest)
	// payload
	switch x := m.Payload.(type) {
	case *IssueTokenRequest_Password:
		s := proto.Size(x.Password)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *IssueTokenRequest_Token:
		s := proto.Size(x.Token)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *IssueTokenRequest_ApplicationCredential:
		s := proto.Size(x.ApplicationCredential)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type IssueTokenResponse struct {
	Token *Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *IssueTokenResponse) Reset()                    { *m = IssueTokenResponse{} }
func (m *IssueTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*IssueTokenResponse) ProtoMessage()               {}
func (*IssueTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor35, []int{5} }

func (m *IssueTokenResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenScope)(nil), "ai.metathings.service.identity.TokenScope")
	proto.RegisterType((*PasswordPayload)(nil), "ai.metathings.service.identity.PasswordPayload")
	proto.RegisterType((*TokenPayload)(nil), "ai.metathings.service.identity.TokenPayload")
	proto.RegisterType((*ApplicationCredentialPayload)(nil), "ai.metathings.service.identity.ApplicationCredentialPayload")
	proto.RegisterType((*IssueTokenRequest)(nil), "ai.metathings.service.identity.IssueTokenRequest")
	proto.RegisterType((*IssueTokenResponse)(nil), "ai.metathings.service.identity.IssueTokenResponse")
	proto.RegisterEnum("ai.metathings.service.identity.AUTH_METHOD", AUTH_METHOD_name, AUTH_METHOD_value)
}

func init() { proto.RegisterFile("issue_token.proto", fileDescriptor35) }

var fileDescriptor35 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x4f, 0xdb, 0x3e,
	0x14, 0xc6, 0x69, 0xa0, 0x69, 0x7b, 0x8a, 0xfe, 0xff, 0x62, 0x69, 0x08, 0x21, 0xc4, 0x50, 0xa5,
	0x49, 0x13, 0x83, 0x74, 0x62, 0x2f, 0xda, 0xc4, 0xa6, 0xad, 0xd0, 0x4a, 0x8d, 0x80, 0xa4, 0x4b,
	0xc3, 0xb8, 0xac, 0x4c, 0xe2, 0x05, 0x8f, 0x24, 0xce, 0x62, 0x87, 0x8a, 0xfb, 0x7d, 0x81, 0x7d,
	0xc2, 0x4d, 0xfb, 0x02, 0xbb, 0xdc, 0xed, 0x14, 0x27, 0x7d, 0x61, 0x9a, 0x20, 0xe5, 0x2e, 0x89,
	0x9f, 0xc7, 0x7e, 0xce, 0x39, 0x3f, 0x07, 0x56, 0x28, 0xe7, 0x09, 0x19, 0x0a, 0x76, 0x49, 0x42,
	0x2d, 0x8a, 0x99, 0x60, 0x68, 0x13, 0x53, 0x2d, 0x20, 0x02, 0x8b, 0x0b, 0x1a, 0x7a, 0x5c, 0xe3,
	0x24, 0xbe, 0xa2, 0x0e, 0xd1, 0xa8, 0x4b, 0x42, 0x41, 0xc5, 0xf5, 0xfa, 0xa6, 0xc7, 0x98, 0xe7,
	0x93, 0x96, 0x54, 0x9f, 0x27, 0x9f, 0x5a, 0xa3, 0x18, 0x47, 0x11, 0x89, 0x79, 0xe6, 0x5f, 0x7f,
	0xe9, 0x51, 0x71, 0x91, 0x9c, 0x6b, 0x0e, 0x0b, 0x5a, 0xc1, 0x88, 0x8a, 0x4b, 0x36, 0x6a, 0x79,
	0x6c, 0x57, 0x2e, 0xee, 0x5e, 0x61, 0x9f, 0xba, 0x58, 0xb0, 0x98, 0xb7, 0x26, 0x8f, 0xb9, 0xaf,
	0x3e, 0x13, 0xa2, 0xf9, 0xb5, 0x04, 0x60, 0xa7, 0xef, 0x03, 0x87, 0x45, 0x04, 0xbd, 0x86, 0x9a,
	0xcb, 0x02, 0x4c, 0xc3, 0x21, 0x75, 0xd7, 0x4a, 0x5b, 0xa5, 0xc7, 0xf5, 0xbd, 0x0d, 0x2d, 0xcb,
	0xa1, 0x8d, 0x73, 0x68, 0x03, 0x11, 0xd3, 0xd0, 0xfb, 0x88, 0xfd, 0x84, 0x58, 0xd5, 0x4c, 0xae,
	0xbb, 0x68, 0x1f, 0x20, 0x8a, 0xd9, 0x67, 0xe2, 0x88, 0xd4, 0xab, 0x14, 0xf0, 0xd6, 0x72, 0xbd,
	0xee, 0x36, 0x7f, 0x29, 0xf0, 0x7f, 0x1f, 0x73, 0x3e, 0x62, 0xb1, 0xdb, 0xc7, 0xd7, 0x3e, 0xc3,
	0x2e, 0xda, 0x01, 0xa5, 0x60, 0x08, 0x85, 0xba, 0xe8, 0x15, 0x54, 0x13, 0x4e, 0xe2, 0x10, 0x07,
	0xa4, 0xd0, 0xe1, 0x13, 0x75, 0xea, 0x8c, 0xf2, 0xa3, 0xd7, 0x16, 0x8b, 0x38, 0xc7, 0xea, 0x9b,
	0xdd, 0x5a, 0x9a, 0xab, 0x5b, 0x6f, 0xa1, 0x9e, 0x5b, 0x65, 0xe2, 0x72, 0x01, 0x33, 0x64, 0x06,
	0x23, 0xcd, 0xfc, 0x1e, 0xca, 0x3c, 0x1d, 0xd8, 0x9a, 0x2a, 0x8d, 0xdb, 0xda, 0xed, 0x2c, 0x69,
	0xd3, 0x11, 0x5b, 0x99, 0xb1, 0xf9, 0xad, 0x04, 0xcb, 0xf2, 0xeb, 0xb8, 0xdd, 0xef, 0xa0, 0x2a,
	0xc1, 0x28, 0x38, 0xf9, 0x03, 0xf5, 0xe7, 0xf7, 0x87, 0xca, 0x56, 0xc9, 0xaa, 0x48, 0x97, 0xee,
	0x4e, 0x33, 0x29, 0xf7, 0xcd, 0xf4, 0x43, 0x81, 0x8d, 0x76, 0x14, 0xf9, 0xd4, 0xc1, 0x82, 0xb2,
	0xf0, 0x30, 0x26, 0x52, 0x8b, 0xfd, 0xfb, 0x21, 0xf1, 0x14, 0x96, 0x0a, 0xe3, 0x20, 0x95, 0xe8,
	0x39, 0xa8, 0x9c, 0x38, 0x31, 0x11, 0x85, 0x40, 0xc8, 0xb5, 0xe8, 0x05, 0x54, 0x52, 0x98, 0x8a,
	0x42, 0xa0, 0xa6, 0x62, 0xfd, 0x26, 0xb1, 0xe5, 0xb9, 0x88, 0xbd, 0xc1, 0x9d, 0x3a, 0x0f, 0x77,
	0xcd, 0xdf, 0x0a, 0xac, 0xe8, 0xe9, 0xaf, 0x48, 0x76, 0xdf, 0x22, 0x5f, 0x12, 0xc2, 0x05, 0x3a,
	0x02, 0x35, 0x20, 0xe2, 0x82, 0x65, 0xbd, 0xfd, 0x6f, 0xef, 0xc9, 0x5d, 0xb3, 0x6b, 0x9f, 0xda,
	0xbd, 0xe1, 0x49, 0xd7, 0xee, 0x99, 0x9d, 0x09, 0x08, 0xf9, 0x16, 0xe8, 0x64, 0xe6, 0x3e, 0x65,
	0xad, 0x6f, 0xdd, 0xb5, 0xdd, 0x5f, 0x57, 0xbf, 0xb7, 0x30, 0x73, 0xc9, 0x3a, 0x50, 0x96, 0x84,
	0xe5, 0x23, 0xd9, 0x29, 0x84, 0xd5, 0x74, 0xa3, 0xcc, 0x8c, 0x12, 0x58, 0xc5, 0x53, 0xb2, 0x86,
	0xce, 0x04, 0xad, 0x7c, 0x64, 0x6f, 0xee, 0xac, 0xf8, 0x16, 0x2e, 0x7b, 0x0b, 0xd6, 0x03, 0xfc,
	0xaf, 0xf5, 0x83, 0x1a, 0x54, 0xa2, 0x4c, 0xd3, 0xfc, 0x00, 0x68, 0xb6, 0xf1, 0x3c, 0x62, 0x21,
	0x27, 0x68, 0x7f, 0x5c, 0x5d, 0x06, 0xf5, 0xa3, 0x42, 0xd5, 0xe5, 0x45, 0x6d, 0x9b, 0x50, 0x9f,
	0x19, 0x04, 0xaa, 0x43, 0xe5, 0xd4, 0x38, 0x32, 0xcc, 0x33, 0xa3, 0xb1, 0x80, 0x96, 0xa1, 0xda,
	0x6f, 0x0f, 0x06, 0x67, 0xa6, 0xd5, 0x69, 0x94, 0x50, 0x0d, 0xca, 0xb6, 0x79, 0xd4, 0x35, 0x1a,
	0x0a, 0x5a, 0x87, 0xd5, 0x76, 0xbf, 0x7f, 0xac, 0x1f, 0xb6, 0x6d, 0xdd, 0x34, 0x86, 0x87, 0x56,
	0xb7, 0xd3, 0x35, 0x6c, 0xbd, 0x7d, 0xdc, 0x58, 0x3c, 0x57, 0x25, 0x3d, 0xcf, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x9a, 0x33, 0x93, 0x04, 0xae, 0x06, 0x00, 0x00,
}