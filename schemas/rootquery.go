package schemas

import (
	"checkout-system-gql/services"

	"github.com/graphql-go/graphql"
)

func generateRootQuery(cs services.ICheckoutServices) *graphql.Object {
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name: "Items",
		Fields: graphql.Fields{
			"buyerId": &graphql.Field{
				Type:        graphql.String,
				Description: "Buyer Id",
			},
			"quantity": &graphql.Field{
				Type:        graphql.String,
				Description: "Item Quantity",
			},
			"sku": &graphql.Field{
				Type:        graphql.String,
				Description: "Item SKU",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Item Name",
			},
			"price": &graphql.Field{
				Type:        graphql.String,
				Description: "Item Price",
			},
		},
	})

	field := graphql.Fields{
		"buyer": &graphql.Field{
			Type: graphql.NewList(obj),
			Args: graphql.FieldConfigArgument{
				"buyerId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve:     cs.GetItemsResolver,
			Description: "Query all Items from Cart",
		},
	}

	rootCfg := graphql.ObjectConfig{
		Name:   "BuyerQuery",
		Fields: field,
	}

	rootQuery := graphql.NewObject(rootCfg)

	return rootQuery
}
