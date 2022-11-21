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
}

type User struct {
	ID        int `gorm:"primaryKey"`
	SizeStore int
}

func InitMongo() (*gorm.DB, error) {
	dbURL := "postgres://pg:pass@172.17.0.3:5432"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(User{}, File{})

	return db, err
}

func FindFile(db *gorm.DB, name string, idUser int) (string, int, error) {
	var result File
	tx := db.Raw(
		`SELECT name
			 FROM files
			 WHERE id_user = ? and name = ?`, idUser, name).Scan(&result)
	if tx.Error != nil {
		return "", 0, tx.Error
	}
	return name, idUser, nil
}
