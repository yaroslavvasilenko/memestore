package fileSystem

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

func (d *Document) AnswerInlineQuery(bot *tgbotapi.BotAPI, inlineQueryId, url, description string, nameFile string) error {
	inlineDocument := tgbotapi.NewInlineQueryResultDocument(inlineQueryId, url, nameFile, d.MimeType)
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

func (d *Document) GiveFile() *postgres.File {
	doc := &postgres.File{
		ID:       d.ID,
		Name:     d.Name,
		Size:     d.Size,
		IdUser:   d.IdUser,
		TypeFile: postgres.TyDocument,
		MimeType: d.MimeType,
	}
	return doc
}
