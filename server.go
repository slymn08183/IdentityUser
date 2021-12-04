package main

import (
	"IdentityUser/dal"
	"IdentityUser/database"
	"IdentityUser/middleware"
	"IdentityUser/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main(){

	//<editor-fold desc="DB CONFIGURATION">
	database.CreateUniqueIndex(dal.GetUserCollection(), "email")
	//</editor-fold>

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
		log.Fatal(port)
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	// API-2
	router.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	// API-1
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
