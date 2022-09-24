package infrastructure

import (
	"context"
	"database/sql"

	"golang.org/x/xerrors"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/samber/lo"

	"github.com/syuparn/fridgesim/domain"
	"github.com/syuparn/fridgesim/ent"
)

type ingredientRepository struct {
	client *ent.Client
}

var _ domain.IngredientRepository = new(ingredientRepository)

func NewIngredientRepository(db *sql.DB) (domain.IngredientRepository, error) {

	drv := entsql.OpenDB("postgres", db)
	client := ent.NewClient(ent.Driver(drv))

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
