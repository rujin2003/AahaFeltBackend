package model

type User struct {
	UserId     string `json:"userid"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Phoneno    string `json:"phoneno"`
	Country    string `json:"country"`
	PostalCode string `json:"postalcode"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Orders     []int  `json:"orders"`
}
