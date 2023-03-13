package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Mixin for Article.
func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("title").
			MinLen(8).
			MaxLen(1024).
			NotEmpty().
			Annotations(entgql.OrderField("TITLE")).
			Comment("The article's title."),

		field.
			String("slug").
			Unique().
			MinLen(8).
			MaxLen(1024).
			NotEmpty().
			Annotations(entgql.OrderField("SLUG")).
			Comment("The article's slug."),

		field.
			String("description").
			Optional().
			MinLen(64).
			MaxLen(512).
			Annotations(entgql.OrderField("DESCRIPTION")).
			Comment("The article's description."),

		field.
			Text("content").
			MinLen(64).
			NotEmpty().
			Annotations(entgql.OrderField("CONTENT")).
			Comment("The article's content."),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("categories", Category.Type).
			Comment("The article's category"),
	}
}

// Annotations of the Article.
func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
