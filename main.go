package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/KevinArangu/stats_server/config"
	"github.com/KevinArangu/stats_server/interactor"
	"github.com/KevinArangu/stats_server/middleware"
	"github.com/KevinArangu/stats_server/router"
	"github.com/KevinArangu/stats_server/service"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	ctx := context.Background()
	if err := config.ReadConfig(); err != nil {
		log.WithError(err).Fatal("ENVS error")
	}
	e := echo.New()
	p := service.NewPinger()
	r, err := service.NewRedis()
	if err != nil {
		log.WithError(err).Fatal("Failed to create a Redis client")
	}

	e.Use(middleware.StaticMiddleware())

	router.CreateRoutes(e)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		e.Shutdown(ctx)
	}()

	go interactor.StartPing(ctx, e, r, p)

	err = e.Start(config.ServerAddress())
	if err != nil {
		log.WithError(err).Fatal("Server error")
	}
}
