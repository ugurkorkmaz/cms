package resolver

import (
	"app/ent"
	"app/gen"

	"github.com/99designs/gqlgen/graphql"
)

type Resolver struct {
	client *ent.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return gen.NewExecutableSchema(gen.Config{
		Resolvers: &Resolver{
			client: client,
		},
	})
}
