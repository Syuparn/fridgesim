package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"

	"github.com/syuparn/fridgesim/pkg/config"
	"github.com/syuparn/fridgesim/pkg/di"
)

func main() {
	injector := di.New()

	cfg := do.MustInvoke[*config.Specification](injector)
	e := do.MustInvoke[*echo.Echo](injector)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
