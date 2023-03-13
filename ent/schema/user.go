package schema

import (
	"fmt"
	"net/mail"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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
	return []ent.Field{
		field.
			String("name").
			NotEmpty().
			MinLen(3).
			MaxLen(64).
			Annotations(entgql.OrderField("NAME")).
			Comment("The user's name."),

		field.
			String("email").
			NotEmpty().
			Unique().
			Immutable().
			Annotations(entgql.OrderField("EMAIL")).
			Validate(func(s string) error {
				if _, err := mail.ParseAddress(s); err != nil {
					return fmt.Errorf("invalid email address: %w", err)
				}
				return nil
			}).
			Comment("The user's email."),

		field.
			String("password").
			NotEmpty().
			MinLen(8).
			MaxLen(264).
			Sensitive().
			Comment("The user's password."),

		field.
			Enum("role").
			Values("admin", "author", "user", "subscriber").
			Default("user").
			Annotations(entgql.OrderField("ROLE")).
			Comment("The user's role."),

		field.
			String("token").
			Optional().
			Sensitive().
			Annotations(entgql.OrderField("TOKEN")).
			Comment("The user's token."),

		field.
			Time("token_expired").
			Optional().
			Default(time.Now().Add(time.Hour * 24 * 7)).
			Annotations(entgql.OrderField("TOKEN_EXPIRED")).
			Comment("The user's token expiration time."),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Adding the annotations to the schema.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
