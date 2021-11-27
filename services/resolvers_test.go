package services_test

import (
	"checkout-system-gql/mocks"
	"checkout-system-gql/services"
	"checkout-system-gql/structs"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

type GetItemsResolversTestCase struct {
	caseName      string
	params        graphql.ResolveParams
	expectedRes   []structs.ItemsList
	expectedError error
}

type DeleteItemsResolversTestCase struct {
	caseName      string
	params        graphql.ResolveParams
	itemList      structs.ItemsList
	expectedError error
}

func TestGetItemResolver(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCheckoutRepo := mocks.NewMockICheckoutRepositories(mockCtrl)

	testcases := []GetItemsResolversTestCase{
		{
			caseName: "Retrieve 2 Items Success",
			params: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId": "212",
				},
			},
			expectedRes: []structs.ItemsList{
				{
					BuyerID:  "212",
					Sku:      "234234",
					Name:     "MacBook Pro",
					Price:    "5959.35",
					Quantity: 2,
				},
				{
					BuyerID:  "212",
					Sku:      "ALE-7X0",
					Name:     "alexa Speaker",
					Price:    "49.99",
					Quantity: 5,
				},
			},
			expectedError: nil,
		},
		{
			caseName: "Retrieve 4 Items Success",
			params: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId": "6969",
				},
			},
			expectedRes: []structs.ItemsList{
				{
					BuyerID:  "6969",
					Sku:      "234234",
					Name:     "MacBook Pro",
					Price:    "$5959.35",
					Quantity: 1,
				},
				{
					BuyerID:  "6969",
					Sku:      "ALE7X0",
					Name:     "alexa Speaker",
					Price:    "$20.00",
					Quantity: 5,
				},
				{
					BuyerID:  "6969",
					Sku:      "ALE7X0",
					Name:     "Google Home",
					Price:    "$49.99",
					Quantity: 1,
				},
				{
					BuyerID:  "6969",
					Sku:      "234234",
					Name:     "Raspberry Pi B",
					Price:    "$30.00",
					Quantity: 2,
				},
			},
			expectedError: nil,
		},
		{
			caseName: "Empty Items",
			params: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId": "100",
				},
			},
			expectedRes:   []structs.ItemsList{},
			expectedError: nil,
		},
	}

	resolvers := services.NewCheckoutServices(mockCheckoutRepo)

	for in, tc := range testcases {
		mockCheckoutRepo.EXPECT().
			RetrieveItems(tc.params.Args["buyerId"]).
			Return(tc.expectedRes, nil).
			AnyTimes()

		res, err := resolvers.GetItemsResolver(tc.params)
		assert.Equal(t, tc.expectedError, err)
		assert.Equal(t, tc.expectedRes, res)

		fmt.Println("Resolvers Test | Retrieve Items on Testcase:", in+1, " test case name:", tc.caseName)
	}
}

func TestAddItemsResolver(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCheckoutRepo := mocks.NewMockICheckoutRepositories(mockCtrl)

	testCase := []DeleteItemsResolversTestCase{
		{
			caseName: "Stored Items Success",
			params: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId":  "212",
					"itemType": 4,
					"quantity": 2,
				},
			},

			itemList: structs.ItemsList{
				BuyerID:  "212",
				Sku:      "234234",
				Name:     "Raspberry Pi B",
				Price:    "30.00",
				Quantity: 2,
			},
			expectedError: nil,
		},
		{
			caseName: "Missing Mutation Arguments : itemType",
			params: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId":  "100",
					"quantity": 1,
				},
			},
			itemList: structs.ItemsList{
				Sku: "43N23P",
			},
			expectedError: errors.New("missing argument itemType"),
		},
		{
			caseName: "Missing Mutation Arguments : quantity",
			params: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId":  "100",
					"itemType": 1,
				},
			},
			itemList: structs.ItemsList{
				Sku: "43N23P",
			},
			expectedError: errors.New("missing argument quantity"),
		},
		{
			caseName: "argument quantity is not a string value",
			params: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId":  "100",
					"itemType": 1,
					"quantity": "2",
				},
			},
			itemList:      structs.ItemsList{},
			expectedError: errors.New("quantity is not a string value"),
		},
		{
			caseName: "argument itemType is not a string value",
			params: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId":  "100",
					"itemType": "1",
					"quantity": 2,
				},
			},
			itemList:      structs.ItemsList{},
			expectedError: errors.New("itemType is not a string value"),
		},
	}

	resolvers := services.NewCheckoutServices(mockCheckoutRepo)

	for in, tc := range testCase {
		cacheKey := fmt.Sprintf("buyer_%s:%s", tc.params.Args["buyerId"], tc.itemList.Sku)

		mockCheckoutRepo.EXPECT().
			StoreListItems(cacheKey, tc.itemList).
			Return(tc.expectedError).
			AnyTimes()

		_, err := resolvers.AddItemsResolver(tc.params)
		assert.Equal(t, tc.expectedError, err)

		fmt.Println("Resolvers Test | Delete Items on Testcase:", in+1, " test case name:", tc.caseName)
	}
}
