package schema

import "entgo.io/ent"

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Mixin for Comment.
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		UpdatedAt{},
		CreatedAt{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return nil
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
