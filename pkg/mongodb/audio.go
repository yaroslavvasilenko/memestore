package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type AudioCollection struct {
	handle *mongo.Collection
}

func (d *AudioCollection) InsertAudio(str interface{}) error {
	_, err := d.handle.InsertOne(context.Background(), str)
	if err != nil {
		return err
	}
	return nil
}
