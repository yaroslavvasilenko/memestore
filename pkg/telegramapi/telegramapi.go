package telegramapi

import (
	"context"
	telebot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"memestore/pkg/config"
)

func InitBot(cfg *config.Config, asd func(ctx context.Context, bot *telebot.Bot, update *models.Update)) (*telebot.Bot, error) {
	opts := []telebot.Option{
		telebot.WithDefaultHandler(asd),
	}
	if cfg.Debug {
		opts = append(opts, telebot.WithDebug())
	}

	return telebot.New(cfg.TeleToken, opts...)
}
