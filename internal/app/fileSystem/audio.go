package fileSystem

import (
	"gorm.io/gorm"
	"memestore/pkg/mongodb"
)

type Audio struct {
	ITypeFile
	ID   string
	Name string
	Size int
}

func (a *Audio) DownloadFile() error {
	randName := makeRandom()
	err := downloadAny(a.ID, audioPath+randName)
	if err != nil {
		return err
	}
	a.ID = randName
	return nil
}

func (a *Audio) InsertDB(db *gorm.DB, idUser int64) error {
	tx := db.Create(mongodb.Audio{
		ID:     a.ID,
		Name:   a.Name,
		Size:   a.Size,
		IdUser: idUser,
	})
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.Exec(
		`UPDATE users
			SET size_store = size_store + ? 
			WHERE id = ?`, a.Size, idUser)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (a *Audio) DeleteDB(db *gorm.DB, idUser int64) error {
	tx := db.Delete(mongodb.Document{
		ID:     a.ID,
		Name:   a.Name,
		Size:   a.Size,
		IdUser: idUser,
	})
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.Exec(
		`UPDATE users 
			SET size_store = size_store - ? 
			WHERE id = ?`, a.Size, idUser)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
