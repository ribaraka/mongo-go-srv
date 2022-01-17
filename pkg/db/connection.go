package db

import (
	"context"
	"fmt"
	"github.com/ribaraka/mongo-go-srv/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetMongoDbConnection(c config.Config) *mongo.Client {
	var credentials = options.Credential{
		Username: c.Mongo.Credentials.Username,
		Password: c.Mongo.Credentials.Password,
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.Mongo.ServerHost).SetAuth(credentials))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}
