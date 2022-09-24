package domain

import "context"

type IngredientRepository interface {
	List(context.Context) ([]*Ingredient, error)
	Get(context.Context, IngredientID) (*Ingredient, error)
	Upsert(context.Context, *Ingredient) error
	Delete(context.Context, *Ingredient) error
}
