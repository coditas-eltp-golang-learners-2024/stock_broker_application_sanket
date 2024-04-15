package handlers

import (
	"authentication/constants"
	"authentication/models"
	"authentication/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateOtp(user *service.OtpVerificationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData models.User

		if err := ctx.ShouldBind(&userData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{constants.ErrGeneric.Error(): err.Error()})
			return
		}
		if err:= user.OtpVerification(userData); err !=nil {
			ctx.JSON(http.StatusBadRequest,gin.H{constants.ErrGeneric.Error(): err.Error()})
		}
	

		ctx.JSON(http.StatusOK, userData)
	}
}
