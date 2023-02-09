package main

import (
	"tutorial/configs"
	"tutorial/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	configs.ConnectDB()

	v1 := router.Group("/v1")

	routes.HandwashRoute(v1)
	routes.StateRoute(v1)

	router.Run("localhost:8080")
}
