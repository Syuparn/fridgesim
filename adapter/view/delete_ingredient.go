package view

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syuparn/fridgesim/usecase"
)

func DeleteIngredient(c echo.Context, out *usecase.DeleteIngredientOutputData) error {
	return c.JSON(http.StatusNoContent, nil)
}
