package main

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/internal/app"
	"musicBot/pkg/logging"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logging.GetLogger(ctx)

	logger.Info("config initializing")

	ctx = logging.ContextWithLogger(ctx, logger)

	app, err := app.NewApp()
	if err != nil {
		logger.Fatal("fatal init")
	}

	for update := range *app.BUpdateChannel {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := app.Bapi.Send(msg)
			if err != nil {
				logger.Debug("deb")
				return
			}
		}
	}

}
