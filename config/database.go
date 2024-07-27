package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


type MongoInstance struct {
	Client *mongo.Client
	Db    *mongo.Database
}

var MongoInstanceDB MongoInstance

const MONGO_URI = "mongodb://localhost:27017/fiber"


func ConnectToMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer cancel()
	if err != nil {
		defer client.Disconnect(ctx)
		log.Fatal(err)
	}
	
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}
	println("Connected to MongoDB!..")
	todoDB := client.Database("fiber-todo")
	
	MongoInstanceDB = MongoInstance{
		Client: client,
		Db: todoDB,
	}
}