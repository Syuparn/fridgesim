package usecase

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/syuparn/fridgesim/domain"
)

type DeleteIngredientInputPort interface {
	Handle(context.Context, *DeleteIngredientInputData) (*DeleteIngredientOutputData, error)
}

type DeleteIngredientInputData struct {
	ID string
}

type DeleteIngredientOutputData struct{}

type deleteIngredientInteractor struct {
	ingredientRepository domain.IngredientRepository
}

func NewDeleteIngredientInputPort(
	ingredientRepository domain.IngredientRepository,
) DeleteIngredientInputPort {
	return &deleteIngredientInteractor{
		ingredientRepository: ingredientRepository,
	}
}

var _ DeleteIngredientInputPort = new(deleteIngredientInteractor)

func (i *deleteIngredientInteractor) Handle(
	ctx context.Context,
	in *DeleteIngredientInputData,
) (*DeleteIngredientOutputData, error) {
	ingredient, err := i.ingredientRepository.Get(ctx, domain.IngredientID(in.ID))
	if err != nil {
		return nil, xerrors.Errorf("failed to get ingredient: %w", err)
	}

	err = i.ingredientRepository.Delete(ctx, ingredient)
	if err != nil {
		return nil, xerrors.Errorf("failed to get ingredients: %w", err)
	}

	return &DeleteIngredientOutputData{}, nil
}
