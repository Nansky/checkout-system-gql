package commons_test

import (
	"checkout-system-gql/commons"
	"checkout-system-gql/structs"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ItemPromotion struct {
	caseName         string
	ItemName         string
	ItemPrice        float64
	ItemQuantity     int
	ResponseCheckout structs.ResponseCheckout
}

func TestMacbookPromotion(t *testing.T) {
	testcases := []ItemPromotion{
		{
			caseName:     "Get macbook promotion with 2 items. Free 2 Raspberry",
			ItemName:     "MacBook Pro",
			ItemPrice:    5499.99,
			ItemQuantity: 2,
			ResponseCheckout: structs.ResponseCheckout{
				Desc:  "Scanned Items: MacBook Pro, Raspberry Pi B, MacBook Pro, Raspberry Pi B",
				Total: "$10999.98",
			},
		},
		{
			caseName:     "Get macbook promotion with 4 items. Free 4 Raspberry",
			ItemName:     "MacBook Pro",
			ItemPrice:    5499.99,
			ItemQuantity: 4,
			ResponseCheckout: structs.ResponseCheckout{
				Desc:  "Scanned Items: MacBook Pro, Raspberry Pi B, MacBook Pro, Raspberry Pi B, MacBook Pro, Raspberry Pi B, MacBook Pro, Raspberry Pi B",
				Total: "$21999.96",
			},
		},
	}

	for in, tc := range testcases {
		res := commons.GetMacbookPromotion(tc.ItemName, tc.ItemPrice, tc.ItemQuantity)
		assert.Equal(t, tc.ResponseCheckout, res)

		fmt.Println("Test Macbook Promotion | Testcase:", in+1, " test case name:", tc.caseName)
	}
}

func TestGoogleHomePromotion(t *testing.T) {
	testCase := []ItemPromotion{
		{
			caseName:     "Get Google Home promotion with 4 items. buyer will get price  3",
			ItemName:     "Google Home",
			ItemPrice:    49.99,
			ItemQuantity: 4,
			ResponseCheckout: structs.ResponseCheckout{
				Desc:  "Scanned Items: Google Home, Google Home, Google Home, Google Home",
				Total: "$149.97",
			},
		},
		{
			caseName:     "Get Google Home promotion with 2 items. buyer will get price  2",
			ItemName:     "Google Home",
			ItemPrice:    49.99,
			ItemQuantity: 2,
			ResponseCheckout: structs.ResponseCheckout{
				Desc:  "Scanned Items: Google Home, Google Home",
				Total: "$99.98",
			},
		},
	}

	for in, tc := range testCase {
		res := commons.GetGoogleHomePromotion(tc.ItemName, tc.ItemPrice, tc.ItemQuantity)
		assert.Equal(t, tc.ResponseCheckout, res)

		fmt.Println("Test Google Home Promotion | Testcase:", in+1, " test case name:", tc.caseName)
	}
}

func TestAlexaSpeakerPromotion(t *testing.T) {
	testCase := []ItemPromotion{
		{
			caseName:     "Get Google Home promotion with 2 items. buyer will not get any discount ",
			ItemName:     "Alexa Speaker",
			ItemPrice:    109.50,
			ItemQuantity: 2,
			ResponseCheckout: structs.ResponseCheckout{
				Desc:  "Scanned Items: Alexa Speaker, Alexa Speaker",
				Total: "$219.00",
			},
		},
		{
			caseName:     "Get Google Home promotion with 3 items. buyer will get 10% discount of all items",
			ItemName:     "Alexa Speaker",
			ItemPrice:    109.50,
			ItemQuantity: 3,
			ResponseCheckout: structs.ResponseCheckout{
				Desc:  "Scanned Items: Alexa Speaker, Alexa Speaker, Alexa Speaker",
				Total: "$328.50",
			},
		},
	}

	for in, tc := range testCase {
		res := commons.GetAlexaSpeakerPromotion(tc.ItemName, tc.ItemPrice, tc.ItemQuantity)
		assert.Equal(t, tc.ResponseCheckout, res)

		fmt.Println("Test Alexa Speaker Promotion | Testcase:", in+1, " test case name:", tc.caseName)
	}
}
