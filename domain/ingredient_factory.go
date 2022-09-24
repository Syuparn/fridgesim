package domain

import (
	"github.com/oklog/ulid/v2"
)

type IngredientFactory interface {
	New(IngredientKind, IngredientAmount) *Ingredient
}

type ingredientFactory struct {
	generateID func() string
}

var _ IngredientFactory = new(ingredientFactory)

func NewIngredientFacotry(idGenerator func() string) IngredientFactory {
	// use default ULID generator
	if idGenerator == nil {
		idGenerator = func() string {
			return ulid.Make().String()
		}
	}

	return &ingredientFactory{
		generateID: idGenerator,
	}
}

func (f *ingredientFactory) New(kind IngredientKind, amount IngredientAmount) *Ingredient {
	return &Ingredient{
		ID:     IngredientID(f.generateID()),
		Kind:   kind,
		Amount: amount,
	}
}
