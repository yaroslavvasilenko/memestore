package telegramapi

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitBot() (*tgbotapi.BotAPI, *tgbotapi.UpdatesChannel) {
	bot, err := tgbotapi.NewBotAPI("5496447413:AAENVSjTJw_3Uk7CUEzoNX23XC185eY7hH8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	return bot, &updates
}
