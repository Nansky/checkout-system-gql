package schemas

import (
	"checkout-system-gql/services"

	"github.com/graphql-go/graphql"
)

type ISchemas interface {
	GenerateNewSchemas() (*graphql.Schema, error)
}

type Schemas struct {
	checkoutService services.ICheckoutServices
}

func NewCheckoutSchema(cs services.ICheckoutServices) ISchemas {
	return &Schemas{
		checkoutService: cs,
	}
}

func (s *Schemas) GenerateNewSchemas() (*graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{
		Query:    generateRootQuery(s.checkoutService),
		Mutation: generateNewMutation(s.checkoutService),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
