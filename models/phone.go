package models

type Phone struct {
	ID       int    `json:"id"`
	Number   string `json:"number"`
	PersonID int    `json:"person_id"`
}
