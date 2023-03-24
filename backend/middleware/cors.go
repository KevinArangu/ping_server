package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CorsMiddleware() echo.MiddlewareFunc {
	configs := middleware.CORSConfig{}
	return middleware.CORSWithConfig(configs)
}
