// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_token_by_text.proto

package identityd2

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetTokenByTextRequest) Validate() error {
	if nil == this.Text {
		return go_proto_validators.FieldError("Text", fmt.Errorf("message must exist"))
	}
	if this.Text != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Text); err != nil {
			return go_proto_validators.FieldError("Text", err)
		}
	}
	return nil
}
func (this *GetTokenByTextResponse) Validate() error {
	if this.Token != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Token); err != nil {
			return go_proto_validators.FieldError("Token", err)
		}
	}
	return nil
}