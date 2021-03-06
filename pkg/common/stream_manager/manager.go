package stream_manager

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	cored_pb "github.com/nayotta/metathings/pkg/proto/cored"
)

var (
	NotFound              = errors.New("not found")
	Timeout               = errors.New("timeout")
	Registered            = errors.New("registered")
	UnexpectedMessageType = errors.New("unexpected message type")
)

type StreamManager interface {
	Register(core_id string, stream cored_pb.CoredService_StreamServer) (chan interface{}, error)
	UnaryCall(core_id string, req *cored_pb.UnaryCallRequestPayload) (*cored_pb.UnaryCallResponsePayload, error)
	StreamCall(core_id string, req *cored_pb.StreamCallRequestPayload) (cored_pb.CoredService_StreamServer, func(), error)
	Close(core_id string) error
}

type streamManager struct {
	streams          map[string]cored_pb.CoredService_StreamServer
	stream_exit_chan map[string]chan interface{}
	sessions         map[string]chan *cored_pb.StreamResponse

	streaming_streams map[string]cored_pb.CoredService_StreamServer
	streaming_close   map[string]chan interface{}
	streaming_notity  map[string]chan interface{}

	lock   *sync.Mutex
	logger log.FieldLogger
}

func (mgr *streamManager) getSessionIdFromContext(ctx context.Context) string {
	return metautils.ExtractIncoming(ctx).Get("session-id")
}

func (mgr *streamManager) clear_default_stream(core_id string) error {
	if _, ok := mgr.streams[core_id]; ok {
		delete(mgr.streams, core_id)
		mgr.logger.WithField("core_id", core_id).Debugf("delete stream in stream manager")
	}
	if ch, ok := mgr.stream_exit_chan[core_id]; ok {
		close(ch)
		delete(mgr.stream_exit_chan, core_id)
		mgr.logger.WithField("core_id", core_id).Debugf("delete and close stream exit channel in stream manager")
	}
	return nil
}

func (mgr *streamManager) send_exit_signal(core_id string) error {
	if ch, ok := mgr.stream_exit_chan[core_id]; ok {
		select {
		case ch <- nil:
			mgr.logger.WithField("core_id", core_id).Debugf("send exit signal to Stream")
		default:
			mgr.logger.WithField("core_id", core_id).Debugf("send exit signal blocked, maybe closed")
		}
	} else {
		mgr.logger.WithField("core_id", core_id).Debugf("failed to send exit signal to Stream, maybe closed")
	}
	return nil
}

func (mgr *streamManager) register_default(core_id string, stream cored_pb.CoredService_StreamServer) (chan interface{}, error) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	if _, ok := mgr.streams[core_id]; ok {
		return nil, Registered
	}
	mgr.streams[core_id] = stream

	exit := make(chan interface{})
	mgr.stream_exit_chan[core_id] = exit

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				mgr.lock.Lock()
				defer mgr.lock.Unlock()
				mgr.send_exit_signal(core_id)
				mgr.clear_default_stream(core_id)
				gerr, ok := status.FromError(err)
				if (!ok && err != io.EOF) || (ok && gerr.Code() != codes.Canceled) {
					mgr.logger.WithError(err).WithField("core_id", core_id).Warningf("core agent stream closed with unexpected error")
				} else {
					mgr.logger.WithField("core_id", core_id).Debugf("core agent stream closed")
				}
				return
			}

			if ch, ok := mgr.sessions[res.SessionId]; !ok {
				mgr.logger.WithField("session_id", res.SessionId).Errorf("unknown session id")
			} else {
				ch <- res
			}
		}

	}()

	return exit, nil
}

func (mgr *streamManager) register_streaming(core_id string, sess_id string, stream cored_pb.CoredService_StreamServer) (chan interface{}, error) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	notity, ok := mgr.streaming_notity[sess_id]
	if !ok {
		mgr.logger.WithField("session_id", sess_id).Errorf("invalid session id")
		return nil, NotFound
	}

	if _, ok := mgr.streaming_streams[sess_id]; ok {
		return nil, Registered
	}
	mgr.streaming_streams[sess_id] = stream

	close_ch := make(chan interface{})
	mgr.streaming_close[sess_id] = close_ch
	notity <- nil

	return close_ch, nil
}

func (mgr *streamManager) Register(core_id string, stream cored_pb.CoredService_StreamServer) (chan interface{}, error) {
	ctx := stream.Context()
	sess_id := mgr.getSessionIdFromContext(ctx)

	if sess_id == "" {
		mgr.logger.WithFields(log.Fields{"core_id": core_id}).Debugf("register default stream")
		return mgr.register_default(core_id, stream)
	} else {
		mgr.logger.WithFields(log.Fields{"core_id": core_id, "session_id": sess_id}).Debugf("register streaming stream")
		return mgr.register_streaming(core_id, sess_id, stream)
	}
}

