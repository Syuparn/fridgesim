package model

type CreateIngredientRequest struct {
	Kind   string  `json:"kind"`
	Amount float64 `json:"amount"`
}
