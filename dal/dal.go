package dal

import (
	"IdentityUser/database"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func GetUserCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, os.Getenv("USER_COLLECTION_NAME"))
}

func GetDatabase() *mongo.Client {
	return database.Client
}