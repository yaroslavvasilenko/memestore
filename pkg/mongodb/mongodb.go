package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	Client     *mongo.Client
	Collection *Collection
}

type Collection struct {
	Audio *AudioCollection
	Doc   *DocCollection
}

func InitMongo() (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	db := client.Database("memestore")

	dbAndCol := &Database{
		Client: client,
		Collection: &Collection{
			Audio: &AudioCollection{handle: db.Collection("audio")},
			Doc:   &DocCollection{handle: db.Collection("document")},
		},
	}

	return dbAndCol, err
}
