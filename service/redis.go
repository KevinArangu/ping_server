package service

import (
	"context"

	"github.com/KevinArangu/stats_server/config"
	"github.com/KevinArangu/stats_server/model"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

const (
	totalPings           = "total_pings"
	totalLocalPings      = "total_local_pings"
	totalRemotePings     = "total_remote_pings"
	completedLocalPings  = "completed_local_pings"
	completedRemotePings = "completed_remote_pings"
	errorsLocalPings     = "errors_local_pings"
	errorsRemotePings    = "errors_remote_pings"
)

var log = logrus.New()

type Redis struct {
	client *redis.Client
}

func NewRedis() (*Redis, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress(),
		Password: config.RedisPass(),
		DB:       config.RedisDB(),
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := c.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Redis{
		client: c,
	}, nil
}

func (r *Redis) AddPingComplete(ctx context.Context, isRemotePing bool) (int64, error) {
	var result int64
	var err error
	if isRemotePing {
		t := r.client.Incr(ctx, totalPings)
		if t.Err() != nil {
			log.WithError(t.Err()).Error("Error in AddPingComplete (totalPings)")
		}
		n := r.client.Incr(ctx, totalRemotePings)
		if n.Err() != nil {
			log.WithError(n.Err()).Error("Error in AddPingComplete (totalRemotePings)")
		}
		p := r.client.Incr(ctx, completedRemotePings)
		if p.Err() != nil {
			log.WithError(p.Err()).Error("Error in AddPingComplete (completedRemotePings)")
			return 0, p.Err()
		}
		res, er := p.Result()
		result = res
		err = er
	} else {
		t := r.client.Incr(ctx, totalPings)
		if t.Err() != nil {
			log.WithError(t.Err()).Error("Error in AddPingComplete (totalPings)")
		}
		n := r.client.Incr(ctx, totalLocalPings)
		if n.Err() != nil {
			log.WithError(n.Err()).Error("Error in AddPingComplete (totalLocalPings)")
		}
		p := r.client.Incr(ctx, completedLocalPings)
		if p.Err() != nil {
			log.WithError(p.Err()).Error("Error in AddPingComplete (completedLocalPings)")
			return 0, p.Err()
		}
		res, er := p.Result()
		result = res
		err = er
	}
	return result, err
}

func (r *Redis) AddPingError(ctx context.Context, isRemotePing bool) (int64, error) {
	var result int64
	var err error
	if isRemotePing {
		t := r.client.Incr(ctx, totalPings)
		if t.Err() != nil {
			log.WithError(t.Err()).Error("Error in AddPingError (totalPings)")
		}
		n := r.client.Incr(ctx, totalRemotePings)
		if n.Err() != nil {
			log.WithError(n.Err()).Error("Error in AddPingError (totalRemotePings)")
		}
		p := r.client.Incr(ctx, errorsRemotePings)
		if p.Err() != nil {
			log.WithError(p.Err()).Error("Error in AddPingError (errorsRemotePings)")
			return 0, p.Err()
		}
		res, er := p.Result()
		result = res
		err = er
	} else {
		t := r.client.Incr(ctx, totalPings)
		if t.Err() != nil {
			log.WithError(t.Err()).Error("Error in AddPingError (totalPings)")
		}
		n := r.client.Incr(ctx, totalLocalPings)
		if n.Err() != nil {
			log.WithError(n.Err()).Error("Error in AddPingError (totalLocalPings)")
		}
		p := r.client.Incr(ctx, errorsLocalPings)
		if p.Err() != nil {
			log.WithError(p.Err()).Error("Error in AddPingError (errorsLocalPings)")
			return 0, p.Err()
		}
		res, er := p.Result()
		result = res
		err = er
	}
	return result, err
}

func (r *Redis) GetStats(ctx context.Context) (model.StatsResponse, error) {
	var tp, tlp, trp, clp, crp, elp, erp string
	mg := r.client.MGet(ctx, totalPings, totalLocalPings, totalRemotePings, completedLocalPings, completedRemotePings, errorsLocalPings, errorsRemotePings)
	if mg.Err() != nil {
		return model.StatsResponse{}, mg.Err()
	}
	values := mg.Val()
	if values[0] != nil {
		tp = values[0].(string)
	}
	if values[1] != nil {
		tlp = values[1].(string)
	}
	if values[2] != nil {
		trp = values[2].(string)
	}
	if values[3] != nil {
		clp = values[3].(string)
	}
	if values[4] != nil {
		crp = values[4].(string)
	}
	if values[5] != nil {
		elp = values[5].(string)
	}
	if values[6] != nil {
		erp = values[6].(string)
	}

	return model.StatsResponse{
		TotalPings:           tp,
		TotalLocalPings:      tlp,
		TotalRemotePings:     trp,
		CompletedLocalPings:  clp,
		CompletedRemotePings: crp,
		ErrorLocalPings:      elp,
		ErrorRemotePings:     erp,
	}, nil
}
