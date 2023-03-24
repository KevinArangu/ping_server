package router

import (
	"context"

	"github.com/KevinArangu/stats_server/model"
	"github.com/KevinArangu/stats_server/service"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func pingHandler(c echo.Context) error {
	return c.JSON(200, model.Response{OK: true, Message: "Is OK!"})
}

func statsHandler(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r, err := service.NewRedis()
	if err != nil {
		return c.JSON(500, model.Response{OK: false, Message: err.Error()})
	}
	val, err := r.GetStats(ctx)
	if err != nil {
		return c.JSON(500, model.Response{OK: false, Message: err.Error()})
	}
	return c.JSON(200, val)
}
