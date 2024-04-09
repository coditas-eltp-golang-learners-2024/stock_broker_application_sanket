package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/utils/db"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

// validateWithPattern validates a string against a regular expression pattern.
func validateWithPattern(input string, pattern string) bool {
	regExp, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("%s: %s", constants.ErrCompileRegex, err)
	}
	return regExp.MatchString(input)
}

// ValidatedEmail checks email format.
func ValidatedEmail(email string) bool {
	return validateWithPattern(email, constants.ValidatedEmailPattern)
}

// ValidatePANCard checks PAN card number  format.
func ValidatePANCard(panCard string) bool {
	return validateWithPattern(panCard, constants.ValidatePANCardNumber)
}

// ValidatePassword checks password meets required criteria or not.
func ValidatePassword(password string) bool {
	return validateWithPattern(password, constants.ValidatePassword)
}

// Validation checks for required fields
func ValidateCustomer(customer models.Customer) error {
	if err := Validate.Struct(customer); err != nil {
		return err
	}
	return nil
}

// inserting data into database
func CreateNewCustomer(customerRecords models.Customer, c *gin.Context) error {
	// Get database connection
	DatabaseConnection, err := db.ConnectionWithDb()
	if err != nil {
		log.Fatalf("%s: %s", constants.ErrConnectingDB, err)
	}

	// query to inert record to database
	query := "INSERT INTO customerdata (Name, EMail, PhoneNumber, PancardNumber, Password) VALUES (?, ?, ?, ?, ?)"

	_, err = DatabaseConnection.Exec(query, customerRecords.Name, customerRecords.Email, customerRecords.PhoneNumber, customerRecords.PanCardNumber, customerRecords.Password)

	if err != nil {
		log.Fatalf("%s: %s", constants.ErrExecutingQuery, err)
	}

	// Appending data
	models.CustomerData = append(models.CustomerData, customerRecords)

	c.JSON(http.StatusCreated, customerRecords)

	return nil
}
