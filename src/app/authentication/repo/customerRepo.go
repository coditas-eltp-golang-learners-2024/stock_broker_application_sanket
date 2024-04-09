package repo

import (
	"authentication/constants"
	"authentication/utils/db"

	"log"
)

func IsEmailExist(email string) bool {

	DatabaseConnection, err := db.ConnectionWithDb()

	if err != nil {
		log.Fatalf("%s :%s", constants.ErrConnectingDB.Error(), err)
	}

	query := "SELECT COUNT(*) FROM customerdata WHERE email = ?"

	var count int
	err = DatabaseConnection.QueryRow(query, email).Scan(&count)
	if err != nil {
		log.Fatalf("%s :%s", constants.ErrCheckingEmail.Error(), err)
	}

	return count > 0
}
func IsPanCardExist(Pancard string) bool {

	DatabaseConnection, err := db.ConnectionWithDb()

	if err != nil {
		log.Fatalf("%s :%s", constants.ErrConnectingDB.Error(), err)
	}

	query := "SELECT COUNT(*) FROM customerdata WHERE pancardnumber = ?"

	var count int
	err = DatabaseConnection.QueryRow(query, Pancard).Scan(&count)
	if err != nil {
		log.Fatalf("%s :%s", constants.ErrCheckingPanCard.Error(), err)
	}

	return count > 0
}
