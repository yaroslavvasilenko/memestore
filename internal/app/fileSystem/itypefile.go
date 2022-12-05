package fileSystem

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"memestore/pkg/postgres"
)

type ITypeFile interface {
	AnswerInlineQuery(bot *tgbotapi.BotAPI, inlineQueryId, url, description string, nameFile string) error
	GiveFile() *postgres.File
}
