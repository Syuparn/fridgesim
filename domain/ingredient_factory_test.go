package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIngredientFactoryNew(t *testing.T) {
	mockIDGenerator := func() string {
		return "01GDPBP8XX4C0BPRCCBN52W070"
	}

	tests := []struct {
		name     string
		kind     IngredientKind
		amount   IngredientAmount
		expected *Ingredient
	}{
		{
			"generate new ingredient",
			"cabbage",
			2.0,
			&Ingredient{
				ID:     "01GDPBP8XX4C0BPRCCBN52W070",
				Kind:   "cabbage",
				Amount: 2.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory := NewIngredientFacotry(mockIDGenerator)

			actual := factory.New(tt.kind, tt.amount)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
