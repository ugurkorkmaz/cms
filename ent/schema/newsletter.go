package schema

import "entgo.io/ent"

// Newsletter holds the schema definition for the Newsletter entity.
type Newsletter struct {
	ent.Schema
}

// Mixin for Newsletter.
func (Newsletter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		UpdatedAt{},
		CreatedAt{},
	}
}

// Fields of the Newsletter.
func (Newsletter) Fields() []ent.Field {
	return nil
}

// Edges of the Newsletter.
func (Newsletter) Edges() []ent.Edge {
	return nil
}
