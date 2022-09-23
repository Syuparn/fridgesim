package di

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/do"

	"github.com/syuparn/fridgesim/adapter"
	"github.com/syuparn/fridgesim/pkg/config"
)

func New() *do.Injector {
	injector := do.New()

	do.Provide(injector, newConfig)
	do.Provide(injector, newServer)

	return injector
}

func newConfig(i *do.Injector) (*config.Specification, error) {
	return config.New()
}

func newServer(i *do.Injector) (*echo.Echo, error) {
	return adapter.NewServer(), nil
}
