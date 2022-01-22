package controller

import (
	"IdentityUser/helper"
	"IdentityUser/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Authentication validates token and authorizes users
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, model.Error{Message: "No Authorization header provided"}.GetAsEnvelope())
			return
		}

		_, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, model.Error{Message: err}.GetAsEnvelope())
			return
		}

		//c.Set("email", claims.Email)
		//c.Set("firstName", claims.UserName)
		//c.Set("uid", claims.Uid)

		c.JSON(http.StatusOK, model.Success{}.True())
		return
	}
}
