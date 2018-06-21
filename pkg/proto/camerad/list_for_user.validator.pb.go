// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: list_for_user.proto

package camerad

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/nayotta/metathings/pkg/proto/camera"

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
	if this.Core != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Core); err != nil {
			return go_proto_validators.FieldError("Core", err)
		}
	}
	if this.Entity != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Entity); err != nil {
			return go_proto_validators.FieldError("Entity", err)
		}
	}
	return nil
}
func (this *ListForUserResponse) Validate() error {
	for _, item := range this.Cameras {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Cameras", err)
			}
		}
	}
	return nil
}
