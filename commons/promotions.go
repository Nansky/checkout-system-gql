package commons

import (
	"checkout-system-gql/structs"
	"fmt"
	"strings"
)

func getMacbookPromotion(name string, price float64, quantity int) (resp structs.ResponseCheckout) {
	var scanned []string
	priceStr := fmt.Sprintf("$%.2f", price*float64(quantity))

	for i := 0; i < quantity; i++ {
		scanned = append(scanned, name+", Raspberry Pi B")
	}

	joinedStr := strings.Join(scanned[:], ", ")

	resp.Desc = "Scanned Items: " + joinedStr
	resp.Total = priceStr

	return resp
}

func getGoogleHomePromotion(name string, price float64, quantity int) (resp structs.ResponseCheckout) {
	var scanned []string

	quantityDiv, quantityMod := (quantity/3)*2, quantity%3
	totalQuantity := quantityDiv + quantityMod

	priceStr := fmt.Sprintf("$%.2f", price*float64(totalQuantity))

	for i := 0; i < quantity; i++ {
		scanned = append(scanned, name)
	}

	joinedStr := strings.Join(scanned[:], ", ")

	resp.Desc = "Scanned Items: " + joinedStr
	resp.Total = priceStr

	return resp
}

func getAlexaSpeakerPromotion(name string, price float64, quantity int) (resp structs.ResponseCheckout) {
	var scanned []string

	if quantity > 3 {
		price = price - (price * 0.1)
	}
	priceStr := fmt.Sprintf("$%.2f", price*float64(quantity))

	for i := 0; i < quantity; i++ {
		scanned = append(scanned, name)
	}

	joinedStr := strings.Join(scanned[:], ", ")

	resp.Desc = "Scanned Items: " + joinedStr
	resp.Total = priceStr

	return resp
}
