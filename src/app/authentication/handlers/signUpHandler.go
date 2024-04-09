package handlers

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repo"
	"authentication/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// customer Registeration function
func RegisterCustomer(c *gin.Context) {
	var customerRecords models.Customer

	// binding data into json
	if err := c.ShouldBind(&customerRecords); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	// Email check
	if !service.ValidatedEmail(customerRecords.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidEmail.Error()})
		return
	}

	if repo.IsEmailExist(customerRecords.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrEmailExsits.Error()})
		return

	}

	// PanCardNumber check
	if !service.ValidatePANCard(customerRecords.PanCardNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidPanCardFormat.Error()})
		return
	}

	if repo.IsPanCardExist(customerRecords.PanCardNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrPanCardlExsits.Error()})
		return

	}

	// Password checks
	if !service.ValidatePassword(customerRecords.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidPasswordFormat.Error()})
		return
	}

	// Customer Data Validation
	if err := service.ValidateCustomer(customerRecords); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// FinalCall to insert customer to database
	err := service.CreateNewCustomer(customerRecords,c)
	if err != nil {
		log.Fatalf("%s: %s",constants.ErrInsertingCustomerRecord ,err)
		return
	}
}
