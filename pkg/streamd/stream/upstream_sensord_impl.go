package stream_manager

import (
	"context"
	"fmt"
	"sync"
	"time"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/lovoo/goka"
	log "github.com/sirupsen/logrus"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	sensor_pb "github.com/nayotta/metathings/pkg/proto/sensor"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
)

type sensordUpstreamOption struct {
	id    string
	alias string

	logger       log.FieldLogger
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	cli_fty      *client_helper.ClientFactory
	snr_id       string
	brokers      []string
	targets      []string
	filters      map[string]string
	sym_tbl      SymbolTable
}

type sensordUpstream struct {
	Emitter
	slck     *sync.Mutex
	logger   log.FieldLogger
	state    UpstreamState
	opt      *sensordUpstreamOption
	cfn      client_helper.CloseFn
	emitters map[string]*goka.Emitter
}

func (self *sensordUpstream) Id() string {
	return self.opt.id
}

func (self *sensordUpstream) Symbol() string {
	sym := NewSymbol(self.Id(), COMPONENT_UPSTREAM, self.opt.alias)
	return sym.String()
}

func (self *sensordUpstream) Start() error {
	self.slck.Lock()
	defer self.slck.Unlock()
	if self.state != UPSTREAM_STATE_STOP {
		return ErrUnstartable
	}

	self.state = UPSTREAM_STATE_STARTING

	cli, cfn, err := self.opt.cli_fty.NewSensordServiceClient()
	if err != nil {
		self.state = UPSTREAM_STATE_STOP
		self.Emit(ERROR_EVENT, nil)
		return err
	}
	self.cfn = cfn

	go self.start(cli, cfn)
	self.logger.Debugf("upstream.sensord starting")

	return nil
}

func (self *sensordUpstream) start(cli sensord_pb.SensordServiceClient, cfn client_helper.CloseFn) {
	defer func() {
		self.slck.Lock()
		defer self.slck.Unlock()

		self.state = UPSTREAM_STATE_STOP
		self.Emit(STOP_EVENT, nil)

		self.logger.WithField("sensor_id", self.opt.snr_id).Infof("upstream.sensord terminated")
	}()

	ctx := context.Background()
	tk_ctx := context_helper.WithToken(ctx, self.opt.app_cred_mgr.GetToken())

	stm, err := cli.Subscribe(tk_ctx)
	if err != nil {
		self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Errorf("failed to subscribe")
		return
	}

	sub_reqs := &sensord_pb.SubscribeRequests{
		Requests: []*sensord_pb.SubscribeRequest{
			&sensord_pb.SubscribeRequest{
				Payload: &sensord_pb.SubscribeRequest_SubscribeById{
					SubscribeById: &sensord_pb.SubscribeByIdRequest{
						Id: &gpb.StringValue{Value: self.opt.snr_id},
					},
				},
			},
		},
	}

	err = stm.Send(sub_reqs)
	if err != nil {
		self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Errorf("failed to subscribe sesnor data")
		return
	}
	self.logger.WithFields(log.Fields{
		"sensor_id": self.opt.snr_id,
	}).Debugf("upstream.sensord started")

	self.slck.Lock()
	self.state = UPSTREAM_STATE_RUNNING
	self.Emit(START_EVENT, nil)
	self.slck.Unlock()

	for {
		sub_ress, err := stm.Recv()
		if err != nil {
			self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Errorf("failed to recv data from stream")
			return
		}
		self.logger.WithFields(log.Fields{
			"upstream_id": self.Id(),
			"sensor_id":   self.opt.snr_id,
		}).Debugf("receive data from sensord")

		for _, sub_res := range sub_ress.Responses {
			self.logger.Debugf("%v", sub_res.Data)
			upstm_dat := enc_sensord_upstream_data(sub_res.Data)
			for target, filter := range self.opt.filters {
				ok, err := self.filter_upstream_data(filter, upstm_dat)
				if err != nil {
					self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Warningf("failed to filter upstream data")
				} else if ok {
					if err = self.emit_upstream_data(target, upstm_dat); err != nil {
						self.logger.WithError(err).WithField("sensor_id", self.opt.snr_id).Warningf("failed to emit upstream data")
					}
				}
			}
		}
	}
}

// TODO(Peer): dont compile filter every time
func (self *sensordUpstream) filter_upstream_data(filter string, upstream_data *UpstreamData) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	eng := NewLuaEngine(self.opt.logger)
	defer eng.Close()

	eng.SetContext(ctx)
	return eng.Filter(filter, upstream_data.Metadata().Data(), upstream_data.Data())

}

