package delegate

type Customer struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

type Address struct {
	ID      int64  `json:"id"`
	CID     int64  `json:"cid"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

type Item struct {
	ID          int64   `json:"id"`
	SKU         string  `json:"sku"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Order struct {
	ID          int64   `json:"id"`
	CID         int64   `json:"cid"`
	OrderNumber string  `json:"oid"`
	Items       *[]Item `json:"items"`
}
