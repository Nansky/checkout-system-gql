package schemas_test

import (
	"checkout-system-gql/mocks"
	"checkout-system-gql/schemas"
	"checkout-system-gql/structs"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

type SchemaMutationTest struct {
	casename         string
	params           graphql.ResolveParams
	expectedAddItems structs.ResponseAddItems
	expectedCheckout []structs.ItemsList
	expectedErr      error
}

type SchemaRootqueryTest struct {
	casename              string
	paramsQuery           graphql.ResolveParams
	expectedQueryResponse []structs.ItemsList
	expectedErr           error
}

func TestRootqueryGenerateNewSchema(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCtrl.Finish()

	testcase := []SchemaRootqueryTest{
		{
			casename: "Test Rootquery success retrieve items from added items",
			paramsQuery: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId": "212",
				},
			},
			expectedQueryResponse: []structs.ItemsList{
				{
					Name:     "Raspberry Pi B",
					Price:    "30.00",
					Quantity: 1,
				},
				{
					Name:     "Google Home",
					Price:    "49.99",
					Quantity: 2,
				},
				{
					Name:     "Alexa Speaker",
					Price:    "109.50",
					Quantity: 1,
				},
			},
			expectedErr: nil,
		},
		{
			casename: "Test Rootquery, empty item list",
			paramsQuery: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId": "213",
				},
			},
			expectedQueryResponse: []structs.ItemsList{},
			expectedErr:           nil,
		},
	}

	mockResolver := mocks.NewMockICheckoutServices(mockCtrl)

	for in, tc := range testcase {
		mockResolver.EXPECT().
			GetItemsResolver(tc.paramsQuery).
			Return(tc.expectedQueryResponse, tc.expectedErr)

		schemas := schemas.NewCheckoutSchema(mockResolver)
		_, err := schemas.GenerateNewSchemas()

		assert.Equal(t, tc.expectedErr, err)
		fmt.Println("Test Generate New Schema Rootquery| Testcase:", in+1, " test case name:", tc.casename)
	}
}

func TestMutationGenerateNewSchemas(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testcase := []SchemaMutationTest{
		{
			casename: "Test Mutation Schema",
			params: graphql.ResolveParams{
				Args: map[string]interface{}{
					"buyerId":  "212",
					"quantity": 1,
					"itemType": 3,
				},
			},
			expectedAddItems: structs.ResponseAddItems{
				Items: "Alexa Speaker",
				Total: 1,
				Price: "$109.50",
			},
			expectedCheckout: []structs.ItemsList{
				{
					BuyerID:  "212",
					Sku:      "A304SD",
					Name:     "Alexa Speaker",
					Price:    "$109.50",
					Quantity: 1,
				},
			},
			expectedErr: nil,
		},
	}

	mockSvc := mocks.NewMockICheckoutServices(mockCtrl)
	testSchema := schemas.NewCheckoutSchema(mockSvc)

	for in, tc := range testcase {

		mockSvc.EXPECT().
			AddItemsResolver(gomock.Any()).
			Return(tc.expectedAddItems, nil).
			AnyTimes()

		mockSvc.EXPECT().
			CheckoutItemsResolver(gomock.All()).
			Return(tc.expectedCheckout, nil).
			AnyTimes()

		_, err := testSchema.GenerateNewSchemas()
		assert.Equal(t, tc.expectedErr, err)

		fmt.Println("Test Generate New Schema Mutation| Testcase:", in+1, " test case name:", tc.casename)
	}

}
