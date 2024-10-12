package service

import (
	"errors"
	"fmt"
	"infilon_practicaltask/database"
	"infilon_practicaltask/models"
)

func CreatePerson(person models.PersonRequest) error {
	// Validate required fields
	if person.Name == "" {
		return errors.New("name is required")
	}
	if person.PhoneNumber == "" {
		return errors.New("phone number is required")
	}

	conn, err := database.Connect()
	if err != nil {
		return errors.New("could not establish a database connection")
	}
	defer conn.Close()

	// Start a transaction
	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert person
	personInsert := `INSERT INTO person (name, age) VALUES (?, ?)`
	res, err := tx.Exec(personInsert, person.Name, person.Age)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting person: %v", err)
	}

	personID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error fetching person ID: %v", err)
	}

	//insert phone
	phoneInsert := `INSERT INTO phone (number, person_id) VALUES (?, ?)`
	_, err = tx.Exec(phoneInsert, person.PhoneNumber, personID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting phone: %v", err)
	}

	//insert address
	addressInsert := `INSERT INTO address (city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)`
	res, err = tx.Exec(addressInsert, person.City, person.State, person.StreetLine1, person.StreetLine2, person.PostalCode)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting address: %v", err)
	}

	addressID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error fetching address ID: %v", err)
	}

	//insert address join
	addressJoinInsert := `INSERT INTO address_join (person_id, address_id) VALUES (?, ?)`
	_, err = tx.Exec(addressJoinInsert, personID, addressID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting address join: %v", err)
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	return nil
}
