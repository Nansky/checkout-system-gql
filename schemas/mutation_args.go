package schemas

import (
	"github.com/graphql-go/graphql"
)

func getAddItemArgs() graphql.FieldConfigArgument {
	args := graphql.FieldConfigArgument{
		"buyerId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"itemType": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"quantity": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}

	return args
}

func getCheckoutItemsArgs() graphql.FieldConfigArgument {
	args := graphql.FieldConfigArgument{
		"buyerId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	}

	return args

}
