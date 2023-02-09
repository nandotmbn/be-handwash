package routes

import (
	controller_handwash "tutorial/controllers/handwash"

	"github.com/gin-gonic/gin"
)

func HandwashRoute(router *gin.RouterGroup) { // http://localhost:8080/v1
	router.POST("/handwash", controller_handwash.RegisterHandwash()) // http://localhost:8080/v1/handwash
	router.POST("/handwash/retriveid", controller_handwash.GetIdHandwash())
}
