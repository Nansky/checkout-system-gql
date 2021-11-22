package structs

type ItemsList struct {
	BuyerID  string
	Sku      string
	Name     string
	Price    string
	Quantity int
}

type ResponseAddItems struct {
	Items string
	Total int
	Price string
}

type ResponseCheckout struct {
	Desc  string
	Total string
}
