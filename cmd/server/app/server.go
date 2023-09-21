package app

import (
	"fmt"

	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/term"

	"github.com/huweihuang/gin-api-frame/cmd/server/app/config"
	"github.com/huweihuang/gin-api-frame/cmd/server/app/options"
	"github.com/huweihuang/gin-api-frame/pkg/server"
	"github.com/huweihuang/gin-api-frame/pkg/version/verflag"
)

// NewServerCommand creates gin-api-frame command
func NewServerCommand() *cobra.Command {
	opts := options.NewServerOptions()
	cmd := &cobra.Command{
		Use:  "gin-api-frame",
		Long: "gin-api-frame api server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			verflag.PrintAndExitIfRequested()
			cliflag.PrintFlags(cmd.Flags())

			if err := opts.ValidateOptions(); err != nil {
				return fmt.Errorf("validate options error: %v", err)
			}

			if err := opts.Complete(); err != nil {
				return fmt.Errorf("complete options error, %v", err)
			}

			return Run(opts)
		},
	}

	fs := cmd.Flags()
	namedFlagSets := opts.Flags()
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, namedFlagSets, cols)

	return cmd
}

// Run runs the ServerConfig. This should never exit
func Run(opt *options.ServerOptions) error {
	conf, err := config.InitConfig(opt.ConfFile)
	if err != nil {
		return err
	}
	s := server.NewServer(conf)
	return s.Run()
}
