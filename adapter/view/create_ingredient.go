package view

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syuparn/fridgesim/adapter/model"
	"github.com/syuparn/fridgesim/usecase"
)

func CreateIngredient(c echo.Context, out *usecase.CreateIngredientOutputData) error {
	res := &model.Ingredient{
		ID:     string(out.Ingredient.ID),
		Kind:   string(out.Ingredient.Kind),
		Amount: float64(out.Ingredient.Amount),
	}

	return c.JSON(http.StatusCreated, res)
}
