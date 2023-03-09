package service

import (
	"time"

	"github.com/go-ping/ping"
)

const (
	LocalPingerAddress  = "192.168.0.1"
	RemotePingerAddress = "1.1.1.1"
	PingDuration        = time.Duration(1 * time.Second)
	PingCount           = 1
)

type Pinger struct{}

var pinger Pinger

func NewPinger() *Pinger {
	return &pinger
}

func (p *Pinger) PingLocal() (*ping.Statistics, error) {
	local, err := ping.NewPinger(LocalPingerAddress)
	if err != nil {
		return nil, err
	}
	local.Count = PingCount
	local.Timeout = PingDuration
	local.SetPrivileged(true)
	err = local.Run()
	if err != nil {
		return nil, err
	}
	stats := local.Statistics()
	return stats, nil
}

func (p *Pinger) PingRemote() (*ping.Statistics, error) {
	remote, err := ping.NewPinger(RemotePingerAddress)
	if err != nil {
		return nil, err
	}
	remote.Count = PingCount
	remote.Timeout = PingDuration
	remote.SetPrivileged(true)
	err = remote.Run()
	if err != nil {
		return nil, err
	}
	stats := remote.Statistics()
	return stats, nil
}
