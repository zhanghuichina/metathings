// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream_call.proto

package cored

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *StreamCallRequest) Validate() error {
	if this.CoreId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.CoreId); err != nil {
			return go_proto_validators.FieldError("CoreId", err)
		}
	}
	if this.Payload != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Payload); err != nil {
			return go_proto_validators.FieldError("Payload", err)
		}
	}
	return nil
}
func (this *StreamCallResponse) Validate() error {
	if this.Payload != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Payload); err != nil {
			return go_proto_validators.FieldError("Payload", err)
		}
	}
	return nil
}
