// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: show_by_application_credential_id.proto

package camerad

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

func (this *ShowByApplicationCredentialIdRequest) Validate() error {
	if nil == this.ApplicationCredentialId {
		return go_proto_validators.FieldError("ApplicationCredentialId", fmt.Errorf("message must exist"))
	}
	if this.ApplicationCredentialId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ApplicationCredentialId); err != nil {
			return go_proto_validators.FieldError("ApplicationCredentialId", err)
		}
	}
	return nil
}
func (this *ShowByApplicationCredentialIdResponse) Validate() error {
	if this.Camera != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Camera); err != nil {
			return go_proto_validators.FieldError("Camera", err)
		}
	}
	return nil
}
