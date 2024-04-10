package router

import (
	"authentication/constants"
	"authentication/handlers"
	"authentication/service"

	"github.com/gin-gonic/gin"
)

// router is created with gin
func InitializeRouter() *gin.Engine {
	userRouter := gin.New()
	userRouter.Use(gin.Logger()) // LOG DETAILS ON CONSOLE
	userRouter.POST(constants.CustomerSignupEndpoint, handlers.RegisterCustomer)



	// SignIn 
	isAuthorizedEmail := service.IsAuthorizedEmail
	isAuthorizedPswd := service.IsAuthorizedPswd
	isAuthorizedCustomer := service.IsAuthorizedCustomer

	// Define your handler function with injected dependencies
	handler := func(c *gin.Context) {
		handlers.CustomerAuthentication(c, isAuthorizedEmail, isAuthorizedPswd, isAuthorizedCustomer)
	}

	// Register your handler with the router
	userRouter.POST(constants.CustomerSigninEndpoint, handler)
	
	return userRouter
}
