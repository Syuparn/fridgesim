package infrastructure

import (
	"context"

	"golang.org/x/xerrors"

	_ "github.com/lib/pq"
	"github.com/samber/lo"

	"github.com/syuparn/fridgesim/domain"
	"github.com/syuparn/fridgesim/ent"
)

type ingredientRepository struct {
	client *ent.Client
}

var _ domain.IngredientRepository = new(ingredientRepository)

func NewIngredientRepository(client *ent.Client) (domain.IngredientRepository, error) {
	return &ingredientRepository{
		client: client,
	}, nil
}

func (r *ingredientRepository) List(ctx context.Context) ([]*domain.Ingredient, error) {
	ingredients, err := r.client.Ingredient.
		Query().
		All(ctx)

	if err != nil {
		return nil, xerrors.Errorf("failed to list ingredients: %w", err)
	}

	return lo.Map(ingredients, func(i *ent.Ingredient, _ int) *domain.Ingredient {
		return &domain.Ingredient{
			ID:     domain.IngredientID(i.ID),
			Kind:   domain.IngredientKind(i.Kind),
			Amount: domain.IngredientAmount(i.Amount),
		}
	}), nil
}

func (r *ingredientRepository) Upsert(ctx context.Context, ingredient *domain.Ingredient) error {
	// use OnConflict for upsert
	err := r.client.Ingredient.Create().
		SetID(string(ingredient.ID)).
		SetKind(string(ingredient.Kind)).
		SetAmount(float64(ingredient.Amount)).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		return xerrors.Errorf("failed to update ingredient: %w", err)
	}

	return nil
}
