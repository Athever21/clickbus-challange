package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conn *mongo.Client

func initDB() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	conn = client

	if err != nil {
		log.Fatal(err)
	}
}

func GetDb() *mongo.Client {
	if conn == nil {
		initDB()
	}

	return conn
}

func CloseDb() {
	conn.Disconnect(context.TODO())
}
