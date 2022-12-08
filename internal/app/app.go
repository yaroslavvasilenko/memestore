package app

import (
	"context"
	telebot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"memestore/pkg/config"
	"memestore/pkg/logging"
	"net/http"
	"os"
	"strings"

	"memestore/pkg/postgres"
	"memestore/pkg/telegramapi"
)

type App struct {
	Db       *postgres.DB
	Bot      *telebot.Bot
	TokenBot string
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

	app := &App{
		Db:       mdb,
		TokenBot: cfg.TeleToken,
		LogFile:  logF,
	}

	bApi, err := telegramapi.InitBot(cfg, app.handler)
	if err != nil {
		return nil, err
	}

	app.Bot = bApi

	return app, err
}

func (app *App) RunLongPool() {
	//  app.Bot.DeleteWebhook(context.TODO(), &telebot.DeleteWebhookParams{false})
	//  ToDo: No work

	app.Bot.Start(context.TODO())
}

func (app *App) RunWebhook() error {
	//  app.Bot.DeleteWebhook(context.TODO(), &telebot.DeleteWebhookParams{false})
	//  ToDo: No work

	_, err := app.Bot.SetWebhook(context.TODO(), &telebot.SetWebhookParams{
		URL: "https://memestore.onrender.com/"})
	if err != nil {
		return err
	}

	go func() {
		http.ListenAndServe(":2000", app.Bot.WebhookHandler())
	}()

	// Use StartWebhook instead of Start
	app.Bot.StartWebhook(context.TODO())

	return nil
}

func (app *App) handler(ctx context.Context, b *telebot.Bot, update *models.Update) {
	if update.InlineQuery != nil {
		app.myInlineQuery(update)
	} else if update.Message != nil && IsCommand(update.Message) {
		app.myCommand(update)
	} else if update.Message != nil {
		testSplit := strings.Split(update.Message.Text, " ")
		testSplit[0] = strings.ToLower(testSplit[0])
		if testSplit[0] == "удалить" || testSplit[0] == "delete" {
			app.deleteFileForName(testSplit, update)
		} else if testSplit[0] == strings.ToLower("superUserDeleteAll") {
			app.superUserCommand(testSplit, update)
		} else {
			app.myInsertFile(update)
		}
	}
}
