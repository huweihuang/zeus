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
		Use:   "gin-api-frame",
		Short: "gin-api-frame apiserver",
		Long:  "gin-api-frame apiserver",
		RunE: func(cmd *cobra.Command, _ []string) error {
			verflag.PrintAndExitIfRequested()
			cliflag.PrintFlags(cmd.Flags())

			if err := opts.ValidateOptions(); err != nil {
				return fmt.Errorf("validate options: %v", err)
			}

			err := opts.Complete()
			if err != nil {
				return fmt.Errorf("complete gin-api-frame options error, %v", err)
			}

			if err := run(opts); err != nil {
				return fmt.Errorf("run gin-api-frame failed, %v", err)
			}
			return nil
		},
	}

	fs := cmd.Flags()
	namedFlagSets := opts.Flags()
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	})

	return cmd
}

// run runs the ServerConfig. This should never exit
func run(opt *options.ServerOptions) error {
	conf := config.MustLoad(opt.ConfFile)
	s := server.NewAPIServer(conf)
	return s.Run()
}
