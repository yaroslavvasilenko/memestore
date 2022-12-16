package fileSystem

import (
	"context"
	telebot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"memestore/pkg/postgres"
)

type Video struct {
	ITypeFile
	ID       string
	Name     string
	Size     int
	IdUser   int
	MimeType string
}

func (d *Video) AnswerInlineQuery(bot *telebot.Bot, inlineQueryId, url, description string, nameFile string) error {
	inlineVideo := models.InlineQueryResultVideo{
		ID:          inlineQueryId,
		Title:       nameFile,
		VideoURL:    url,
		ThumbURL:    "https://memestore-q0oy.onrender.com/thumb_url",
		MimeType:    d.MimeType,
		Description: description,
	}

	results := []models.InlineQueryResult{
		&inlineVideo,
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

func (d *Video) GiveFile() *postgres.File {
	video := &postgres.File{
		ID:       d.ID,
		Name:     d.Name,
		Size:     d.Size,
		IdUser:   d.IdUser,
		TypeFile: postgres.TyVideo,
		MimeType: d.MimeType,
	}
	return video
}
