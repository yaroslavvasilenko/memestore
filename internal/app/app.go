package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"memestore/pkg/config"
	"memestore/pkg/logging"
	"os"

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
	_, err := postgres.FindFile(app.Db, update.InlineQuery.Query, update.InlineQuery.From.ID)
	if err != nil {
		//  ToDO: make msg "file not found"
		return
	}

	article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Echo", "file find")
	article.Description = update.InlineQuery.Query

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       []interface{}{article},
	}

	if _, err := app.Bot.AnswerInlineQuery(inlineConf); err != nil {
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

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "File downloaded")
	_, err := app.Bot.Send(msg)
	if err != nil {
		log.Debug(err)
	}
}
