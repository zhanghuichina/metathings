// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: get_region.proto

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

func (this *GetRegionRequest) Validate() error {
	if nil == this.RegionId {
		return go_proto_validators.FieldError("RegionId", fmt.Errorf("message must exist"))
	}
	if this.RegionId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.RegionId); err != nil {
			return go_proto_validators.FieldError("RegionId", err)
		}
	}
	return nil
}
func (this *GetRegionResponse) Validate() error {
	if this.Region != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Region); err != nil {
			return go_proto_validators.FieldError("Region", err)
		}
	}
	return nil
}
