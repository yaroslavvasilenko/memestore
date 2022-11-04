package fileSystem

import (
	"memestore/pkg/mongodb"
)

type Document struct {
	IFile
	ID     string
	DBName string
	Name   string
	Size   int
}

func (d *Document) DownloadFile() error {
	randName := makeRandom()
	err := downloadAny(d.ID, documentPath+randName)
	if err != nil {
		return err
	}
	d.DBName = randName
	return nil
}

func (d *Document) InsertDB(m *mongodb.Collection) error {
	doc := struct {
		Id   string `bson:"id_file"`
		Name string `bson:"name"`
		Size int    `bson:"size"`
	}{
		d.ID,
		d.Name,
		d.Size,
	}
	err := m.Doc.InsertDoc(doc)
	if err != nil {
		return err
	}
	return nil
}
