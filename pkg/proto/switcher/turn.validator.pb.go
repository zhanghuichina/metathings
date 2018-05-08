// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: turn.proto

package switcher

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

func (this *TurnRequest) Validate() error {
	return nil
}
func (this *TurnResponse) Validate() error {
	if this.Switcher != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Switcher); err != nil {
			return go_proto_validators.FieldError("Switcher", err)
		}
	}
	return nil
}
