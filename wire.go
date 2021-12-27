//go:build wireinject
// +build wireinject

//go:generate wire .

package main

import (
	"go-echo-contract-first-sample/internal/petstore/web"
	"go-echo-contract-first-sample/internal/server/config"
	"go-echo-contract-first-sample/internal/server/restful"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type AppContext struct {
	*config.Config
	*echo.Echo
	web.ServerInterface
}

func Initialize() *AppContext {
	panic(wire.Build(
		// Configuration
		config.Provide,

		// REST
		restful.Provide,
		web.ProvidePetstoreRouter,
		// AppContext
		wire.Struct(new(AppContext), "*"),
	))
}
