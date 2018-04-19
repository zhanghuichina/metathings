package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bigdatagz/metathings/pkg/common/cmd/helper"
)

const (
	METATHINGS_PREFIX = "MT_"
)

var (
	V helper.GetArgument
	A helper.WithPrefix
)

var (
	root_opts struct {
		verbose                       bool
		addr                          string
		token                         string
		log_level                     string
		application_credential_id     string
		application_credential_secret string
	}
)

var (
	RootCmd = &cobra.Command{
		Use:   "metathings",
		Short: "MetaThings Command Line Toolkits",
	}
)

func init() {
	h := helper.NewArgumentHelper(METATHINGS_PREFIX)
	V = h.GetString
	A = h.PrefixWith
	viper.AutomaticEnv()

	RootCmd.PersistentFlags().BoolVar(&root_opts.verbose, "verbose", false, "Verbose mode")

	RootCmd.PersistentFlags().StringVar(&root_opts.token, "mt-token", "", "MetaThings Token")
	viper.BindPFlag(A("TOKEN"), RootCmd.PersistentFlags().Lookup("mt-token"))

	RootCmd.PersistentFlags().StringVar(&root_opts.log_level, "log-level", "info", "Logging Level[debug, info, warn, error]")
	viper.BindPFlag(A("LOG_LEVEL"), RootCmd.PersistentFlags().Lookup("log-level"))

	RootCmd.PersistentFlags().StringVar(&root_opts.addr, "addr", "", "Metathings Service Address")
	viper.BindPFlag(A("ADDR"), RootCmd.PersistentFlags().Lookup("addr"))

	RootCmd.PersistentFlags().StringVar(&root_opts.application_credential_id, "mt-application-credential-id", "", "MetaThings Application Credential ID")
	viper.BindPFlag(A("APPLICATION_CREDENTIAL_ID"), RootCmd.PersistentFlags().Lookup("mt-application-credential-id"))

	RootCmd.PersistentFlags().StringVar(&root_opts.application_credential_secret, "mt-application-credential-secret", "", "MetaThings Application Credential Secret")
	viper.BindPFlag(A("APPLICATION_CREDENTIAL_SECRET"), RootCmd.PersistentFlags().Lookup("mt-application-credential-secret"))
}

func initialize() {
	lvl, err := log.ParseLevel(V("log_level"))
	if err != nil {
		log.Fatalf("bad log level %v: %v", V("log_level"), err)
	}
	log.SetLevel(lvl)
}

func globalPreRunHook(cmd *cobra.Command, args []string) {
	initialize()
}
