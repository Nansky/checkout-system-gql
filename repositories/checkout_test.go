package repositories_test

import (
	"checkout-system-gql/mocks"
	"checkout-system-gql/repositories"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CheckoutRepositoriesSuite struct {
	suite.Suite
	repo     repositories.ICheckoutRepositories
	mockRepo *mocks.MockCacheInterface
}

func TestCheckoutRepo(t *testing.T) {
	suite.Run(t, new(CheckoutRepositoriesSuite))
}

// func (c *CheckoutRepositoriesSuite) SetupTestSuite() {
// 	mockCtrl := gomock.NewController(c.T())

// 	c.mockRepo = mocks.NewMockCacheInterface(mockCtrl)
// 	c.repo = repositories.NewCheckoutRepositories(c.mockRepo)
// }

// func (c *CheckoutRepositoriesSuite) TestGetItemResolver() {
// 	expectedKeys := []string{"123-123", "B90X90", "EX092"}

// 	c.mockRepo.EXPECT().Get("any-key").Return("any-val", nil)
// 	c.mockRepo.EXPECT().GetAllKeys("buyer_id123").Return(expectedKeys, nil)

// 	_, err := c.repo.RetrieveItems("buyer_id123")
// 	assert.Equal(c.T(), nil, err)
// }
