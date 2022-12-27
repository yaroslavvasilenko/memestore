package app

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"memestore/pkg/config"
	"net/http"
	"strings"

	memeModels "github.com/yaroslavvasilenko/meme_store_models"
	"memestore/pkg/telegramapi"
)

type App struct {
	Db       *memeModels.DB
	Bot      *bot.Bot
	TokenBot string
	UrlLink  string
}

func NewApp(cfg *config.Config) (*App, error) {
	mdb, err := memeModels.PostgresInit(cfg.PostgresURL)
	if err != nil {
		return nil, err
	}

	app := &App{
		Db:       mdb,
		TokenBot: cfg.TeleToken,
	}

	bApi, err := telegramapi.InitBot(cfg, app.handler)
	if err != nil {
		return nil, err
	}

	app.Bot = bApi
	app.UrlLink = cfg.UrlLink

	return app, err
}

func (app *App) RunLongPool() {
	//  app.Bot.DeleteWebhook(context.TODO(), &bot.DeleteWebhookParams{false})
	//  ToDo: No work

	app.Bot.Start(context.TODO())
}

func (app *App) RunWebhook() error {
	//  app.Bot.DeleteWebhook(context.TODO(), &bot.DeleteWebhookParams{false})
	//  ToDo: No work

	_, err := app.Bot.SetWebhook(context.TODO(), &bot.SetWebhookParams{
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

func (app *App) handler(ctx context.Context, b *bot.Bot, update *models.Update) {
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
