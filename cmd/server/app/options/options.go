package options

import (
	cliflag "k8s.io/component-base/cli/flag"
)

// ServerOptions runs a api server.
type ServerOptions struct {
	ConfFile string
}

// NewServerOptions creates a new ServerOptions object with default parameters
func NewServerOptions() *ServerOptions {
	o := &ServerOptions{}
	return o
}

// Flags returns flags for a specific APIServer by section name
func (o *ServerOptions) Flags() (fss cliflag.NamedFlagSets) {
	fs := fss.FlagSet("server")
	fs.StringVarP(&o.ConfFile, "config", "c", "configs/config.yaml", "config file path.")
	return fss
}

// ValidateOptions validates ServerOptions
func (o *ServerOptions) ValidateOptions() error {
	return nil
}

// Complete set default ServerOptions.
// Should be called after flags parsed.
func (o *ServerOptions) Complete() error {
	return nil
}
