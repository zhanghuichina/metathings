// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_role_in_group_on_domain.proto

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

func (this *ListRoleInGroupOnDomainRequest) Validate() error {
	if nil == this.DomainId {
		return go_proto_validators.FieldError("DomainId", fmt.Errorf("message must exist"))
	}
	if this.DomainId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.DomainId); err != nil {
			return go_proto_validators.FieldError("DomainId", err)
		}
	}
	if nil == this.GroupId {
		return go_proto_validators.FieldError("GroupId", fmt.Errorf("message must exist"))
	}
	if this.GroupId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.GroupId); err != nil {
			return go_proto_validators.FieldError("GroupId", err)
		}
	}
	return nil
}
func (this *ListRoleInGroupOnDomainResponse) Validate() error {
	for _, item := range this.Roles {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	return nil
}
