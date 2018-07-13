package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	pb "github.com/nayotta/metathings/pkg/proto/sensor"
)

func unary_list_data(cli pb.SensorServiceClient, ctx context.Context, req *any.Any) (*any.Any, error) {
	req1 := new(pb.ListDataRequest)
	err := ptypes.UnmarshalAny(req, req1)
	if err != nil {
		return nil, err
	}

	res, err := cli.ListData(ctx, req1)
	if err != nil {
		return nil, err
	}

	res1, err := ptypes.MarshalAny(res)
	if err != nil {
		return nil, err
	}

	return res1, nil
}

func init() {
	unary_call_methods["ListData"] = unary_list_data
}
