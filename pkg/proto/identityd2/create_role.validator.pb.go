// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create_role.proto

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

func (this *CreateRoleRequest) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if nil == this.Domain {
		return go_proto_validators.FieldError("Domain", fmt.Errorf("message must exist"))
	}
	if this.Domain != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Domain); err != nil {
			return go_proto_validators.FieldError("Domain", err)
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
	if nil == this.Alias {
		return go_proto_validators.FieldError("Alias", fmt.Errorf("message must exist"))
	}
	if this.Alias != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return go_proto_validators.FieldError("Alias", err)
		}
	}
	if this.Description != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Description); err != nil {
			return go_proto_validators.FieldError("Description", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CreateRoleResponse) Validate() error {
	if this.Role != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Role); err != nil {
			return go_proto_validators.FieldError("Role", err)
		}
	}
	return nil
}
