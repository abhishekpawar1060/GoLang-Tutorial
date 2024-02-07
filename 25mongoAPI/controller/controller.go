package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

// Most IMP
var collection *mongo.Collection

// connect with mongoDB
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongdb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection sucess")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is Ready")
}
