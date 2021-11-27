package commons_test

import (
	"checkout-system-gql/commons"
	"checkout-system-gql/structs"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CaseAddItemsResponse struct {
	caseName string
	params   structs.ItemsList
	Response structs.ResponseAddItems
}

type CaseGetPromotionResponse struct {
	caseName string
	params   []structs.ItemsList
	Response []structs.ResponseCheckout
}

func TestAddToCartResponse(t *testing.T) {
	testCase := []CaseAddItemsResponse{
		{
			caseName: "Add Item MacBook Pro",
			params: structs.ItemsList{
				Name:     "MacBook Pro",
				Price:    "5399.99",
				Quantity: 1,
			},
			Response: structs.ResponseAddItems{
				Items: "MacBook Pro",
				Total: 1,
				Price: "$5399.99",
			},
		},
		{
			caseName: "Add Item Alexa Speaker",
			params: structs.ItemsList{
				Name:     "Alexa Speaker",
				Price:    "109.50",
				Quantity: 2,
			},
			Response: structs.ResponseAddItems{
				Items: "Alexa Speaker",
				Total: 2,
				Price: "$109.50",
			},
		},
		{
			caseName: "Add Item Google Home",
			params: structs.ItemsList{
				Name:     "Google Home",
				Price:    "49.99",
				Quantity: 5,
			},
			Response: structs.ResponseAddItems{
				Items: "Google Home",
				Total: 5,
				Price: "$49.99",
			},
		},
		{
			caseName: "Add Item Raspberry Pi B",
			params: structs.ItemsList{
				Name:     "Raspberry Pi B",
				Price:    "30.00",
				Quantity: 3,
			},
			Response: structs.ResponseAddItems{
				Items: "Raspberry Pi B",
				Total: 3,
				Price: "$30.00",
			},
		},
	}

	for in, tc := range testCase {
		res := commons.AddToCartResponse(tc.params)

		assert.Equal(t, res, tc.Response)
		fmt.Println("Test Response Add Item | Testcase:", in+1, " test case name:", tc.caseName)
	}
}

func TestGetpromotionResponse(t *testing.T) {
	testCase := []CaseGetPromotionResponse{
		{
			caseName: "Add Item MacBook Pro",
			params: []structs.ItemsList{
				{
					Name:     "MacBook Pro",
					Sku:      "43N23P",
					Price:    "5499.99",
					Quantity: 1,
				},
				{
					Name:     "Google Home",
					Sku:      "120P90",
					Price:    "49.99",
					Quantity: 3,
				},
				{
					Name:     "Alexa Speaker",
					Sku:      "A304SD",
					Price:    "109.50",
					Quantity: 5,
				},
			},
			Response: []structs.ResponseCheckout{
				{
					Desc:  "Scanned Items: MacBook Pro, Raspberry Pi B",
					Total: "$5499.99",
				},
				{
					Desc:  "Scanned Items: Google Home, Google Home, Google Home",
					Total: "$99.98",
				},
				{
					Desc:  "Scanned Items: Alexa Speaker, Alexa Speaker, Alexa Speaker, Alexa Speaker, Alexa Speaker",
					Total: "$492.75",
				},
			},
		},
	}

	for in, tc := range testCase {
		res := commons.GetPromotionResponse(tc.params)

		assert.Equal(t, res, tc.Response)
		fmt.Println("Test Get Promotion Response| Testcase:", in+1, " test case name:", tc.caseName)
	}
}
