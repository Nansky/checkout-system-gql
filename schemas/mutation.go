package schemas

import (
	"checkout-system-gql/services"

	"github.com/graphql-go/graphql"
)

func generateNewMutation(s services.ICheckoutServices) *graphql.Object {
	mutationFields := graphql.Fields{
		// Create a mutations
		"addItems": &graphql.Field{
			Type:        getAddItemsOutput(),
			Args:        getAddItemArgs(),
			Resolve:     s.AddItemsResolver,
			Description: "Add Items to cart",
		},
		"checkout": &graphql.Field{
			Type:        getCheckoutItemsOutput(),
			Args:        getCheckoutItemsArgs(),
			Resolve:     s.CheckoutItemsResolver,
			Description: "Checkout Items from cart",
		},
	}

	mutationConfig := graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields}
	return graphql.NewObject(mutationConfig)
}
