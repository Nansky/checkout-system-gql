package repositories

import (
	"checkout-system-gql/cache"
	"checkout-system-gql/structs"
	"encoding/json"
	"time"
)

type ICheckoutRepositories interface {
	RetrieveItems(buyerId string) (res []structs.ItemsList, err error)
	StoreListItems(cacheKey string, data structs.ItemsList) (err error)
	DeleteItemsByBuyerID(buyerId string) (err error)
}

type CheckoutRepositories struct {
	Cache cache.CacheInterface
}

func NewCheckoutRepositories(c cache.CacheInterface) ICheckoutRepositories {

	return &CheckoutRepositories{
		Cache: c,
	}
}

func (cr *CheckoutRepositories) RetrieveItems(buyerId string) (res []structs.ItemsList, err error) {
	var item structs.ItemsList
	var items []structs.ItemsList

	keys, err := cr.Cache.GetAllKeys(buyerId)
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		v, _ := cr.Cache.Get(key)
		_ = json.Unmarshal(v, &item)

		items = append(items, item)
	}

	return items, nil
}

func (cr *CheckoutRepositories) StoreListItems(cacheKey string, data structs.ItemsList) (err error) {
	ttl := 1 * time.Hour
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = cr.Cache.Write(cacheKey, dataByte, ttl)
	if err != nil {
		return err
	}

	return nil

}

func (cr *CheckoutRepositories) DeleteItemsByBuyerID(buyerId string) (err error) {
	err = cr.Cache.FlushByBuyerId(buyerId)
	if err != nil {
		return err
	}

	return err
}
