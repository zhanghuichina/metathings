// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: check_role_for_user_on_project.proto

package identity

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

func (this *CheckRoleForUserOnProjectRequest) Validate() error {
	if nil == this.ProjectId {
		return go_proto_validators.FieldError("ProjectId", fmt.Errorf("message must exist"))
	}
	if this.ProjectId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.ProjectId); err != nil {
			return go_proto_validators.FieldError("ProjectId", err)
		}
	}
	if nil == this.UserId {
		return go_proto_validators.FieldError("UserId", fmt.Errorf("message must exist"))
	}
	if this.UserId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.UserId); err != nil {
			return go_proto_validators.FieldError("UserId", err)
		}
	}
	if nil == this.RoleId {
		return go_proto_validators.FieldError("RoleId", fmt.Errorf("message must exist"))
	}
	if this.RoleId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.RoleId); err != nil {
			return go_proto_validators.FieldError("RoleId", err)
		}
	}
	return nil
}
