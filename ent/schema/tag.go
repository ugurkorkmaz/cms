package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Mixin for Tag.
func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		CreatedAt{},
		UpdatedAt{},
	}
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			MinLen(8).
			MaxLen(216).
			Annotations(entgql.OrderField("NAME")).
			Comment("The tag's name."),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}

// Annotations of the Tag.
func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