func (mgr *streamManager) UnaryCall(core_id string, req *cored_pb.UnaryCallRequestPayload) (*cored_pb.UnaryCallResponsePayload, error) {
	stream, ok := mgr.streams[core_id]
	if !ok {
		mgr.logger.WithField("core_id", core_id).Warningf("core not found")
		return nil, NotFound
	}

	sess_id := id_helper.NewId()
	ch := make(chan *cored_pb.StreamResponse)
	mgr.sessions[sess_id] = ch

	stm_req := &cored_pb.StreamRequest{
		SessionId:   &gpb.StringValue{Value: sess_id},
		MessageType: cored_pb.StreamMessageType_STREAM_MESSAGE_TYPE_USER,
		Payload:     &cored_pb.StreamRequest_UnaryCall{req},
	}

	if err := stream.Send(stm_req); err != nil {
		mgr.logger.WithError(err).Errorf("failed to send unary call")
		return nil, err
	}

	defer func() {
		close(ch)
		delete(mgr.sessions, sess_id)
		mgr.logger.WithField("session_id", sess_id).Debugf("close session receive channel")
	}()
	select {
	case stm_res := <-ch:
		switch stm_res.Payload.(type) {
		case *cored_pb.StreamResponse_UnaryCall:
			return stm_res.GetUnaryCall(), nil
		case *cored_pb.StreamResponse_Err:
			return nil, errors.New(stm_res.GetErr().GetContext())
		default:
			return nil, UnexpectedMessageType
		}
	case <-time.After(30 * time.Second):
		return nil, Timeout
	}
}

func (mgr *streamManager) StreamCall(core_id string, req *cored_pb.StreamCallRequestPayload) (cored_pb.CoredService_StreamServer, func(), error) {
	agstm, ok := mgr.streams[core_id]
	if !ok {
		mgr.logger.WithField("core_id", core_id).Warningf("core stream not foound")
		return nil, nil, NotFound
	}

	sess_id := id_helper.NewId()
	ch := make(chan *cored_pb.StreamResponse)
	notity := make(chan interface{})
	mgr.sessions[sess_id] = ch
	mgr.streaming_notity[sess_id] = notity

	stm_req := &cored_pb.StreamRequest{
		SessionId:   &gpb.StringValue{Value: sess_id},
		MessageType: cored_pb.StreamMessageType_STREAM_MESSAGE_TYPE_SYSTEM,
		Payload: &cored_pb.StreamRequest_StreamCall{
			StreamCall: req,
		},
	}

	if err := agstm.Send(stm_req); err != nil {
		mgr.logger.WithError(err).Errorf("failed to send stream call config")
		return nil, nil, err
	}
	mgr.logger.WithFields(log.Fields{
		"core_id":    core_id,
		"session_id": sess_id,
	}).Debugf("send stream call config request to core agent")

	defer func() {
		close(notity)
		delete(mgr.streaming_notity, sess_id)
		mgr.logger.WithField("session_id", sess_id).Debugf("close session receive notity channel")
	}()
	select {
	case <-ch:
		mgr.logger.WithField("session_id", sess_id).Debugf("receive stream call config response from core agent")
	case <-time.After(30 * time.Second):
		mgr.logger.WithField("session_id", sess_id).Errorf("receive stream call config response timeout")
		return nil, nil, Timeout
	}

	select {
	case <-notity:
		stm, ok := mgr.streaming_streams[sess_id]
		if !ok {
			mgr.logger.WithFields(log.Fields{
				"core_id":    core_id,
				"session_id": sess_id,
			}).Errorf("core streaming stream not found")
			return nil, nil, NotFound
		}
		return stm, func() { mgr.streaming_close[sess_id] <- nil }, nil
	case <-time.After(10 * time.Second):
		mgr.logger.WithField("session_id", sess_id).Errorf("receive stream call confirm timeout")
		return nil, nil, Timeout
	}
}

func (mgr *streamManager) Close(core_id string) error {
	var ch chan interface{}
	var ok bool

	if ch, ok = mgr.stream_exit_chan[core_id]; !ok {
		return NotFound
	}
	ch <- nil
	return nil
}

func NewStreamManager(logger log.FieldLogger) (StreamManager, error) {
	mgr := &streamManager{
		streams:          make(map[string]cored_pb.CoredService_StreamServer),
		stream_exit_chan: make(map[string]chan interface{}),
		sessions:         make(map[string]chan *cored_pb.StreamResponse),

		streaming_streams: make(map[string]cored_pb.CoredService_StreamServer),
		streaming_close:   make(map[string]chan interface{}),
		streaming_notity:  make(map[string]chan interface{}),

		lock:   new(sync.Mutex),
		logger: logger,
	}

	return mgr, nil
}
