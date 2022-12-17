package fileSystem

import (
	"context"
	telebot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	memeModels "github.com/yaroslavvasilenko/meme_store_models"
)

type Audio struct {
	ITypeFile
	ID       string
	Name     string
	Size     int
	IdUser   int
	MimeType string
}

func (a *Audio) AnswerInlineQuery(bot *telebot.Bot, inlineQueryId, url, description string, nameFile string) error {
	inlineAudio := models.InlineQueryResultAudio{
		ID:       inlineQueryId,
		Title:    nameFile,
		AudioURL: url,
	}

	results := []models.InlineQueryResult{
		&inlineAudio,
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

func (a *Audio) GiveFile() *memeModels.File {
	audio := &memeModels.File{
		ID:       a.ID,
		Name:     a.Name,
		Size:     a.Size,
		IdUser:   a.IdUser,
		TypeFile: memeModels.TyAudio,
		MimeType: a.MimeType,
	}
	return audio
}
