package app

import (
	"fmt"

	"github.com/huweihuang/golib/config"
	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/term"

	"github.com/huweihuang/zeus/cmd/server/app/configs"
	"github.com/huweihuang/zeus/cmd/server/app/options"
	"github.com/huweihuang/zeus/pkg/server"
	"github.com/huweihuang/zeus/pkg/version/verflag"
)

// NewServerCommand creates zeus command
func NewServerCommand() *cobra.Command {
	opts := options.NewServerOptions()
	cmd := &cobra.Command{
		Use:  "zeus",
		Long: "zeus api server",
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
	err := config.InitConfigObjectByPath(opt.ConfFile, &configs.GlobalConfig)
	if err != nil {
		return err
	}
	s := server.NewServer(&configs.GlobalConfig)
	return s.Run()
}
