package fileSystem

import (
	"github.com/go-telegram/bot"
	"memestore/pkg/postgres"
)

type ITypeFile interface {
	AnswerInlineQuery(bot *bot.Bot, inlineQueryId, url, description string, nameFile string) error
	GiveFile() *postgres.File
}
