package fileSystem

import (
	"context"
	telebot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	memeModels "github.com/yaroslavvasilenko/meme_store_models"
)

type Photo struct {
	ITypeFile
	ID       string
	Name     string
	Size     int
	IdUser   int
	MimeType string
}

func (p *Photo) AnswerInlineQuery(bot *telebot.Bot, inlineQueryId, url, description string, nameFile string) error {
	inlinePhoto := models.InlineQueryResultPhoto{
		ID:       inlineQueryId,
		PhotoURL: url,
		ThumbURL: "https://memestore-q0oy.onrender.com/thumb_url", //ToDo: plug
	}

	results := []models.InlineQueryResult{
		&inlinePhoto,
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

func (p *Photo) GiveFile() *memeModels.File {
	photo := &memeModels.File{
		ID:       p.ID,
		Name:     p.Name,
		Size:     p.Size,
		IdUser:   p.IdUser,
		TypeFile: memeModels.TyPhoto,
		MimeType: p.MimeType,
	}
	return photo
}
