package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(ctrl *Controller) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/ingredients", ctrl.CreateIngredient)
	e.GET("/ingredients", ctrl.ListIngredients)
	e.DELETE("/ingredients/:ingredient", ctrl.DeleteIngredient)

	return e
}
