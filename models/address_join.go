package models

type AddressJoin struct {
	ID        int `json:"id"`
	PersonID  int `json:"person_id"`
	AddressID int `json:"address_id"`
}
