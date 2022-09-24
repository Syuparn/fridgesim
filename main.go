package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
	log "github.com/sirupsen/logrus"

	"github.com/syuparn/fridgesim/ent"
	"github.com/syuparn/fridgesim/pkg/config"
	"github.com/syuparn/fridgesim/pkg/di"
)

func main() {
	injector := di.New()

	cfg := do.MustInvoke[*config.Specification](injector)
	e := do.MustInvoke[*echo.Echo](injector)
	entClient := do.MustInvoke[*ent.Client](injector)

	// auto migration
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
