package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type AudioCollection struct {
	handle *mongo.Collection
}

func NewAudioCollection(handle *mongo.Collection) *AudioCollection {
	return &AudioCollection{handle: handle}
}
