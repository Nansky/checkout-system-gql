package commons

import (
	"checkout-system-gql/structs"
	"fmt"
	"strconv"
	"strings"
)

const (
	MacbookProSku = "43N23P"
	AlexaSpeaker  = "A304SD"
	GoogleHome    = "120P90"
	RaspberryPi   = "234234"
)

func AddToCartResponse(itemList structs.ItemsList) interface{} {
	price, _ := strconv.ParseFloat(itemList.Price, 64)
	priceStr := fmt.Sprintf("$%.2f", price)

	return structs.ResponseAddItems{
		Items: itemList.Name,
		Total: itemList.Quantity,
		Price: priceStr,
	}
}

func GetPromotionResponse(itemList []structs.ItemsList) interface{} {
	responseCheckout := []structs.ResponseCheckout{}
	var res structs.ResponseCheckout

	for _, s := range itemList {
		price, _ := strconv.ParseFloat(s.Price, 64)

		if s.Sku != RaspberryPi {
			switch s.Sku {
			case MacbookProSku:
				res = GetMacbookPromotion(s.Name, price, s.Quantity)
			case GoogleHome:
				res = GetGoogleHomePromotion(s.Name, price, s.Quantity)
			case AlexaSpeaker:
				res = GetAlexaSpeakerPromotion(s.Name, price, s.Quantity)
			}
		} else {
			res = getRaspberryResponse(s.Name, price, s.Quantity)
		}

		responseCheckout = append(responseCheckout, res)
	}

	return responseCheckout
}

func getRaspberryResponse(name string, price float64, quantity int) structs.ResponseCheckout {
	var scanned []string

	priceStr := fmt.Sprintf("$%.2f", price*float64(quantity))

	for i := 0; i < quantity; i++ {
		scanned = append(scanned, name)
	}

	joinedStr := strings.Join(scanned[:], ", ")

	rbp := structs.ResponseCheckout{
		Desc:  "Scanned Items: " + joinedStr,
		Total: priceStr,
	}

	return rbp
}
