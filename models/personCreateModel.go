package models

type PersonRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	StreetLine1 string `json:"street1,omitempty"`
	StreetLine2 string `json:"street2,omitempty"`
	PostalCode  string `json:"zip_code,omitempty"`
	Age         int    `json:"age,omitempty"`
}
