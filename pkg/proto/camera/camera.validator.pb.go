// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: camera.proto

/*
Package camera is a generated protocol buffer package.

It is generated from these files:
	camera.proto
	service.proto
	show.proto
	start.proto
	stop.proto

It has these top-level messages:
	Dimension
	CameraConfig
	Camera
	ShowResponse
	StartConfig
	StartRequest
	StartResponse
	StopResponse
*/
package camera

import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Dimension) Validate() error {
	return nil
}
func (this *CameraConfig) Validate() error {
	if this.Dimension != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Dimension); err != nil {
			return go_proto_validators.FieldError("Dimension", err)
		}
	}
	return nil
}
func (this *Camera) Validate() error {
	if this.Config != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return go_proto_validators.FieldError("Config", err)
		}
	}
	return nil
}
