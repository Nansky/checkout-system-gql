package cmd

import (
	"checkout-system-gql/cache"
	"checkout-system-gql/repositories"
	"checkout-system-gql/schemas"
	"checkout-system-gql/services"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

// StartServer will trigger the server with a Playground
func StartServer() {

	redisCache := cache.NewCache()
	repositories := repositories.NewCheckoutRepositories(redisCache)
	checkoutServices := services.NewCheckoutServices(repositories)

	schemaObj := schemas.NewCheckoutSchema(checkoutServices)
	schema, err := schemaObj.GenerateNewSchemas()

	if err != nil {
		fmt.Printf("ERROR %v", err)
		panic(err)
	}

	// Create a new HTTP handler
	h := handler.New(&handler.Config{
		Schema:     schema,
		Pretty:     true,
		GraphiQL:   true, // GraphiQL used for testing Queries
		Playground: true,
	})

	http.Handle("/add_checkout", h)

	log.Print("Server Running at port 6969")
	log.Fatal(http.ListenAndServe(":6969", nil))
}
