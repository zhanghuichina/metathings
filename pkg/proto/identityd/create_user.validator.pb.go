// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_user.proto

package identityd

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

func (this *CreateUserRequest) Validate() error {
	if nil == this.DomainId {
		return go_proto_validators.FieldError("DomainId", fmt.Errorf("message must exist"))
	}
	if this.DomainId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.DomainId); err != nil {
			return go_proto_validators.FieldError("DomainId", err)
		}
	}
	if this.DefaultProjectId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.DefaultProjectId); err != nil {
			return go_proto_validators.FieldError("DefaultProjectId", err)
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
	if nil == this.Password {
		return go_proto_validators.FieldError("Password", fmt.Errorf("message must exist"))
	}
	if this.Password != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Password); err != nil {
			return go_proto_validators.FieldError("Password", err)
		}
	}
	if this.Enabled != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Enabled); err != nil {
			return go_proto_validators.FieldError("Enabled", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CreateUserResponse) Validate() error {
	if this.User != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
