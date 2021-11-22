package services

import (
	"checkout-system-gql/commons"
	"checkout-system-gql/repositories"
	"checkout-system-gql/structs"
	"fmt"

	"github.com/graphql-go/graphql"
)

type ICheckoutServices interface {
	GetItemsResolver(params graphql.ResolveParams) (dataresp interface{}, err error)
	AddItemsResolver(params graphql.ResolveParams) (dataresp interface{}, err error)
	CheckoutItemsResolver(params graphql.ResolveParams) (dataresp interface{}, err error)
}

type CheckoutServices struct {
	checkoutRepo repositories.ICheckoutRepositories
}

func NewCheckoutServices(r repositories.ICheckoutRepositories) ICheckoutServices {
	return &CheckoutServices{
		checkoutRepo: r,
	}
}

func (cs *CheckoutServices) GetItemsResolver(params graphql.ResolveParams) (dataresp interface{}, err error) {
	buyerId, err := getStringParams("buyerId", params.Args, true)
	if err != nil {
		return nil, err
	}

	return cs.checkoutRepo.RetrieveItems(buyerId)
}

func (cs *CheckoutServices) AddItemsResolver(params graphql.ResolveParams) (dataresp interface{}, err error) {
	buyerId, err := getStringParams("buyerId", params.Args, true)
	if err != nil {
		return nil, err
	}

	itemType, err := getIntParams("itemType", params.Args, true)
	if err != nil {
		return nil, err
	}
	sku, name, price := commons.GetItemDetail(itemType)

	quantity, err := getIntParams("quantity", params.Args, true)
	if err != nil {
		return nil, err
	}

	addedItems := structs.ItemsList{
		BuyerID:  buyerId,
		Sku:      sku,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}

	// store to Redis
	cacheKey := fmt.Sprintf("buyer_%s:%s", buyerId, sku)
	err = cs.checkoutRepo.StoreListItems(cacheKey, addedItems)
	if err != nil {
		return nil, err
	}

	dataresp = commons.AddToCartResponse(addedItems)

	return dataresp, nil
}

func (cs *CheckoutServices) CheckoutItemsResolver(params graphql.ResolveParams) (dataresp interface{}, err error) {
	buyerId, err := getStringParams("buyerId", params.Args, true)
	if err != nil {
		return nil, err
	}

	allItemList, err := cs.checkoutRepo.RetrieveItems(buyerId)
	if err != nil {
		return nil, err
	}

	err = cs.checkoutRepo.DeleteItemsByBuyerID(buyerId)
	if err != nil {
		return []structs.ResponseCheckout{}, err
	}

	return commons.GetPromotionResponse(allItemList), nil
}

func getStringParams(k string, args map[string]interface{}, required bool) (string, error) {
	// first check presense of arg
	if value, ok := args[k]; ok {
		// check string datatype
		v, o := value.(string)
		if !o {
			return "", fmt.Errorf("%s is not a string value", k)
		}
		return v, nil
	}
	if required {
		return "", fmt.Errorf("missing argument %s", k)
	}
	return "", nil
}

func getIntParams(k string, args map[string]interface{}, required bool) (int, error) {
	// first check presense of arg
	if value, ok := args[k]; ok {
		// check string datatype
		v, o := value.(int)
		if !o {
			return 0, fmt.Errorf("%s is not a string value", k)
		}
		return v, nil
	}
	if required {
		return 0, fmt.Errorf("missing argument %s", k)
	}
	return 0, nil
}
