package view

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/syuparn/fridgesim/adapter/model"
	"github.com/syuparn/fridgesim/domain"
	"github.com/syuparn/fridgesim/usecase"
)

func ListIngredients(c echo.Context, out *usecase.ListIngredientsOutputData) error {
	res := &model.ListIngredientsResponse{
		Ingredients: lo.Map(out.Ingredients, func(i *domain.Ingredient, _ int) *model.Ingredient {
			return &model.Ingredient{
				ID:     string(i.ID),
				Kind:   string(i.Kind),
				Amount: float64(i.Amount),
			}
		}),
	}

	return c.JSON(http.StatusOK, res)
}
