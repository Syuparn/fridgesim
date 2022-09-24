package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Ingredient holds the schema definition for the Ingredient entity.
type Ingredient struct {
	ent.Schema
}

// Fields of the Ingredient.
func (Ingredient) Fields() []ent.Field {
	return []ent.Field{
		field.String("kind"),
		field.Float("amount"),
	}
}

// Edges of the Ingredient.
func (Ingredient) Edges() []ent.Edge {
	return nil
}
