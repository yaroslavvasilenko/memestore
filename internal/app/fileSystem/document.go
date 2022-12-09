package fileSystem

import (
	"context"
	telebot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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

func (d *Document) AnswerInlineQuery(bot *telebot.Bot, inlineQueryId, url, description string, nameFile string) error {
	inlineDocument := models.InlineQueryResultDocument{
		ID:          inlineQueryId,
		Title:       nameFile,
		DocumentURL: url,
		MimeType:    d.MimeType,
		Description: description,
	}

	results := []models.InlineQueryResult{
		&inlineDocument,
	}

	inlineConf := &telebot.AnswerInlineQueryParams{
		InlineQueryID: inlineQueryId,
		IsPersonal:    true,
		Results:       results,
	}

	if _, err := bot.AnswerInlineQuery(context.TODO(), inlineConf); err != nil {
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
