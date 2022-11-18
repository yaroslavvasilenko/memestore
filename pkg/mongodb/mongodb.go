package mongodb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Audio struct {
	ID     string `gorm:"primaryKey"`
	Name   string
	Size   int
	IdUser int64
}

type Document struct {
	ID     string `gorm:"primaryKey"`
	Name   string
	Size   int
	IdUser int64
}

type Video struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Size int
}

type Text struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Size int
}

type Photo struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Size int
}

type Voise struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Size int
}

type Music struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Size int
}

type User struct {
	ID        int64 `gorm:"primaryKey"`
	SizeStore int
}

func InitMongo() (*gorm.DB, error) {
	dbURL := "postgres://pg:pass@172.17.0.3:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(User{}, Audio{}, Document{})

	return db, err
}
