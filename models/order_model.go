package model

type Order struct {
	ID int `json:"id"`

	ProductName string `json:"productname"`
	ProductID   int    `json:"product_id"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	TotalPrice  int    `json:"totalprice"`
	OrderDate   string `json:"orderdate"`

	UserName   string `json:"name"`
	Phoneno    string `json:"phoneno"`
	Country    string `json:"country"`
	PostalCode string `json:"postalcode"`
	Street     string `json:"street"`
	City       string `json:"city"`
	ClientId   int    `json:"client_id"`
}
