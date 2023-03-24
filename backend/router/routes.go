package router

import "github.com/labstack/echo"

func CreateRoutes(e *echo.Echo) {
	router := e.Group("")
	router.GET("/ping", pingHandler)
	router.GET("/stats", statsHandler)
}
