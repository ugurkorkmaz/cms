package schema

import (
	"entgo.io/ent"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin for User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Id{},
		UpdatedAt{},
		CreatedAt{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
