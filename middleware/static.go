package middleware

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StaticMiddleware() echo.MiddlewareFunc {
	configs := middleware.StaticConfig{
		Root:    "front",
		Browse:  true,
		HTML5:   true,
		Skipper: StaticSkipper,
	}
	return middleware.StaticWithConfig(configs)
}

func StaticSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/stats") || strings.HasPrefix(c.Path(), "/ping")
}
