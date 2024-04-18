package handlers

import (
	"authentication/constants"
	"authentication/models"
	"authentication/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary User Sign In
// @Description Signs in a user with provided credentials.
// @Tags users
// @Accept json
// @Produce json
// @Param body body models.SignInCredentials true "Sign In Credentials"
// @Success 200 {object} models.SignInCredentials "User signed in successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /customer-signin [post]
func UserSignInHandler(user *service.SignInChecker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signInData models.SignInCredentials
		if err := ctx.ShouldBind(&signInData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{constants.ErrGeneric.Error(): err.Error()})
			return
		}

		if err := user.SignIn(signInData); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{constants.ErrGeneric.Error(): err.Error()})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{constants.Otp:constants.OtpMessage})
	}
}
