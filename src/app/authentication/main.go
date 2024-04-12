package main

import (
	_ "authentication/docs"
	"authentication/router"
	"fmt"
)

//@title Stock Broker Application
//@description   api for Stock Broker using gin and gorm
//@version 2.0

// @host localhost:8080
func main() {
	// router initialize on port 8080
	apiRouter := router.InitializeRouter()
	fmt.Println("The Connection is open at port-8080")
	apiRouter.Run(":8080")
}
