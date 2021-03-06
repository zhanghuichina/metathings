// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_application_credential.proto

package identityd

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateApplicationCredentialRequest) Validate() error {
	if nil == this.UserId {
		return go_proto_validators.FieldError("UserId", fmt.Errorf("message must exist"))
	}
	if this.UserId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.UserId); err != nil {
			return go_proto_validators.FieldError("UserId", err)
		}
	}
	if nil == this.Name {
		return go_proto_validators.FieldError("Name", fmt.Errorf("message must exist"))
	}
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	if this.Description != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return go_proto_validators.FieldError("Description", err)
		}
	}
	if this.Secret != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Secret); err != nil {
			return go_proto_validators.FieldError("Secret", err)
		}
	}
	if this.Unrestricted != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Unrestricted); err != nil {
			return go_proto_validators.FieldError("Unrestricted", err)
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	if this.ExpiresAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ExpiresAt); err != nil {
			return go_proto_validators.FieldError("ExpiresAt", err)
		}
	}
	return nil
}
func (this *CreateApplicationCredentialRequest__Role) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	return nil
}
func (this *CreateApplicationCredentialResponse) Validate() error {
	if this.ApplicationCredential != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ApplicationCredential); err != nil {
			return go_proto_validators.FieldError("ApplicationCredential", err)
		}
	}
	return nil
}
