package service

import (
	"github.com/KevinArangu/stats_server/config"
	"github.com/go-ping/ping"
)

type Pinger struct{}

var pinger Pinger

func NewPinger() *Pinger {
	return &pinger
}

func (p *Pinger) PingLocal() (*ping.Statistics, error) {
	local, err := ping.NewPinger(config.LocalPingAddress())
	if err != nil {
		log.WithField("addr", config.LocalPingAddress()).Error("Error in PingLocal")
		return nil, err
	}
	local.Count = config.PingCount()
	local.Timeout = config.PingDuration()
	local.SetPrivileged(true)
	err = local.Run()
	if err != nil {
		return nil, err
	}
	stats := local.Statistics()
	return stats, nil
}

func (p *Pinger) PingRemote() (*ping.Statistics, error) {
	remote, err := ping.NewPinger(config.RemotePingAddress())
	if err != nil {
		log.WithField("addr", config.RemotePingAddress()).Error("Error in PingRemote")
		return nil, err
	}
	remote.Count = config.PingCount()
	remote.Timeout = config.PingDuration()
	remote.SetPrivileged(true)
	err = remote.Run()
	if err != nil {
		return nil, err
	}
	stats := remote.Statistics()
	return stats, nil
}
