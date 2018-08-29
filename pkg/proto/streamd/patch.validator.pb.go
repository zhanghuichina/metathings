// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: patch.proto

package streamd

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

func (this *PatchRequest) Validate() error {
	if nil == this.Id {
		return go_proto_validators.FieldError("Id", fmt.Errorf("message must exist"))
	}
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
func (this *PatchResponse) Validate() error {
	if this.Stream != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Stream); err != nil {
			return go_proto_validators.FieldError("Stream", err)
		}
	}
	return nil
}
