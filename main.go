package main

import (
	"memestore/internal/app"
	"memestore/pkg/config"
	"memestore/pkg/log"
)

func main() {
	pathLog := log.InitLog()
	log.Info(pathLog)
	log.Info("config initializing")
	cfg, err := config.GetConf()
	if err != nil {
		log.Panic(err)
	}

	myApp, err := app.NewApp(cfg)
	if err != nil {
		log.Panic(err)
	}

	if cfg.Webhook {
		log.Info("Bot webhook started")
		if err := myApp.RunWebhook(); err != nil {
			log.Panic(err)
		}
	} else {
		log.Info("Bot polling started")
		myApp.RunLongPool()
	}
}
