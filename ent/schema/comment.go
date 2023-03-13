package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Mixin for Comment.
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		CreatedAt{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("content").
			MinLen(64).
			MaxLen(512).
			Annotations(entgql.OrderField("content")).
			Comment("The comment's content."),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}

// Annotations of the Comment.
func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
