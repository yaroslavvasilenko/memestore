package app

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"memestore/pkg/config"
	"memestore/pkg/logging"

	"memestore/pkg/mongodb"
	"memestore/pkg/telegramapi"
)

type App struct {
	Mdb      *mongo.Client
	Bot      *tgbotapi.BotAPI
	MessChan *tgbotapi.UpdatesChannel
	LogFile  *os.File
}

func NewApp(cfg *config.Config) (*App, error) {
	logF, err := logging.InitLog(cfg)
	if err != nil {
		return nil, err
	}
	mdb, err := mongodb.InitMongo()
	if err != nil {
		return nil, err
	}

	bApi, mesChan, err := telegramapi.InitBot(cfg)
	if err != nil {
		return nil, err
	}
	log.Printf("Authorized on account %s", bApi.Self.UserName)

	app := &App{
		Mdb:      mdb,
		Bot:      bApi,
		MessChan: mesChan,
		LogFile:  logF,
	}

	return app, err
}
