package adapter

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	return e
}

type helloResponse struct {
	Message string `json:"message"`
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, &helloResponse{Message: "Hello, world!"})
}
