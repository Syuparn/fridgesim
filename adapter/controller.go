package adapter

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/syuparn/fridgesim/adapter/model"
	"github.com/syuparn/fridgesim/adapter/view"
	"github.com/syuparn/fridgesim/usecase"
)

type Controller struct {
	createIngredientInputPort usecase.CreateIngredientInputPort
	listIngredientsInputPort  usecase.ListIngredientsInputPort
	deleteIngredientInputPort usecase.DeleteIngredientInputPort
}

func NewController(
	createIngredientInputPort usecase.CreateIngredientInputPort,
	listIngredientsInputPort usecase.ListIngredientsInputPort,
	deleteIngredientInputPort usecase.DeleteIngredientInputPort,
) *Controller {
	return &Controller{
		createIngredientInputPort: createIngredientInputPort,
		listIngredientsInputPort:  listIngredientsInputPort,
		deleteIngredientInputPort: deleteIngredientInputPort,
	}
}

func (ctr *Controller) CreateIngredient(c echo.Context) error {
	var req model.CreateIngredientRequest
	if err := c.Bind(&req); err != nil {
		// TODO: error handling
		return view.Error(c, err)
	}

	in := &usecase.CreateIngredientInputData{
		Kind:   req.Kind,
		Amount: req.Amount,
	}
	out, err := ctr.createIngredientInputPort.Handle(c.Request().Context(), in)
	if err != nil {
		log.Warnf("%+v", err)
		return view.Error(c, err)
	}

	return view.CreateIngredient(c, out)
}

func (ctr *Controller) ListIngredients(c echo.Context) error {
	in := &usecase.ListIngredientsInputData{}
	out, err := ctr.listIngredientsInputPort.Handle(c.Request().Context(), in)
	if err != nil {
		log.Warnf("%+v", err)
		return view.Error(c, err)
	}

	return view.ListIngredients(c, out)
}

func (ctr *Controller) DeleteIngredient(c echo.Context) error {
	in := &usecase.DeleteIngredientInputData{
		ID: c.Param("ingredient"),
	}
	out, err := ctr.deleteIngredientInputPort.Handle(c.Request().Context(), in)
	if err != nil {
		log.Warnf("%+v", err)
		return view.Error(c, err)
	}

	return view.DeleteIngredient(c, out)
}
