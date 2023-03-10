package schema

import "entgo.io/ent"

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Mixin for Article.
func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		UpdatedAt{},
		CreatedAt{},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return nil
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
