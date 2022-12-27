package fileSystem

import (
	"github.com/go-telegram/bot"
	memeModels "github.com/yaroslavvasilenko/meme_store_models"
)

type ITypeFile interface {
	AnswerInlineQuery(myBot *bot.Bot, inlineQueryId, url, description string, nameFile string) error
	GiveFile() *memeModels.File
}
