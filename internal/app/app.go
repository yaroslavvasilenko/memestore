package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"memestore/pkg/config"
	"memestore/pkg/logging"
	"os"
	"strconv"

	"memestore/pkg/postgres"
	"memestore/pkg/telegramapi"
)

type App struct {
	Db       *gorm.DB
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
	mdb, err := postgres.PostgresInit(cfg.PostgresURL)
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
		if update.InlineQuery != nil && update.InlineQuery.Query != "" {
			app.myInlineQuery(update)
		} else if update.Message != nil {
			app.myInsertFile(update)
		}
	}
}

func (app *App) myInlineQuery(update tgbotapi.Update) {
	f, err := postgres.FindFile(app.Db, update.InlineQuery.Query, update.InlineQuery.From.ID)
	log.Info("Find file")
	if err != nil {
		log.Debug(err, "file not found")
		return
	}

	file := makeTypeFileForDB(f)
	log.Info("Type file")
	if file == nil {
		log.Debug("no find file")
		return
	}
	idUser := strconv.Itoa(f.IdUser)

	url := fmt.Sprintf("https://memestore-q0oy.onrender.com/for_telegram?id_user=%s&id_file=%s", idUser, f.ID)

	err = file.AnswerInlineQuery(app.Bot, update.InlineQuery.ID, url, update.InlineQuery.Query)
	log.Info("Yes")
	if err != nil {
		log.Debug(err)
	}
}

func (app *App) myInsertFile(update tgbotapi.Update) {
	userID := update.Message.From.ID
	file := app.makeTypeFile(update.Message)
	if file == nil {
		log.Debug("no type file")
		return
	}

	if app.execUser(userID) == false {
		tx := app.Db.Create(&postgres.User{
			ID:        userID,
			SizeStore: 0,
		})
		if tx.Error != nil {
			log.Debug(tx.Error)
			return
		}
	}
	if err := file.DownloadFile(); err != nil {
		log.Debug(err)
		return
	}
	if err := file.InsertDB(app.Db, userID); err != nil {
		log.Debug(err)
		return
	}

	if err := app.sendMessageFast(update.Message.Chat.ID, "File downloaded"); err != nil {
		log.Debug(err)
	}

}
