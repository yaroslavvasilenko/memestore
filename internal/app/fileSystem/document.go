package fileSystem

import (
	"gorm.io/gorm"
	"memestore/pkg/mongodb"
)

type Document struct {
	ITypeFile
	ID   string
	Name string
	Size int
}

func (d *Document) DownloadFile() error {
	randName := makeRandom()
	err := downloadAny(d.ID, documentPath+randName)
	if err != nil {
		return err
	}
	d.ID = randName
	return nil
}

func (d *Document) InsertDB(db *gorm.DB, idUser int64) error {
	tx := db.Create(mongodb.Document{
		ID:     d.ID,
		Name:   d.Name,
		Size:   d.Size,
		IdUser: idUser,
	})
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.Exec(
		`UPDATE users 
			SET size_store = size_store + ? 
			WHERE id = ?`, d.Size, idUser)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (d *Document) DeleteDB(db *gorm.DB, idUser int64) error {
	tx := db.Delete(mongodb.Document{
		ID:     d.ID,
		Name:   d.Name,
		Size:   d.Size,
		IdUser: idUser,
	})
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.Exec(
		`UPDATE users 
			SET size_store = size_store - ? 
			WHERE id = ?`, d.Size, idUser)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
