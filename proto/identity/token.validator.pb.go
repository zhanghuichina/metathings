// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token.proto

package identity

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Token) Validate() error {
	if this.ExipresAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ExipresAt); err != nil {
			return go_proto_validators.FieldError("ExipresAt", err)
		}
	}
	if this.IssuedAt != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.IssuedAt); err != nil {
			return go_proto_validators.FieldError("IssuedAt", err)
		}
	}
	if this.User != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return go_proto_validators.FieldError("User", err)
		}
	}
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	if this.Project != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Project); err != nil {
			return go_proto_validators.FieldError("Project", err)
		}
	}
	return nil
}
func (this *Token__Domain) Validate() error {
	return nil
}
func (this *Token__Project) Validate() error {
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	return nil
}
func (this *Token__User) Validate() error {
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
		}
	}
	return nil
}
func (this *Token__Role) Validate() error {
	return nil
}
