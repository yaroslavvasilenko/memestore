package main

import (
	"memestore/internal/app"
	"memestore/pkg/config"
	"memestore/pkg/log"
)

func main() {
	pathLog := log.InitLog()
	log.Log().Info(pathLog)
	log.Log().Info("config initializing")
	cfg, err := config.GetConf()
	if err != nil {
		log.Log().Panic(err)
	}

	myApp, err := app.NewApp(cfg)
	if err != nil {
		log.Log().Panic(err)
	}

	if cfg.Webhook {
		log.Log().Info("Bot webhook started")
		if err := myApp.RunWebhook(); err != nil {
			log.Log().Panic(err)
		}
	} else {
		log.Log().Info("Bot polling started")
		myApp.RunLongPool()
	}
}
