package telegramapi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"memestore/pkg/config"
	"net/http"
)

func InitBot(cfg *config.Config) (*tgbotapi.BotAPI, *tgbotapi.UpdatesChannel, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TeleToken)
	if err != nil {
		return nil, nil, err
	}

	bot.Debug = cfg.Debug //  debug or no
	if cfg.Webhook != true {
		log.Info("Authorized on account %s", bot.Self.UserName)

		_, err := bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://memestore.onrender.com/"+bot.Token, "cert.pem"))
		if err != nil {
			return nil, nil, err
		}

		info, err := bot.GetWebhookInfo()
		if err != nil {
			return nil, nil, err
		}

		if info.LastErrorDate != 0 {
			log.Info("Telegram callback failed: %s", info.LastErrorMessage)
		}

		updates := bot.ListenForWebhook("/" + bot.Token)
		go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
		return bot, &updates, nil

	} else {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)
		if err != nil {
			return nil, nil, err
		}

		return bot, &updates, nil
	}

}
