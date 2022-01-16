package controller

import (
	"IdentityUser/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Success{}.True())
		return
	}
}
