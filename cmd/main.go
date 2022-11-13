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

	app, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Bot polling started")

	app.Run()
}
