package service

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

const (
	timeout      = time.Duration(5 * time.Second)
	completedKey = "pings_completed"
	errorsKey    = "pings_errors"
)

var log = logrus.New()

type Redis struct {
	client *redis.Client
}

func NewRedis() *Redis {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Redis{
		client: c,
	}
}

func (r *Redis) AddPingComplete(c context.Context, ping interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	p := r.client.RPush(ctx, completedKey, ping)
	if p.Err() != redis.Nil {
		log.WithError(p.Err()).Error("Error in AddPingComplete")
		return 0, p.Err()
	}
	return p.Result()
}

func (r *Redis) AddPingError(c context.Context, ping interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	p := r.client.RPush(ctx, errorsKey, ping)
	if p.Err() != redis.Nil {
		log.WithError(p.Err()).Error("Error in AddPingError")
		return 0, p.Err()
	}
	return p.Result()
}
