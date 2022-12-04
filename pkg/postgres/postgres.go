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

type DB struct {
	Postgres *gorm.DB
}

func PostgresInit(urlPostgres string) (*DB, error) {
	dbURL := urlPostgres

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(User{}, File{})

	return &DB{db}, err
}

// Переделать на методы для DB

func (db *DB) FindFile(nameFile string, idUser int) (*File, error) {
	var result File
	tx := db.Postgres.Raw(
		`SELECT id, name, size, id_user, type_file, mime_type
			 FROM files
			 WHERE id_user = ? and name = ?`, idUser, nameFile).Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

func (db *DB) DeleteDB(f *File) error {
	tx := db.Postgres.Delete(f)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Postgres.Exec(
		`UPDATE users 
			SET size_store = size_store - ? 
			WHERE id = ?`, f.Size, f.IdUser)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (db *DB) DeleteFile(name string, idUser int) error {
	fileFinding, err := db.FindFile(name, idUser)
	if err != nil {
		return err
	}
	tx := db.Postgres.Delete(fileFinding)
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

func (db *DB) InsertDB(f *File) error {
	tx := db.Postgres.Create(f)
	if tx.Error != nil {
		return tx.Error
	}

	tx = db.Postgres.Exec(
		`UPDATE users 
			SET size_store = size_store + ? 
			WHERE id = ?`, f.Size, f.IdUser)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (db *DB) CreateUser(f *File) error {
	tx := db.Postgres.Create(&User{
		ID:        f.IdUser,
		SizeStore: 0,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (db *DB) AllFileUser(idUser int) []File {
	var files []File
	db.Postgres.Where(&File{IdUser: idUser}).Find(&files)
	return files

}

func (db *DB) ExecUser(userID int) bool {
	user := User{ID: userID}
	tx := db.Postgres.First(&user)
	if tx.RowsAffected != 1 {
		return false
	}
	return true
}

func (db *DB) CheckName(file *File) bool {

	f := &File{
		Name:   file.Name,
		IdUser: file.IdUser,
	}
	resultFile := &File{Size: 0}
	db.Postgres.Model(f).First(resultFile)

	if resultFile.Size != 0 {
		return true
	}
	return false
}
