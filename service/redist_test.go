package service

import (
	"context"
	"testing"
)

func TestRedisConnection(t *testing.T) {
	r := NewRedis()
	r.client.RPush(context.TODO(), "algo", "algo")
}
