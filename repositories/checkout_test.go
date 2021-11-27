package repositories_test

import (
	"checkout-system-gql/mocks"
	"checkout-system-gql/repositories"
	"checkout-system-gql/structs"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

type TestCaseRetrieveItems struct {
	caseName           string
	buyerId            string
	expectedKeys       []string
	expectedValFromKey []byte
	expectedError      error
}

type TestCaseStoredItems struct {
	caseName      string
	expectedError error
	cacheKey      string
	val           structs.ItemsList
}

type TestCaseDeleteItems struct {
	caseName      string
	buyerId       string
	expectedError error
}

func TestRetrieveItems(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	testCase := []TestCaseRetrieveItems{
		{
			caseName:           "Buyer ID Found, Get Value",
			buyerId:            "123",
			expectedKeys:       []string{"123-123", "B90X90", "EX092"},
			expectedValFromKey: []byte{69, 69},
			expectedError:      nil,
		},
		{
			caseName:           "Buyer ID Found, empty Value",
			buyerId:            "234",
			expectedKeys:       []string{"enak_234"},
			expectedValFromKey: []byte{},
			expectedError:      nil,
		},
		{
			caseName:      "Key Not Found, empty val",
			buyerId:       "488",
			expectedKeys:  []string{},
			expectedError: redis.ErrNil,
		},
		{
			caseName:      "Redis Error",
			buyerId:       "any-IDs",
			expectedKeys:  []string{},
			expectedError: redis.ErrPoolExhausted,
		},
	}

	mockRepo := mocks.NewMockCacheInterface(mockCtrl)

	for in, c := range testCase {
		mockRepo.EXPECT().
			GetAllKeys(c.buyerId).
			Return(c.expectedKeys, c.expectedError).
			AnyTimes()

		mockRepo.EXPECT().
			Get(gomock.Any()).
			Return(c.expectedValFromKey, c.expectedError).
			AnyTimes()

		repo := repositories.NewCheckoutRepositories(mockRepo)

		_, err := repo.RetrieveItems(c.buyerId)
		assert.Equal(t, c.expectedError, err)
		fmt.Println("Repositories Test | Retrieve Items on Testcase:", in+1, " test case name:", c.caseName)
	}
}

func TestDeleteItems(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tc := []TestCaseDeleteItems{
		{
			caseName:      "Success Delete redis key",
			buyerId:       "234",
			expectedError: nil,
		},
		{
			caseName:      "Delete failed, key not found or expired",
			buyerId:       "6969",
			expectedError: redis.ErrNil,
		},
		{
			caseName:      "Redis Err, Pool Exhausted",
			buyerId:       "12093",
			expectedError: redis.ErrPoolExhausted,
		},
	}
	mockRepo := mocks.NewMockCacheInterface(mockCtrl)
	repo := repositories.NewCheckoutRepositories(mockRepo)

	for in, c := range tc {
		mockRepo.EXPECT().
			FlushByBuyerId(c.buyerId).
			Return(c.expectedError).
			AnyTimes()

		err := repo.DeleteItemsByBuyerID(c.buyerId)
		assert.Equal(t, c.expectedError, err)

		fmt.Println("Repositories Test | Delete Items on Testcase:", in+1, " test case name:", c.caseName)
	}
}

func TestStoreItemsList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockCacheInterface(mockCtrl)
	ttl := 1 * time.Hour

	testCase := []TestCaseStoredItems{
		{
			caseName:      "Success Storing Items",
			expectedError: nil,
			cacheKey:      "buyer_id9999",
			val: structs.ItemsList{
				BuyerID:  "9999",
				Sku:      "DUN-020",
				Name:     "Dunhill Mild 20",
				Price:    "35000",
				Quantity: 2,
			},
		},
		{
			caseName:      "Key found but Empty Items",
			expectedError: nil,
			cacheKey:      "buyer_id11",
			val:           structs.ItemsList{},
		},
		{
			caseName:      "Failed to store Items, Pool exhausted",
			expectedError: redis.ErrPoolExhausted,
			cacheKey:      "buyer_id2",
			val: structs.ItemsList{
				BuyerID:  "2",
				Sku:      "EDC-4848",
				Name:     "any-item",
				Price:    "120000",
				Quantity: 1,
			},
		},
	}

	repo := repositories.NewCheckoutRepositories(mockRepo)

	for in, c := range testCase {
		data, _ := json.Marshal(c.val)
		mockRepo.EXPECT().
			Write(c.cacheKey, data, ttl).
			Return(c.expectedError).
			AnyTimes()

		err := repo.StoreListItems(c.cacheKey, c.val)
		assert.Equal(t, c.expectedError, err)

		fmt.Println("Repositories Test | Delete Items on Testcase:", in+1, " test case name:", c.caseName)
	}
}
