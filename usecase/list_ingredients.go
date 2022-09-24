package usecase

import (
	"context"

	"github.com/syuparn/fridgesim/domain"
	"golang.org/x/xerrors"
)

type ListIngredientsInputPort interface {
	Handle(context.Context, *ListIngredientsInputData) (*ListIngredientsOutputData, error)
}

type ListIngredientsInputData struct{}

type ListIngredientsOutputData struct {
	Ingredients []*domain.Ingredient
}

type listIngredientInteractor struct {
	ingredientRepository domain.IngredientRepository
}

func NewListIngredientsInputPort(
	ingredientRepository domain.IngredientRepository,
) ListIngredientsInputPort {
	return &listIngredientInteractor{
		ingredientRepository: ingredientRepository,
	}
}

var _ ListIngredientsInputPort = new(listIngredientInteractor)

func (i *listIngredientInteractor) Handle(
	ctx context.Context,
	_ *ListIngredientsInputData,
) (*ListIngredientsOutputData, error) {
	ingredients, err := i.ingredientRepository.List(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to get ingredients: %w", err)
	}

	return &ListIngredientsOutputData{
		Ingredients: ingredients,
	}, nil
}
