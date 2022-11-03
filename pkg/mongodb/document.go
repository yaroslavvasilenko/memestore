package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DocCollection struct {
	handle *mongo.Collection
}

func (d *DocCollection) InsertDoc(str interface{}) error {
	_, err := d.handle.InsertOne(context.Background(), str)
	if err != nil {
		return err
	}
	return nil
}
