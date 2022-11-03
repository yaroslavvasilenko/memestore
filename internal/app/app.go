package app

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"memestore/pkg/config"
	"memestore/pkg/logging"

	"memestore/pkg/mongodb"
	"memestore/pkg/telegramapi"
)

type App struct {
	Db       *mongodb.Database
	Bot      *tgbotapi.BotAPI
	TokenBot string
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
		Db:       mdb,
		Bot:      bApi,
		TokenBot: cfg.TeleToken,
		MessChan: mesChan,
		LogFile:  logF,
	}

	return app, err
}

func (app *App) Run() {
	for update := range *app.MessChan {
		if update.Message.Text != "" { // If we got app message
			log.WithFields(log.Fields{
				"userName": update.Message.From.UserName,
				"mess":     update.Message.Text}).Info("mess user")

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			msg.ReplyToMessageID = update.Message.MessageID

			m, err := app.Bot.Send(msg)
			if err != nil {
				log.Info("%s", m)
			}
		} else {
			file := app.makeTypeFile(update.Message)
			err := file.DownloadFile()
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				app.Bot.Send(msg)
				log.Debug(err)
			}
			err = file.InsertDB(app.Db.Collection)
			if err != nil {
				log.Debug(err)
			}
		}
	}
}
