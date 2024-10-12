package service

import (
	"database/sql"
	"fmt"
	"infilon_practicaltask/database"
	"infilon_practicaltask/models"
)

func GetPersonDetails(personID int) (*models.PersonInfo, error) {
	// Get a database connection
	conn, err := database.Connect()
	if err != nil {
		return nil, fmt.Errorf("could not establish a database connection: %w", err)
	}
	defer conn.Close()

	// SQL query to get person details
	query := `
		SELECT p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code
		FROM person p
		JOIN phone ph ON p.id = ph.person_id
		JOIN address_join aj ON p.id = aj.person_id
		JOIN address a ON aj.address_id = a.id
		WHERE p.id = ?;
	`

	var personDetails models.PersonInfo
	// Execute the query and map the result to the personDetails struct
	err = conn.QueryRow(query, personID).Scan(
		&personDetails.FullName,
		&personDetails.PhoneNumber,
		&personDetails.City,
		&personDetails.State,
		&personDetails.StreetLine1,
		&personDetails.StreetLine2,
		&personDetails.PostalCode,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving person data for ID %d: %w", personID, err)
	}
	return &personDetails, nil
}
