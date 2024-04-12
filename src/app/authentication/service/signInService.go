package service

// import (
// 	"authentication/constants"
// 	"authentication/models"
// 	"authentication/utils/db"
// 	"database/sql"
// 	"log"
// )

// func IsAuthorizedEmail(email string) bool {
// 	// Establish database connection
// 	dbConn, err := db.ConnectionWithDb()
// 	if err != nil {
// 		log.Fatalf("%s: %s", constants.ErrConnectingDB, err)
// 	}

// 	// Prepare query
// 	query := "SELECT COUNT(*) FROM customerdata WHERE Email = ?"

// 	// Execute query
// 	var count int
// 	err = dbConn.QueryRow(query, email).Scan(&count)
// 	if err != nil {
// 		log.Fatalf("%s: %s", constants.ErrQueryInDB.Error(), err)
// 	}

// 	// Check if email exists
// 	return count > 0
// }
// func IsAuthorizedPswd(pswd string) bool {
// 	// Establish database connection
// 	dbConn, err := db.ConnectionWithDb()
// 	if err != nil {
// 		log.Fatalf("%s: %s", constants.ErrConnectingDB.Error(), err)
// 	}

// 	// Prepare query
// 	query := "SELECT COUNT(*) FROM customerdata WHERE password = ?"

// 	// Execute query
// 	var count int
// 	err = dbConn.QueryRow(query, pswd).Scan(&count)
// 	if err != nil {
// 		log.Fatalf("%s: %s", constants.ErrQueryInDB.Error(), err)
// 	}

// 	// Check if email exists
// 	return count > 0
// }

// func IsAuthorizedCustomer(email, pswd string) (*models.Customer, bool) {
// 	// Establish database connection
// 	dbConn, err := db.ConnectionWithDb()
// 	if err != nil {
// 		log.Fatalf("%s: %s", constants.ErrConnectingDB.Error(), err)
// 	}

// 	// Prepare query
// 	query := "SELECT Name, Email, PhoneNumber, PancardNumber, Password FROM customerdata WHERE Email = ? AND Password = ?"

// 	// Execute query
// 	var customerData models.Customer
// 	err = dbConn.QueryRow(query, email, pswd).Scan(
// 		&customerData.Name,
// 		&customerData.Email,
// 		&customerData.PhoneNumber,
// 		&customerData.PancardNumber,
// 		&customerData.Password,
// 	)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// Customer not found
// 			return nil, false
// 		}
// 		log.Fatalf("%s: %s", constants.ErrQueryInDB.Error(), err)
// 	}

// 	return &customerData, true
// }
