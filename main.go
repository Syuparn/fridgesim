package main

import (
	"github.com/syuparn/fridgesim/adapter"
)

func main() {
	e := adapter.NewServer()

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
