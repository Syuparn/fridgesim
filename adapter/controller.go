package adapter

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/syuparn/fridgesim/adapter/view"
	"github.com/syuparn/fridgesim/usecase"
)

type Controller struct {
	listIngredientsInputPort usecase.ListIngredientsInputPort
}

func NewController(
	listIngredientsInputPort usecase.ListIngredientsInputPort,
) *Controller {
	return &Controller{
		listIngredientsInputPort: listIngredientsInputPort,
	}
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
