package cmd

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	pb "github.com/nayotta/metathings/pkg/proto/streamd"
	service "github.com/nayotta/metathings/pkg/streamd/service"
)

type _streamdOptions struct {
	_rootOptions  `mapstructure:",squash"`
	Listen        string
	Storage       cmd_helper.StorageOptions
	ServiceConfig cmd_helper.ServiceConfigOptions `mapstructure:"service_config"`
}

var (
	streamd_opts *_streamdOptions
)

var (
	streamdCmd = &cobra.Command{
		Use:   "streamd",
		Short: "Streamd Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			opts_t := &_streamdOptions{}
			cmd_helper.UnmarshalConfig(opts_t)
			cmd_helper.InitServiceConfigOptions(&opts_t.ServiceConfig, &streamd_opts.ServiceConfig)
			root_opts = &opts_t._rootOptions
			streamd_opts = opts_t
			streamd_opts.Service = "streamd"
			streamd_opts.Stage = cmd_helper.GetStageFromEnv()
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runStreamd(); err != nil {
				log.WithError(err).Fatalf("failed to run streamd")
			}
		},
	}
)

func runStreamd() error {
	lis, err := net.Listen("tcp", streamd_opts.Listen)
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(nil)),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(nil)),
	)
	srv, err := service.NewStreamdService(
		service.SetLogLevel(streamd_opts.Log.Level),
		service.SetMetathingsdAddr(streamd_opts.ServiceConfig.Metathingsd.Address),
		service.SetIdentitydAddr(streamd_opts.ServiceConfig.Identityd.Address),
		service.SetCoredAddr(streamd_opts.ServiceConfig.Cored.Address),
		service.SetStorage(streamd_opts.Storage.Driver, streamd_opts.Storage.Uri),
		service.SetApplicationCredential(
			streamd_opts.ApplicationCredential.Id,
			streamd_opts.ApplicationCredential.Secret,
		),
		service.SetStreamManager(cmd_helper.GetFromStage().Sub("stream_manager")),
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new stream service")
		return err
	}

	pb.RegisterStreamdServiceServer(s, srv)

	log.WithField("listen", streamd_opts.Listen).Infof("metathings stream service listening")
	return s.Serve(lis)
}

func init() {
	streamd_opts = &_streamdOptions{}

	streamdCmd.Flags().StringVarP(&streamd_opts.Listen, "listen", "l", "127.0.0.1:5004", "Metathings Stream Service listening address")
	streamdCmd.Flags().StringVar(&streamd_opts.ServiceConfig.Identityd.Address, "identityd-addr", constant_helper.CONSTANT_METATHINGSD_DEFAULT_HOST, "MetaThings Identity Service Address")
	streamdCmd.Flags().StringVar(&streamd_opts.ServiceConfig.Cored.Address, "cored-addr", constant_helper.CONSTANT_METATHINGSD_DEFAULT_HOST, "MetaThings Core Service Address")
	streamdCmd.Flags().StringVar(&streamd_opts.ServiceConfig.Sensord.Address, "sensord-addr", constant_helper.CONSTANT_METATHINGSD_DEFAULT_HOST, "MetaThings Sensor Service Address")
	streamdCmd.Flags().StringVar(&streamd_opts.ServiceConfig.Metathingsd.Address, "metathingsd-addr", constant_helper.CONSTANT_METATHINGSD_DEFAULT_HOST, "MetaThings Service Address")
	streamdCmd.Flags().StringVar(&streamd_opts.Storage.Driver, "storage-driver", "sqlite3", "Storage Driver [sqlite3]")
	streamdCmd.Flags().StringVar(&streamd_opts.Storage.Uri, "storage-uri", "", "Storage URI")

	RootCmd.AddCommand(streamdCmd)
}
