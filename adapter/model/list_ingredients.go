package model

type ListIngredientsResponse struct {
	Ingredients []*Ingredient `json:"ingredients"`
}
