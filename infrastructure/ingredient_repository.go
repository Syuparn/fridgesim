package infrastructure

import (
	"context"

	"golang.org/x/xerrors"

	_ "github.com/lib/pq"
	"github.com/samber/lo"

	"github.com/syuparn/fridgesim/domain"
	"github.com/syuparn/fridgesim/ent"
	ingredientschema "github.com/syuparn/fridgesim/ent/ingredient"
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

func (r *ingredientRepository) Get(ctx context.Context, id domain.IngredientID) (*domain.Ingredient, error) {
	ingredient, err := r.client.Ingredient.
		Query().
		Where(ingredientschema.ID(string(id))).
		Only(ctx)

	if err != nil {
		return nil, xerrors.Errorf("failed to get ingredient: %w", err)
	}

	return &domain.Ingredient{
		ID:     domain.IngredientID(ingredient.ID),
		Kind:   domain.IngredientKind(ingredient.Kind),
		Amount: domain.IngredientAmount(ingredient.Amount),
	}, nil
}

func (r *ingredientRepository) Upsert(ctx context.Context, ingredient *domain.Ingredient) error {
	// use OnConflict for upsert
	err := r.client.Ingredient.Create().
		SetID(string(ingredient.ID)).
		SetKind(string(ingredient.Kind)).
		SetAmount(float64(ingredient.Amount)).
		OnConflictColumns(ingredientschema.FieldID).
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		return xerrors.Errorf("failed to update ingredient: %w", err)
	}

	return nil
}

func (r *ingredientRepository) Delete(ctx context.Context, ingredient *domain.Ingredient) error {

	_, err := r.client.Ingredient.Delete().Where(ingredientschema.ID(string(ingredient.ID))).Exec(ctx)
	if err != nil {
		return xerrors.Errorf("failed to delete ingredient: %w", err)
	}

	return nil
}
