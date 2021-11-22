package schemas

import (
	"github.com/graphql-go/graphql"
)

func getAddItemsOutput() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "ResponseAddItem",
		Fields: graphql.Fields{
			"items": &graphql.Field{
				Type: graphql.String,
			},
			"total": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func getCheckoutItemsOutput() *graphql.List {
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name: "ResponseCheckout",
		Fields: graphql.Fields{
			"desc": &graphql.Field{
				Type: graphql.String,
			},
			"total": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	return graphql.NewList(obj)
}
