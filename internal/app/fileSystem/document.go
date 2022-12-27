package fileSystem

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	memeModels "github.com/yaroslavvasilenko/meme_store_models"
)

type Document struct {
	ITypeFile
	ID       string
	Name     string
	Size     int
	IdUser   int
	MimeType string
}

func (d *Document) AnswerInlineQuery(myBot *bot.Bot, inlineQueryId, url, description string, nameFile string) error {
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

	inlineConf := &bot.AnswerInlineQueryParams{
		InlineQueryID: inlineQueryId,
		IsPersonal:    true,
		Results:       results,
	}

	if _, err := myBot.AnswerInlineQuery(context.TODO(), inlineConf); err != nil {
		return err
	}
	return nil
}

func (d *Document) GiveFile() *memeModels.File {
	doc := &memeModels.File{
		ID:       d.ID,
		Name:     d.Name,
		Size:     d.Size,
		IdUser:   d.IdUser,
		TypeFile: memeModels.TyDocument,
		MimeType: d.MimeType,
	}
	return doc
}
