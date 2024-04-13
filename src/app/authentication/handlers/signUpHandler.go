package handlers

import (
	"authentication/constants"
	"authentication/models"
	"authentication/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

// RegisterCustomer godoc
// @Summary Register a new customer
// @Description Register a new customer and save their data in the database.
// @Accept json
// @Produce json
// @Param customerRecords body models.Customer true "Customer data"
// @Success 200 {object} models.Customer
// @Router /customer-signup [post]
func RegisterCustomer(user *service.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customerRecords models.Customer
		if err := c.ShouldBind(&customerRecords); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{constants.ErrGeneric.Error(): err.Error()})
			return
		}

		if err := user.SignUp(customerRecords); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{constants.ErrGeneric.Error(): err.Error()})
			return
		}
		c.JSON(http.StatusOK, customerRecords)
	}
}
