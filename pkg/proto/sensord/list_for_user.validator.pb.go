// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_for_user.proto

package sensord

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/nayotta/metathings/pkg/proto/sensor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListForUserRequest) Validate() error {
	if this.Name != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return go_proto_validators.FieldError("Name", err)
		}
	}
	if this.CoreId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.CoreId); err != nil {
			return go_proto_validators.FieldError("CoreId", err)
		}
	}
	if this.EntityName != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.EntityName); err != nil {
			return go_proto_validators.FieldError("EntityName", err)
		}
	}
	return nil
}
func (this *ListForUserResponse) Validate() error {
	for _, item := range this.Sensors {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Sensors", err)
			}
		}
	}
	return nil
}
