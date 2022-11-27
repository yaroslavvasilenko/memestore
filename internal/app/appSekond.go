package app

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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
	idFile := r.URL.Query().Get("id_file")
	idUser, err := strconv.Atoi(r.URL.Query().Get("id_user"))
	if err != nil {
		//  ToDO: err
		return
	}
	log.Println(idFile, idUser)
	//f, err := postgres.FindFile(app.Db, idFile, idUser)
	//if err != nil {
	//	// ToDo: err
	//}
	//  просто напиши как отдавать файл

}
