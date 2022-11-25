package main

import (
	log "github.com/sirupsen/logrus"
	"memestore/internal/app"
	"memestore/pkg/config"
	"net/http"
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

	log.Info("Bot polling started")

	go serverForLink()

	myApp.Run()
}

func serverForLink() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hi"))
}
