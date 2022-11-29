package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
