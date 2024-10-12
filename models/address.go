package models

type Address struct {
	ID         int64  `json:"id"`
	City       string `json:"city"`
	State      string `json:"state"`
	Street1    string `json:"street1"`
	Street2    string `json:"street2"`
	PostalCode string `json:"zip_code"`
}
