package fileSystem

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"memestore/pkg/postgres"
)

type Document struct {
	ITypeFile
	ID       string
	Name     string
	Size     int
	IdUser   int
	MimeType string
}

func (d *Document) DownloadFile() error {
	randName := makeRandom()
	err := downloadAny(d.ID, FilePath+randName)
	if err != nil {
		return err
	}
	d.ID = randName
	return nil
}

func (d *Document) AnswerInlineQuery(bot *tgbotapi.BotAPI, inlineQueryId, url, description string) error {
	inlineDocument := tgbotapi.NewInlineQueryResultDocument(inlineQueryId, url, "Your document", d.MimeType)
	inlineDocument.Description = description
	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: inlineQueryId,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{inlineDocument},
	}

	if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
		return err
	}
	return nil
}

func (d *Document) InsertDB(db *gorm.DB, idUser int) error {
	tx := db.Create(postgres.File{
		ID:       d.ID,
		Name:     d.Name,
		Size:     d.Size,
		IdUser:   idUser,
		TypeFile: postgres.TyDocument,
		MimeType: d.MimeType,
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

func (d *Document) DeleteDB(db *gorm.DB, idUser int) error {
	tx := db.Delete(postgres.File{
		ID:       d.ID,
		Name:     d.Name,
		Size:     d.Size,
		IdUser:   idUser,
		TypeFile: postgres.TyDocument,
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
