package interactor

import (
	"context"
	"time"

	"github.com/KevinArangu/stats_server/service"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const sleepTime = time.Duration(10 * time.Second)

var log = logrus.New()

func StartPing(ctx context.Context, e *echo.Echo, r *service.Redis, p *service.Pinger) {
	log.Info("Start pings")
	for {
		stats, err := p.PingLocal()
		if err != nil {
			log.WithError(err).Error("Ping Local error")
		}
		log.Info(stats)
		if stats.PacketLoss == 0 {
			r.AddPingComplete(ctx, stats)
		}

		stats, err = p.PingRemote()
		if err != nil {
			log.WithError(err).Error("Ping Remote error")
		}
		log.Info(stats)
		if stats.PacketLoss == 0 {
			r.AddPingError(ctx, stats)
		}

		select {
		case <-ctx.Done():
			break
		case <-time.After(sleepTime):
		}
	}

}
