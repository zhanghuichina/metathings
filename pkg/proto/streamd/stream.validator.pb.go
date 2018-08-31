// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream.proto

package streamd

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Stream) Validate() error {
	for _, item := range this.Sources {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Sources", err)
			}
		}
	}
	for _, item := range this.Groups {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Groups", err)
			}
		}
	}
	return nil
}
func (this *Upstream) Validate() error {
	return nil
}
func (this *Input) Validate() error {
	return nil
}
func (this *Output) Validate() error {
	return nil
}
func (this *Source) Validate() error {
	if this.Upstream != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Upstream); err != nil {
			return go_proto_validators.FieldError("Upstream", err)
		}
	}
	return nil
}
func (this *Group) Validate() error {
	for _, item := range this.Inputs {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Inputs", err)
			}
		}
	}
	for _, item := range this.Outputs {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Outputs", err)
			}
		}
	}
	return nil
}
func (this *OpUpstream) Validate() error {
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
	if this.Alias != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return go_proto_validators.FieldError("Alias", err)
		}
	}
	for _, item := range this.Targets {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Targets", err)
			}
		}
	}
	return nil
}
func (this *OpInput) Validate() error {
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
	if this.Alias != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return go_proto_validators.FieldError("Alias", err)
		}
	}
	if this.Symbol != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Symbol); err != nil {
			return go_proto_validators.FieldError("Symbol", err)
		}
	}
	for _, item := range this.Targets {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Targets", err)
			}
		}
	}
	return nil
}
func (this *OpOutput) Validate() error {
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
	if this.Alias != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Alias); err != nil {
			return go_proto_validators.FieldError("Alias", err)
		}
	}
	if this.Symbol != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Symbol); err != nil {
			return go_proto_validators.FieldError("Symbol", err)
		}
	}
	return nil
}
func (this *OpSource) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Upstream != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Upstream); err != nil {
			return go_proto_validators.FieldError("Upstream", err)
		}
	}
	return nil
}
func (this *OpGroup) Validate() error {
	if this.Id != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return go_proto_validators.FieldError("Id", err)
		}
	}
	for _, item := range this.Inputs {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Inputs", err)
			}
		}
	}
	for _, item := range this.Outputs {
		if item != nil {
			if err := go_proto_validators.CallValidatorIfExists(item); err != nil {
				return go_proto_validators.FieldError("Outputs", err)
			}
		}
	}
	return nil
}
