package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type DocCollection struct {
	handle *mongo.Collection
}

func NewDocCollection(handle *mongo.Collection) *DocCollection {
	return &DocCollection{handle: handle}
}
