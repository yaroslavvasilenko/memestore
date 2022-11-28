package fileSystem

import (
	"gorm.io/gorm"
	"memestore/pkg/postgres"
)

type Audio struct {
	ITypeFile
	ID     string
	Name   string
	Size   int
	IdUser int
}

func (a *Audio) DownloadFile() error {
	randName := makeRandom()
	err := downloadAny(a.ID, FilePath+randName)
	if err != nil {
		return err
	}
	a.ID = randName
	return nil
}

func (a *Audio) InsertDB(db *gorm.DB, idUser int) error {
	tx := db.Create(postgres.File{
		ID:       a.ID,
		Name:     a.Name,
		Size:     a.Size,
		IdUser:   idUser,
		TypeFile: postgres.TyAudio,
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

func (a *Audio) DeleteDB(db *gorm.DB, idUser int) error {
	tx := db.Delete(postgres.File{
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
