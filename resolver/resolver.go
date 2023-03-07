package resolver

import (
	"app/gen"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func NewSchema() graphql.ExecutableSchema {
	return gen.NewExecutableSchema(gen.Config{
		Resolvers: &Resolver{},
	})
}
