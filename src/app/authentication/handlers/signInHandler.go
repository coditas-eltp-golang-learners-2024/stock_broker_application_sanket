package handlers

import (
	"authentication/constants"
	"authentication/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerAuthentication(c *gin.Context, isAuthorizedEmail func(email string) bool,
	isAuthorizedPswd func(password string) bool,
	isAuthorizedCustomer func(email, password string) (*models.Customer, bool)) {


	var custCreds models.UserCredentials

	if err := c.ShouldBind(&custCreds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isAuthorizedEmail(custCreds.Email) {
		// If email is not authorized, return response with bad request status
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrEmail.Error()})
		return
	}

	if !isAuthorizedPswd(custCreds.Password) {
		// If password is not authorized, return response with bad request status
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrPassword.Error()})
		return
	}

	data, status := isAuthorizedCustomer(custCreds.Email, custCreds.Password)

	if !status {
		// Customer not found, return appropriate JSON response
		c.JSON(http.StatusNotFound, gin.H{"error": constants.ErrCustomerNotFound.Error()})
		return
	}

	// Customer found, return their data
	c.JSON(http.StatusOK, data)
}
