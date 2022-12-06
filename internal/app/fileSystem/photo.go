package fileSystem

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"memestore/pkg/postgres"
)

type Photo struct {
	ITypeFile
	ID       string
	Name     string
	Size     int
	IdUser   int
	MimeType string
}

func (p *Photo) AnswerInlineQuery(bot *tgbotapi.BotAPI, inlineQueryId, url, description string, nameFile string) error {
	inlinePhoto := tgbotapi.NewInlineQueryResultPhotoWithThumb(inlineQueryId, url, url)
	inlinePhoto.Description = description
	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: inlineQueryId,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{inlinePhoto},
	}

	if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
		return err
	}
	return nil
}

func (p *Photo) GiveFile() *postgres.File {
	photo := &postgres.File{
		ID:       p.ID,
		Name:     p.Name,
		Size:     p.Size,
		IdUser:   p.IdUser,
		TypeFile: postgres.TyPhoto,
		MimeType: p.MimeType,
	}
	return photo
}
