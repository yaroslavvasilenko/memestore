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

	bot.Debug = cfg.Debug //  debug or no

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	return bot, &updates, err
}
