package main

import (
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

	myApp, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Webhook {
		log.Info("Bot webhook started")
		if err := myApp.RunWebhook(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Info("Bot polling started")
		myApp.RunLongPool()
	}
}
