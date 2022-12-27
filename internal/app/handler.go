package app

import (
	"fmt"
	"github.com/go-telegram/bot/models"
	"memestore/pkg/log"
	"strconv"
)

func (app *App) myInlineQuery(update *models.Update) {
	f, err := app.Db.FindFile(update.InlineQuery.Query, update.InlineQuery.From.ID)
	log.Info("Find file")
	if err != nil {
		log.Debug(err, "file not found")
		return
	}

	file := makeTypeFileForDB(f)
	log.Info("Type file")
	if file == nil {
		log.Debug("no find file")
		return
	}
	idUser := strconv.Itoa(f.IdUser)

	url := fmt.Sprintf(app.UrlLink+"for_telegram?id_user=%s&id_file=%s", idUser, f.ID)

	err = file.AnswerInlineQuery(app.Bot, update.InlineQuery.ID, url, update.InlineQuery.Query, f.Name)
	log.Info("Yes")
	if err != nil {
		log.Debug(err)
	}
}

func (app *App) myInsertFile(update *models.Update) {
	if update.Message.Audio != nil && update.Message.Caption == "" {
		app.sendMessageFast(update.Message.Chat.ID, "С аудио файлом сразу прописывайте его название в надписи")
		return
	}
	file := app.makeTypeFile(update.Message)
	if file == nil {
		log.Debug("no type file")
		return
	}

	f := newFullFile(file)

	f.FileDB.IdUser = update.Message.From.ID

	if app.Db.ExecUser(f.FileDB.IdUser) == false {
		err := app.Db.CreateUser(f.FileDB)
		if err != nil {
			log.Debug(err)
			return
		}
	}

	if app.Db.CheckName(f.FileDB) {
		app.sendMessageFast(update.Message.Chat.ID, "Такое имя файла занято")
		return
	}

	if err := app.Download(f.FileDB); err != nil {
		log.Debug(err)
		return
	}

	if err := app.Db.InsertDB(f.FileDB); err != nil {
		log.Debug(err)
		return
	}

	if err := app.sendMessageFast(update.Message.Chat.ID, "File downloaded"); err != nil {
		log.Debug(err)
	}

}

func (app *App) myCommand(update *models.Update) {
	var msg string
	switch update.Message.Text {
	case "/start":
		msg = `Hi v0.2
Что бы посмотреть что я могу введи /help`
	case "/help":
		msg = `Вы можете отправить мне документ(pdf), mp3 пока что
и через инлайн запрос потом отправить его любому человеку в любой момент
имя файла надо вводить полностью

@MemesStore_bot 'название файла'
Что бы удалить напишите "удалить 'имя файла'"`
	case "/files":
		files := app.Db.AllFileUser(update.Message.From.ID)
		for _, value := range files {
			msg += value.Name + "\n"
		}
		msg = "вот\n" + msg
	case "delete":

	default:
		msg = "такого не знаю"
	}
	app.sendMessageFast(update.Message.Chat.ID, msg)
}

func (app *App) deleteFileForName(arrayText []string, update *models.Update) {
	for i := 1; i < len(arrayText); i++ {
		err := app.Db.DeleteFile(arrayText[i], update.Message.From.ID)
		if err != nil {
			app.sendMessageFast(update.Message.Chat.ID, "файл "+arrayText[i]+" не удалён")
			continue
		}
		app.sendMessageFast(update.Message.Chat.ID, "удаленно "+arrayText[i])
	}

}

func (app *App) superUserCommand(arrayText []string, update *models.Update) {
	if update.Message.From.ID != 767640121 {
		app.sendMessageFast(update.Message.Chat.ID, "уходи")
		return
	}
	if arrayText[0] == "superuserdeleteall" {
		if err := app.Db.AllDelete(); err != nil {
			log.Debug(err)
			app.sendMessageFast(update.Message.Chat.ID, "что то не так")
		}
		app.sendMessageFast(update.Message.Chat.ID, "ready")
	}

}
