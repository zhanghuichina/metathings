// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: create.proto

package camerad

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateRequest) Validate() error {
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
	if this.Config != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
func (this *CreateResponse) Validate() error {
	if this.Camera != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Camera); err != nil {
			return go_proto_validators.FieldError("Camera", err)
		}
	}
	return nil
}