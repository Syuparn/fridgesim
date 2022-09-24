package domain

import "context"

type IngredientRepository interface {
	List(context.Context) ([]*Ingredient, error)
}