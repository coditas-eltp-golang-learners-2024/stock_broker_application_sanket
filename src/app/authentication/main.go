package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router:=gin.New()

	router.Use(gin.Logger())

	router.Run(":9091")
}