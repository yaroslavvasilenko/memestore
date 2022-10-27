package telegramapi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"memestore/pkg/config"
)

func InitBot(cfg *config.Config) (*tgbotapi.BotAPI, *tgbotapi.UpdatesChannel, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TeleToken) //  env
	if err != nil {
		return nil, nil, err
	}

	bot.Debug = cfg.Debug //  debag or no

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	return bot, &updates, err
}
