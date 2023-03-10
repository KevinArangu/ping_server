package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisConnection(t *testing.T) {
	r, err := NewRedis()
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	rp := r.client.RPush(context.TODO(), "algo", "algo")
	res, err := rp.Result()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	lp := r.client.LPop(context.TODO(), "algo")
	ress, err := lp.Result()
	assert.NoError(t, err)
	assert.NotEmpty(t, ress)
}
