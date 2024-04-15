package main

import (
	"gowork/go-gin-gorm/config"
	"gowork/go-gin-gorm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
