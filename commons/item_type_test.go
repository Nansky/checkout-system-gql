package commons_test

import (
	"checkout-system-gql/commons"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ItemTypeInt struct {
	itemType         int
	expectedResSKU   string
	expectedResPrice string
	expectedResName  string
}

func TestGetItemDetailFromType(t *testing.T) {
	testCase := []ItemTypeInt{
		{
			itemType:         1,
			expectedResSKU:   "43N23P",
			expectedResName:  "MacBook Pro",
			expectedResPrice: "5399.99",
		},
		{
			itemType:         2,
			expectedResSKU:   "120P90",
			expectedResName:  "Google Home",
			expectedResPrice: "49.99",
		},
		{
			itemType:         3,
			expectedResSKU:   "A304SD",
			expectedResName:  "Alexa Speaker",
			expectedResPrice: "109.50",
		},
		{
			itemType:         4,
			expectedResSKU:   "234234",
			expectedResName:  "Raspberry Pi B",
			expectedResPrice: "30.00",
		},
		{
			itemType: 99,
		},
	}

	for in, tc := range testCase {
		sku, name, price := commons.GetItemDetail(tc.itemType)
		assert.Equal(t, tc.expectedResSKU, sku)
		assert.Equal(t, tc.expectedResName, name)
		assert.Equal(t, tc.expectedResPrice, price)

		fmt.Println("Test get item Detail From itemType case: ", in+1, " | Item name", name)
	}

}
