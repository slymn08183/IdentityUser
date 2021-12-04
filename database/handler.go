package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// CreateUniqueIndex create UniqueIndex
func CreateUniqueIndex(collection *mongo.Collection, key string) {

	idxRet, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: key, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		options.CreateIndexes().SetMaxTime(10*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.Indexes().CreateOne:", idxRet)
}
