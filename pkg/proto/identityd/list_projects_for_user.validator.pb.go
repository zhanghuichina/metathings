// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_projects_for_user.proto

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

func (this *ListProjectsForUserRequest) Validate() error {
	if nil == this.UserId {
		return go_proto_validators.FieldError("UserId", fmt.Errorf("message must exist"))
	}
	if this.UserId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.UserId); err != nil {
			return go_proto_validators.FieldError("UserId", err)
		}
	}
	return nil
}
func (this *ListProjectsForUserResponse) Validate() error {
	for _, item := range this.Projects {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Projects", err)
			}
		}
	}
	return nil
}
