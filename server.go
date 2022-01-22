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

func main() {

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
	routes.HealthRoutes(router)

	router.Use(middleware.Authentication())

	router.Run(":" + port)
}
