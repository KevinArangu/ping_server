package main

import (
	"github.com/KevinArangu/stats_server/config"
	"github.com/KevinArangu/stats_server/router"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	if err := config.ReadConfig(); err != nil {
		e.Logger.Fatal("Error read envs")
	}
	router.CreateRoutes(e)
	err := e.Start(config.ServerAddress())
	if err != nil {
		e.Logger.Fatal("Error starting server")
	}
}
