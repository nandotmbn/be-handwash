package routes

import (
	controller_state "tutorial/controllers/state"

	"github.com/gin-gonic/gin"
)

func StateRoute(router *gin.RouterGroup) {
	router.GET("/state/:handwash_id", controller_state.GetState())
	router.PUT("/state/:handwash_id", controller_state.UpdateState())
}
