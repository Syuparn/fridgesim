package domain

// aggregate roots
type Ingredient struct {
	ID     IngredientID
	Kind   IngredientKind
	Amount IngredientAmount
}

// Value objects
type IngredientID string
type IngredientKind string
type IngredientAmount float64

// IngredientEpsilon defines threshold of ingredient amount.
// If remained amount of ingredient is less than this value, ingredient is treated as consumed (or deleted).
// NOTE: without IngredientEpsilon, the system will suffer from rounding error!
const IngredientEpsilon = 0.05
