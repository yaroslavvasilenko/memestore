package fileSystem

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Photo struct {
	ITypeFile
	ID       string
	Name     string
	Size     int
	IdUser   int
	MimeType string
}

func (d *Photo) AnswerInlineQuery(bot *tgbotapi.BotAPI, inlineQueryId, url, description string, nameFile string) error {
	inlineDocument := tgbotapi.NewInlineQueryResultPhoto(inlineQueryId, url)
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
