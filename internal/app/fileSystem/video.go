package fileSystem

import (
	"context"
	telebot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	memeModels "github.com/yaroslavvasilenko/meme_store_models"
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
		ThumbURL:    "https://memestore-q0oy.onrender.com/thumb_url", //ToDo: plug
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

func (d *Video) GiveFile() *memeModels.File {
	video := &memeModels.File{
		ID:       d.ID,
		Name:     d.Name,
		Size:     d.Size,
		IdUser:   d.IdUser,
		TypeFile: memeModels.TyVideo,
		MimeType: d.MimeType,
	}
	return video
}
