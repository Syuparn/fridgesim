package usecase

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/syuparn/fridgesim/domain"
)

type CreateIngredientInputPort interface {
	Handle(context.Context, *CreateIngredientInputData) (*CreateIngredientOutputData, error)
}

type CreateIngredientInputData struct {
	Kind   string
	Amount float64
}

type CreateIngredientOutputData struct {
	Ingredient *domain.Ingredient
}

type createIngredientInteractor struct {
	ingredientFactory    domain.IngredientFactory
	ingredientRepository domain.IngredientRepository
}

func NewCreateIngredientInputPort(
	ingredientFactory domain.IngredientFactory,
	ingredientRepository domain.IngredientRepository,
) CreateIngredientInputPort {
	return &createIngredientInteractor{
		ingredientFactory:    ingredientFactory,
		ingredientRepository: ingredientRepository,
	}
}

var _ CreateIngredientInputPort = new(createIngredientInteractor)

func (i *createIngredientInteractor) Handle(
	ctx context.Context,
	in *CreateIngredientInputData,
) (*CreateIngredientOutputData, error) {
	ingredient := i.ingredientFactory.New(domain.IngredientKind(in.Kind), domain.IngredientAmount(in.Amount))

	err := i.ingredientRepository.Upsert(ctx, ingredient)
	if err != nil {
		return nil, xerrors.Errorf("failed to save ingredients: %w", err)
	}

	return &CreateIngredientOutputData{
		Ingredient: ingredient,
	}, nil
}
