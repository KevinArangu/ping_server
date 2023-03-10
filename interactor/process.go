package interactor

import (
	"context"
	"time"

	"github.com/KevinArangu/stats_server/config"
	"github.com/KevinArangu/stats_server/service"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func StartPing(c context.Context, e *echo.Echo, r *service.Redis, p *service.Pinger) {
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	log.Info("Start pings")
	for {
		stats, err := p.PingLocal()
		if err != nil {
			log.WithError(err).Error("Ping Local error")
		}
		if stats != nil {
			log.Info(stats)
			if stats.PacketLoss == 0 {
				r.AddPingComplete(ctx, true)
			} else {
				r.AddPingError(ctx, true)
			}
		}

		stats, err = p.PingRemote()
		if err != nil {
			log.WithError(err).Error("Ping Remote error")
		}
		if stats != nil {
			log.Info(stats)
			if stats.PacketLoss == 0 {
				r.AddPingComplete(ctx, false)
			} else {
				r.AddPingError(ctx, false)
			}
		}

		select {
		case <-ctx.Done():
			break
		case <-time.After(config.SleepTime()):
		}
	}

}
