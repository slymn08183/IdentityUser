package routes

import (
	"IdentityUser/controller"
	"github.com/gin-gonic/gin"
)

func HealthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/health/check", controller.HealthCheck())
}
