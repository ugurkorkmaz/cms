package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Gallery holds the schema definition for the Gallery entity.
type Gallery struct {
	ent.Schema
}

// Mixin for Gallery.
func (Gallery) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		CreatedAt{},
		UpdatedAt{},
	}
}

// Fields of the Gallery.
func (Gallery) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			MinLen(8).
			MaxLen(216).
			Annotations(entgql.OrderField("NAME")).
			Comment("The tag's name."),
	}
}

// Edges of the Gallery.
func (Gallery) Edges() []ent.Edge {
	return nil
}

// Annotations of the Gallery.
func (Gallery) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
