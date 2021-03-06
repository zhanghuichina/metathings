package metathings_cored_service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	pb "github.com/nayotta/metathings/pkg/proto/cored"
)

func isStreamCallConfigRequestPayload(req *pb.StreamCallRequestPayload) bool {
	_, ok := req.Payload.(*pb.StreamCallRequestPayload_Config)
	return ok
}

func isStreamCallDataRequestPayload(req *pb.StreamCallRequestPayload) bool {
	_, ok := req.Payload.(*pb.StreamCallRequestPayload_Data)
	return ok
}

func isStreamCallConfigResponsePayload(res *pb.StreamResponse) bool {
	res1, ok := res.Payload.(*pb.StreamResponse_StreamCall)
	if !ok {
		return false
	}
	_, ok = res1.StreamCall.Payload.(*pb.StreamCallResponsePayload_Config)
	return ok
}

func isStreamCallDataResponsePayload(res *pb.StreamResponse) bool {
	res1, ok := res.Payload.(*pb.StreamResponse_StreamCall)
	if !ok {
		return false
	}

	_, ok = res1.StreamCall.Payload.(*pb.StreamCallResponsePayload_Data)
	return ok
}

func (srv *metathingsCoredService) GetSessionIdFromContext(ctx context.Context) string {
	return metautils.ExtractIncoming(ctx).Get("session-id")
}

func (srv *metathingsCoredService) handleGRPCError(err error, format string, args ...interface{}) error {
	return grpc_helper.HandleGRPCError(srv.logger, err, format, args)
}
