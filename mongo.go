package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(context context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(context, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the Database!")
	return client, err
}

func Disconnect(context context.Context, client *mongo.Client) {
	if err := client.Disconnect(context); err != nil {
		log.Panic("Unable to disconnect from the database", err)
		panic(err)
	}

	log.Println("Successfully disconnected from the Database.")
}
