package schema

import "entgo.io/ent"

// Meta holds the schema definition for the Meta entity.
type Meta struct {
	ent.Schema
}

// Mixin for Meta.
func (Meta) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		UpdatedAt{},
		CreatedAt{},
	}
}

// Fields of the Meta.
func (Meta) Fields() []ent.Field {
	return nil
}

// Edges of the Meta.
func (Meta) Edges() []ent.Edge {
	return nil
}
