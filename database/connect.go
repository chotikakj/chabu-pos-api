package database

import (
	"context"
	"fmt"
	"log"
	"pos-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB mongo.Database

func InitialMongoDB() {
	clientOptions := options.Client().ApplyURI(config.GetEnvConfig("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	MongoDB = *client.Database("iampos")
}
