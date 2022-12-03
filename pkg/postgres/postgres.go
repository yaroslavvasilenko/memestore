package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const (
	FilePath = "./store/"
)

const (
	TyText = iota
	TyAudio
	TyDocument
	TyPhoto
	TyVideo
	TyVoice
)

type File struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Size     int
	IdUser   int
	TypeFile int
	MimeType string
}

type User struct {
	ID        int `gorm:"primaryKey"`
	SizeStore int
}

func PostgresInit(urlPostgres string) (*gorm.DB, error) {
	dbURL := urlPostgres

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(User{}, File{})

	return db, err
}

// Переделать на методы для DB

func FindFile(db *gorm.DB, nameFile string, idUser int) (*File, error) {
	var result File
	tx := db.Raw(
		`SELECT id, name, size, id_user, type_file, mime_type
			 FROM files
			 WHERE id_user = ? and name = ?`, idUser, nameFile).Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

func DeleteFile(db *gorm.DB, nameFile string, idUser int) error {
	fileFinding, err := FindFile(db, nameFile, idUser)
	if err != nil {
		return err
	}
	tx := db.Delete(fileFinding)
	if tx.Error != nil {
		return tx.Error
	}
	err = DeleteFileStore(fileFinding.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFileStore(idFile string) error {
	err := os.Remove(FilePath + idFile)
	if err != nil {
		return err
	}
	return nil
}
