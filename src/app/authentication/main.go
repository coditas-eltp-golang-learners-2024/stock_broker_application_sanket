package main

import (
	"authentication/router"
	"fmt"
)

func main() {
	// router initialize on port 8080
	apiRouter := router.InitializeRouter()
	fmt.Println("The Connection is open at port-8080")
	apiRouter.Run(":9091")

}
