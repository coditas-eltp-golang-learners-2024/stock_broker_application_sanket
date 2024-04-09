package router

import (
	"authentication/constants"
	"authentication/handlers"
	"github.com/gin-gonic/gin"
)

// router is created with gin
func InitializeRouter() *gin.Engine {
	userRouter := gin.New()
	userRouter.Use(gin.Logger()) // LOG DETAILS ON CONSOLE
	userRouter.POST(constants.CustomerSignupEndpoint, handlers.RegisterCustomer)
	return userRouter
}
