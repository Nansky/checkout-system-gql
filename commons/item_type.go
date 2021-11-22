package commons

func GetItemDetail(itemtype int) (sku, name, price string) {
	switch itemtype {
	case 1:
		sku, name, price = "43N23P", "MacBook Pro", "5399.99"
	case 2:
		sku, name, price = "120P90", "Google Home", "49.99"
	case 3:
		sku, name, price = "A304SD", "Alexa Speaker", "109.50"
	case 4:
		sku, name, price = "234234", "Raspberry Pi B", "30.00"
	}

	return
}
