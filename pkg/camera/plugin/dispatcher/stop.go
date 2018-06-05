package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/nayotta/metathings/pkg/proto/camera"
)

func unary_stop(cli pb.CameraServiceClient, ctx context.Context, req *any.Any) (*any.Any, error) {
	res, err := cli.Stop(ctx, &empty.Empty{})
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
	unary_call_methods["Stop"] = unary_stop
}
