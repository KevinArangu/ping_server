package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/KevinArangu/stats_server/config"
	"github.com/KevinArangu/stats_server/interactor"
	"github.com/KevinArangu/stats_server/router"
	"github.com/KevinArangu/stats_server/service"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	log := logrus.New()
	e := echo.New()
	r := service.NewRedis()
	p := service.NewPinger()

	if err := config.ReadConfig(); err != nil {
		log.WithError(err).Fatal("ENVS error")
	}

	router.CreateRoutes(e)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		e.Shutdown(ctx)
	}()

	go interactor.StartPing(ctx, e, r, p)

	err := e.Start(config.ServerAddress())
	if err != nil {
		log.WithError(err).Fatal("Server error")
	}
}
