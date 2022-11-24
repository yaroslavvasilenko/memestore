package telegramapi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"memestore/pkg/config"
	"net/http"
	"time"
)

func InitBot(cfg *config.Config) (*tgbotapi.BotAPI, *tgbotapi.UpdatesChannel, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TeleToken)
	if err != nil {
		return nil, nil, err
	}

	bot.RemoveWebhook()
	time.Sleep(time.Second * 5)
	bot.Debug = cfg.Debug //  debug or no
	if cfg.Webhook == true {
		log.Info("Start on webhook")

		_, err := bot.SetWebhook(tgbotapi.NewWebhook("https://memestore.onrender.com:443/" + bot.Token))
		if err != nil {
			return nil, nil, err
		}

		info, err := bot.GetWebhookInfo()
		log.Info(info)
		if err != nil {
			return nil, nil, err
		}

		if info.LastErrorDate != 0 {
			log.Info("Telegram callback failed: %s", info.LastErrorMessage)
		}

		updates := bot.ListenForWebhook("/" + bot.Token)
		go http.ListenAndServe("0.0.0.0:443", nil)
		return bot, &updates, nil

	} else {
		log.Info("Start on longpoll")
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)
		if err != nil {
			return nil, nil, err
		}

		return bot, &updates, nil
	}

}
