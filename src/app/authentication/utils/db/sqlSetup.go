package db

import (
	"database/sql"
	"fmt"

	"authentication/constants"
	"authentication/utils"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// databases connection creation (mysql)
func ConnectionWithDb() (*sql.DB, error) {

	// Parse Details from Yaml to struct
	sqlCredentials := utils.ParseYAML()

	DatabaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		sqlCredentials.DBUsername,
		sqlCredentials.DBPassword,
		sqlCredentials.DBHost,
		sqlCredentials.DBPort,
		sqlCredentials.DBName)

	ConnectToDatabase, err := sql.Open("mysql", DatabaseURL)

	if err != nil {
		log.Fatalf("%s: %s",constants.ErrOpenDatabaseConnection.Error() ,err)
		return nil, err
	}

	// Ping connection (Status check)
	err1 := ConnectToDatabase.Ping()

	if err1 != nil {
		log.Fatalf("%s: %s",constants.ErrDatabasePing,err)
	}

	fmt.Println("Connected to database!")

	return ConnectToDatabase, nil
}
