package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// UUID is a mixin that adds a UUID field to an entity.
type Id struct {
	mixin.Schema
}

// Fields of the UUID mixin.
func (Id) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("ID"),
			).
			Default(uuid.New).
			Immutable().
			Comment("The unique identifier of the entity."),
	}
}

// UpdatedAt mixin for the updated_at field.
type UpdatedAt struct {
	mixin.Schema
}

// Fields of the UpdatedAt mixin.
func (UpdatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.
			Time("updated_at").
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			).
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("The time when the entity was updated."),
	}
}

// DeletedAt mixin for the deleted_at field.
type DeletedAt struct {
	mixin.Schema
}

// Fields of the DeletedAt mixin.
func (DeletedAt) Fields() []ent.Field {
	return []ent.Field{
		field.
			Time("deleted_at").
			Annotations(
				entgql.OrderField("DELETED_AT"),
			).
			Optional().
			Nillable().
			Comment("The time when the entity was deleted."),
	}
}

// CreatedAt mixin for the created_at field.
type CreatedAt struct {
	mixin.Schema
}

// Fields of the CreatedAt mixin.
func (CreatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.
			Time("created_at").
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			).
			Default(time.Now).
			Comment("The time when the entity was created."),
	}
}
