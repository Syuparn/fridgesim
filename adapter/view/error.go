package view

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syuparn/fridgesim/adapter/model"
)

func Error(c echo.Context, err error) error {
	// TODO: impl
	res := &model.Error{
		Type:    "Error",
		Message: err.Error(),
	}

	return c.JSON(http.StatusInternalServerError, res)
}
