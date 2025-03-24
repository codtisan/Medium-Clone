package internal

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MongoClient *mongo.Client
)

var (
	UserCollection *mongo.Collection
)

func InitMongoDB() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/"))
	if err != nil {
		panic(err)
	}
	MongoClient = client
	UserCollection = client.Database("local").Collection("user")
}
