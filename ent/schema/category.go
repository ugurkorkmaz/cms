package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Mixin for Category.
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		CreatedAt{},
		UpdatedAt{},
	}
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			MinLen(8).
			MaxLen(216).
			Annotations(entgql.OrderField("NAME")).
			Comment("The tag's name."),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("article", Article.Type).
			Ref("categories").
			Comment("The article's categories."),
	}
}

// Annotations of the Category.
func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
