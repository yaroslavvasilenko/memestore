package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"memestore/internal/app"
	"memestore/pkg/config"
)

func main() {
	cfg, err := config.GetConf()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("config initializing")

	app, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Bot polling started")

	for update := range *app.MessChan {
		if update.Message != nil { // If we got app message
			log.WithFields(log.Fields{
				"userName": update.Message.From.UserName,
				"mess":     update.Message.Text}).Info("mess user")

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			m, err := app.Bot.Send(msg)
			if err != nil {
				log.Info("%s", m)
			}
		}
	}

}
