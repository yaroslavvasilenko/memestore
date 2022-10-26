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

	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Running Application")

	for update := range *a.MessChan {
		if update.Message != nil { // If we got a message
			log.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			m, err := a.Bot.Send(msg)
			if err != nil {
				log.Info(m)
			}
		}
	}

	//a.Run(ctx)

}
