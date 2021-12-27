package restful

import (
	"go-echo-contract-first-sample/internal/server/config"
	"go-echo-contract-first-sample/internal/server/restful/swaggerui"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Provide provides a preconfigured echo.Echo app to use.
func Provide(cfg *config.Config) *echo.Echo {
	e := echo.New()

	swaggerui.Mount(e)

	e.GET("/q/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"status": "UP"})
	})

	return e
}
