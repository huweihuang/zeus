package client

import "github.com/huweihuang/zeus/cmd/server/app/configs"

var GlobalClients *Clients

type Clients struct {
}

func NewClients(c *configs.ClientConfig) (*Clients, error) {
	return GlobalClients, nil
}
