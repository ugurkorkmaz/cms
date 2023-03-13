package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

// Mixin for Metadata.
func (Metadata) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		UpdatedAt{},
		CreatedAt{},
	}
}

// Fields of the Metadata.
func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("title").
			MinLen(8).
			MaxLen(216).
			Annotations(entgql.OrderField("TITLE")).
			Comment("The website's title."),
	}
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return nil
}

// Annotations of the Metadata.
func (Metadata) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField("metadata"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
