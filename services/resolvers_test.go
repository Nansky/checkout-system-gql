package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CheckoutResolversSuite struct {
	suite.Suite
}

type GetItemResolversCase struct {
	param       string
	expectedRes interface{}
}

var keyVal = map[string]interface{}{
	"buyerId":  "212",
	"quantity": 9,
	"itemType": "MacBook Pro",
}

func TestCheckoutResolvers(t *testing.T) {
	suite.Run(t, new(CheckoutResolversSuite))
}

func (c *CheckoutResolversSuite) SetupTest() {
	mockCtrl := gomock.NewController(c.T())
	defer mockCtrl.Finish()

}

func (c *CheckoutResolversSuite) TestGetItemResolver() {
	testcases := []GetItemResolversCase{
		{
			param:       "buyerId",
			expectedRes: keyVal["buyerId"],
		},
		{
			param:       "itemType",
			expectedRes: keyVal["itemType"],
		},
	}

	for _, tc := range testcases {
		res, _ := getStringParams(tc.param, keyVal, true)
		assert.Equal(c.T(), tc.expectedRes, res)

	}
}
