package handlers

import (
	"authentication/constants"
	"authentication/models"
	"authentication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Validate OTP for user
// @Description Validates the OTP for the user provided in the request body.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.User.email models.User.otp  true "User data including OTP for validation"
// @Success 200 {string} string "OTP validation successful"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "OTP validation failed"
// @Router /customer-otpvalidate [post]
func ValidateOtp(user *service.OtpVerificationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData models.User
		if err := ctx.ShouldBind(&userData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{constants.ErrGeneric.Error(): err.Error()})
			return
		}
		if err := user.OtpVerification(userData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{constants.ErrGeneric.Error(): err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{constants.StatusKey: constants.SignInSuccessMessage})
	}
}
