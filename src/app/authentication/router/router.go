package router

import (
	"authentication/constants"
	"authentication/handlers"
	"authentication/repository"
	"authentication/service"
	"authentication/utils/db"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// router is created with gin
func InitializeRouter() *gin.Engine {
	userRouter := gin.New()
	userRouter.Use(gin.Logger()) // LOG DETAILS ON CONSOLE
	ConnectionWithDb, err := db.ConnectionWithDb()
	// checking err while connectingDB
	if err != nil {
		log.Fatalf("%s :%s", constants.ErrConnectingDB.Error(), err.Error())
	}
	userDatabaseRepo := repository.NewUserDBRepository(ConnectionWithDb)
	userService := service.NewUserRepo(userDatabaseRepo)
	userRouter.POST(constants.CustomerSignupEndpoint, handlers.RegisterCustomer(userService))
	userRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))		// swaggerAdded

	// SignIn
	// isAuthorizedEmail := service.IsAuthorizedEmail
	// isAuthorizedPswd := service.IsAuthorizedPswd
	// isAuthorizedCustomer := service.IsAuthorizedCustomer

	// // Define your handler function with injected dependencies
	// handler := func(c *gin.Context) {
	// 	handlers.CustomerAuthentication(c, isAuthorizedEmail, isAuthorizedPswd, isAuthorizedCustomer)
	// }

	// // Register your handler with the router
	// userRouter.POST(constants.CustomerSigninEndpoint, handler)

	return userRouter
}
