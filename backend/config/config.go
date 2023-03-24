package config

import (
	"context"
	"fmt"
	"time"

	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type config struct {
	PortHTTP          int    `env:"PORT_HTTP,required"`
	RedisAddress      string `env:"REDIS_ADDRESS,required"`
	RedisPass         string `env:"REDIS_PASS,required"`
	RedisDB           int    `env:"REDIS_DB,required"`
	LocalPingAddress  string `env:"LOCAL_PING_ADDRESS,required"`
	RemotePingAddress string `env:"REMOTE_PING_ADDRESS,required"`
	PingDuration      int    `env:"PING_DURATION,required"`
	PingCount         int    `env:"PING_COUNT,required"`
	SleepTime         int    `env:"SLEEP_TIME,required"`
	FrontendDirectory string `env:"FRONTEND_DIRECTORY,required"`
}

var c config

func ReadConfig() error {
	ctx := context.Background()
	err := envconfig.Process(ctx, &c)
	log.WithField("config", c).Info("Envs loaded")
	return err
}

func PortHTTP() int {
	return c.PortHTTP
}

func ServerAddress() string {
	return fmt.Sprintf(":%d", PortHTTP())
}

func RedisAddress() string {
	return c.RedisAddress
}

func RedisPass() string {
	return c.RedisPass
}

func RedisDB() int {
	return c.RedisDB
}

func LocalPingAddress() string {
	return c.LocalPingAddress
}

func RemotePingAddress() string {
	return c.RemotePingAddress
}

func PingDuration() time.Duration {
	return time.Duration(time.Duration(c.PingDuration) * time.Second)
}

func PingCount() int {
	return c.PingCount
}

func SleepTime() time.Duration {
	return time.Duration(time.Duration(c.SleepTime) * time.Second)
}

func FrontendDirectory() string {
	return c.FrontendDirectory
}
