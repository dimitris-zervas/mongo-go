package mongo

import (
	"context"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect(context context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(context, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	log.Info("Successfully connected to the Database!")
	return client, err
}

func Disconnect(context context.Context, client *mongo.Client) {
	if err := client.Disconnect(context); err != nil {
		log.Panic("Unable to disconnect from the database", err)
		panic(err)
	}

	log.Info("Successfully disconnected from the Database.")
}

func Ping(client *mongo.Client) {
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Panic(err)
	}
	log.Println("Pinged the database")
}
