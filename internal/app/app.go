package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"musicBot/pkg/mongodb"
	"musicBot/pkg/telegramapi"
)

type App struct {
	Mdb            *mongo.Client
	Bapi           *tgbotapi.BotAPI
	BUpdateChannel *tgbotapi.UpdatesChannel
}

func NewApp() (*App, error) {
	mdb, err := mongodb.InitMongo()
	bApi, bUpdateChannel := telegramapi.InitBot()

	return &App{Mdb: mdb,
		Bapi:           bApi,
		BUpdateChannel: bUpdateChannel}, err
}
