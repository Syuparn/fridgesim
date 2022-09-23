package main

import (
	"fmt"

	"github.com/syuparn/fridgesim/adapter"
	"github.com/syuparn/fridgesim/pkg/config"
)

func main() {
	e := adapter.NewServer()
	cfg, _ := config.New()

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
