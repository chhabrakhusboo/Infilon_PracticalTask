package models

type PersonInfo struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	StreetLine1 string `json:"street1"`
	StreetLine2 string `json:"street2"`
	PostalCode  string `json:"zip_code"`
}
