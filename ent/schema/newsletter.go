package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

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
	return []ent.Field{
		field.
			String("message").
			MinLen(64).
			MaxLen(1024).
			Annotations(entgql.OrderField("MESSAGE")).
			Comment("The newsletter's message."),
	}
}

// Edges of the Newsletter.
func (Newsletter) Edges() []ent.Edge {
	return nil
}

// Annotations of the Newsletter.
func (Newsletter) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
