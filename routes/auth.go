package routes

import (
	"IdentityUser/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/token/validate", controller.Authentication())
}
