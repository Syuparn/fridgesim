package model

type Ingredient struct {
	ID     string  `json:"id"`
	Kind   string  `json:"kind"`
	Amount float64 `json:"amount"`
}
