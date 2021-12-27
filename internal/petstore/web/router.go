//go:generate oapi-codegen --config ./oapi-codegen-config.yaml ./spec.yaml

package web

import (
	_ "embed"
	"go-echo-contract-first-sample/internal/server/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed spec.yaml
var spec []byte

var _ ServerInterface = &PetstoreRouter{}

type PetstoreRouter struct {
}

func ProvidePetstoreRouter(cfg *config.Config, e *echo.Echo) ServerInterface {
	handler := &PetstoreRouter{}

	RegisterHandlersWithBaseURL(e, handler, cfg.ApiBaseUri)

	e.GET("/q/swagger-ui", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/swagger-ui/index.html?spec=/q/spec.yaml")
	})
	e.GET("/q/spec.yaml", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "text/x-yaml", spec)
	})

	return handler
}

func (p PetstoreRouter) FindPets(ctx echo.Context, params FindPetsParams) error {
	ctx.Error(echo.NewHTTPError(http.StatusNotImplemented))
	return nil
}

func (p PetstoreRouter) AddPet(ctx echo.Context) error {
	ctx.Error(echo.NewHTTPError(http.StatusNotImplemented))
	return nil
}

func (p PetstoreRouter) DeletePet(ctx echo.Context, id int64) error {
	ctx.Error(echo.NewHTTPError(http.StatusNotImplemented))
	return nil
}

func (p PetstoreRouter) FindPetById(ctx echo.Context, id int64) error {
	ctx.Error(echo.NewHTTPError(http.StatusNotImplemented))
	return nil
}