func (self *sensordUpstream) emit_upstream_data(target string, upstream_data *UpstreamData) error {
	sym := self.opt.sym_tbl.Lookup(target)

	var codec goka.Codec
	var msg interface{}

	switch sym.Component() {
	case COMPONENT_INPUT:
		input_data := UpstreamDataToInputData(upstream_data)
		input_data.Metadata().Set("from", self.Symbol())
		msg = input_data
		codec = new(InputDataCodec)
	case COMPONENT_OUTPUT:
		output_data := UpstreamDataToOutputData(upstream_data)
		output_data.Metadata().Set("from", self.Symbol())
		msg = output_data
		codec = new(OutputDataCodec)
	}

	var emitter *goka.Emitter
	var ok bool
	var err error

	if emitter, ok = self.emitters[sym.String()]; !ok {
		emitter, err = goka.NewEmitter(self.opt.brokers, goka.Stream(sym.String()), codec)
		if err != nil {
			return err
		}
		self.emitters[sym.String()] = emitter

		self.logger.WithFields(log.Fields{
			"brokers": self.opt.brokers,
			"symbol":  sym.String(),
		}).Debugf("create goka emitter")
	}

	err = emitter.EmitSync("upstream.sensord", msg)
	if err != nil {
		return err
	}
	self.logger.WithField("symbol", sym.String()).Debugf("emit upstream data")

	return nil
}

func enc_sensord_upstream_data(snr_dat *sensord_pb.SensorData) *UpstreamData {
	md := map[string]interface{}{}
	md["sensor_id"] = snr_dat.SensorId
	md["sensor_name"] = snr_dat.Data["$sensor.name"].GetString_()
	md["created_at"] = fmt.Sprint(protobuf_helper.ToTime(*snr_dat.CreatedAt).UnixNano())
	md["arrvied_at"] = fmt.Sprint(protobuf_helper.ToTime(*snr_dat.ArrivedAt).UnixNano())

	d := map[string]interface{}{}
	for k, v := range snr_dat.Data {
		if len(k) > 0 && k[0] != '$' {
			d[k] = dec_sensord_data_value(v)
		}
	}

	return NewUpstreamData(d, md)
}

func dec_sensord_data_value(v *sensor_pb.SensorValue) string {
	switch v.Value.(type) {
	case *sensor_pb.SensorValue_Double:
		return fmt.Sprint(v.GetDouble())
	case *sensor_pb.SensorValue_Float:
		return fmt.Sprint(v.GetFloat())
	case *sensor_pb.SensorValue_Int64:
		return fmt.Sprint(v.GetInt64())
	case *sensor_pb.SensorValue_Uint64:
		return fmt.Sprint(v.GetUint64())
	case *sensor_pb.SensorValue_Int32:
		return fmt.Sprint(v.GetInt32())
	case *sensor_pb.SensorValue_Uint32:
		return fmt.Sprint(v.GetUint32())
	case *sensor_pb.SensorValue_String_:
		return v.GetString_()
	default:
		return ""
	}
}

func (self *sensordUpstream) Stop() error {
	self.slck.Lock()
	defer self.slck.Unlock()

	if self.state != UPSTREAM_STATE_RUNNING {
		return ErrUnterminable
	}

	self.state = UPSTREAM_STATE_TERMINATING
	self.cfn()

	for _, emitter := range self.emitters {
		emitter.Finish()
	}

	self.logger.WithField("sensor_id", self.opt.snr_id).Debugf("upstream.sensord terminating")
	return nil
}

func (self *sensordUpstream) State() UpstreamState {
	self.slck.Lock()
	defer self.slck.Unlock()

	return self.state
}

func (self *sensordUpstream) Close() {
	self.slck.Lock()
	defer self.slck.Unlock()

	for _, emitter := range self.emitters {
		emitter.Finish()
	}

	self.logger.WithField("sensor_id", self.opt.snr_id).Debugf("upstream closed")
}

type sensordUpstreamFactory struct {
	opt *sensordUpstreamOption
}

func (self *sensordUpstreamFactory) Set(key string, val interface{}) UpstreamFactory {
	switch key {
	case "logger":
		self.opt.logger = val.(log.FieldLogger)
	case "application_credential_manager":
		self.opt.app_cred_mgr = val.(app_cred_mgr.ApplicationCredentialManager)
	case "client_factory":
		self.opt.cli_fty = val.(*client_helper.ClientFactory)
	case "symbol_table":
		self.opt.sym_tbl = val.(SymbolTable)
	case "brokers":
		self.opt.brokers = val.([]string)
	case "option":
		opt := val.(*UpstreamOption)
		self.opt.id = opt.Id
		self.opt.alias = opt.Alias
		self.opt.snr_id = opt.Config["sensor_id"]
		self.opt.targets = split_and_trim(opt.Config["targets"])
		self.opt.filters = group_by_prefix(opt.Config, "filter.")
	}

	return self
}

func (self *sensordUpstreamFactory) New() (Upstream, error) {
	upstream := &sensordUpstream{
		Emitter: NewEmitter(),
		slck:    &sync.Mutex{},
		logger: self.opt.logger.WithFields(log.Fields{
			"id":         self.opt.id,
			"#component": "upstream:sensord",
		}),
		state:    UPSTREAM_STATE_STOP,
		opt:      self.opt,
		emitters: map[string]*goka.Emitter{},
	}
	return upstream, nil
}

func init() {
	RegisterUpstreamFactory("sensord", func() UpstreamFactory {
		return &sensordUpstreamFactory{opt: &sensordUpstreamOption{}}
	})
}
