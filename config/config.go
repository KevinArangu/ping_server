package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	PortHTTP int `env:"PORT_HTTP,required"`
}

var c config

func ReadConfig() error {
	ctx := context.Background()
	return envconfig.Process(ctx, &c)
}

func PortHTTP() int {
	return c.PortHTTP
}

func ServerAddress() string {
	return fmt.Sprintf(":%d", PortHTTP())
}
