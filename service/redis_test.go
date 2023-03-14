package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testListKey = "test_list"

func TestRedisConnection(t *testing.T) {
	r := newTestRedis(t)
	// Creamos una lista y un valor de pruebas
	newValue := "test"
	rp := r.client.RPush(context.TODO(), testListKey, newValue)
	res, err := rp.Result()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	// Eliminamos ese valor de prueba
	lp := r.client.LPop(context.TODO(), testListKey)
	ress, err := lp.Result()
	assert.NoError(t, err)
	assert.NotEmpty(t, ress)
}

func TestAddPingComplete(t *testing.T) {
	r := newTestRedis(t)
	r.AddPingComplete(context.TODO(), false)
	clearRedis(t, r)
}

// Crea, revisa y devuelve una instancia de Redis.
func newTestRedis(t *testing.T) *Redis {
	// Creamos instancia de redis
	r, err := NewRedis()
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	return r
}

func clearRedis(t *testing.T, r *Redis) {
	f := r.client.FlushAll(context.TODO())
	res, err := f.Result()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
