package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Size holds the schema definition for the Size entity.
type Size struct {
	ent.Schema
}

// Fields of the Size.
func (Size) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Immutable().
			Unique(),
		field.Int("size").
			Unique().
			Positive(),
	}
}

// Edges of the Size.
func (Size) Edges() []ent.Edge {
	return nil
}
