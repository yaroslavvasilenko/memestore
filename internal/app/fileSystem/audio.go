package fileSystem

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"memestore/pkg/postgres"
)

type Audio struct {
	ITypeFile
	ID       string
	Name     string
	Size     int
	IdUser   int
	MimeType string
}

func (a *Audio) AnswerInlineQuery(bot *tgbotapi.BotAPI, inlineQueryId, url, description string, nameFile string) error {
	inlineAudio := tgbotapi.NewInlineQueryResultAudio(inlineQueryId, url, nameFile)
	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: inlineQueryId,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{inlineAudio},
	}

	if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
		return err
	}
	return nil
}

func (a *Audio) GiveFile() *postgres.File {
	audio := &postgres.File{
		ID:       a.ID,
		Name:     a.Name,
		Size:     a.Size,
		IdUser:   a.IdUser,
		TypeFile: postgres.TyAudio,
		MimeType: a.MimeType,
	}
	return audio
}
