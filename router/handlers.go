package router

import (
	"github.com/KevinArangu/stats_server/model"
	"github.com/labstack/echo"
)

func pingHandler(c echo.Context) error {
	return c.JSON(200, model.Response{OK: true, Message: "Is OK!"})
}
