package app

import (
	log "github.com/sirupsen/logrus"
	"memestore/internal/app/fileSystem"
	"net/http"
	"os"
)

func (app *App) ServerForLink() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/for_telegram", app.getFile)

	log.Info("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func (app *App) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hi"))
}

func (app *App) getFile(w http.ResponseWriter, r *http.Request) {
	log.Info("Yes?")
	idFile := r.URL.Query().Get("id_file")
	idUser := r.URL.Query().Get("id_user")
	log.Println(idFile, idUser)

	f, err := os.ReadFile(fileSystem.FilePath + idFile)
	if err != nil {
		log.Debug(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	//  ToDo: Need MIME type. Possibly - https://github.com/gabriel-vasile/mimetype
	//   or write one yourself
	w.Write(f)
}
