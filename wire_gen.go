// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/labstack/echo/v4"
	"go-echo-contract-first-sample/internal/petstore/web"
	"go-echo-contract-first-sample/internal/server/config"
	"go-echo-contract-first-sample/internal/server/restful"
)

// Injectors from wire.go:

func Initialize() *AppContext {
	configConfig := config.Provide()
	echo := restful.Provide(configConfig)
	serverInterface := web.ProvidePetstoreRouter(configConfig, echo)
	appContext := &AppContext{
		Config:          configConfig,
		Echo:            echo,
		ServerInterface: serverInterface,
	}
	return appContext
}

// wire.go:

type AppContext struct {
	*config.Config
	*echo.Echo
	web.ServerInterface
}
