package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"memestore/pkg/config"
	"memestore/pkg/logging"
	"os"
	"strconv"
	"strings"

	"memestore/pkg/postgres"
	"memestore/pkg/telegramapi"
)

type App struct {
	Db       *postgres.DB
	Bot      *tgbotapi.BotAPI
	TokenBot string
	MessChan *tgbotapi.UpdatesChannel
	LogFile  *os.File
}

func NewApp(cfg *config.Config) (*App, error) {
	logF, err := logging.InitLog(cfg)
	if err != nil {
		return nil, err
	}
	mdb, err := postgres.PostgresInit(cfg.PostgresURL)
	if err != nil {
		return nil, err
	}

	bApi, mesChan, err := telegramapi.InitBot(cfg)
	if err != nil {
		return nil, err
	}
	log.Printf("Authorized on account %s", bApi.Self.UserName)

	app := &App{
		Db:       mdb,
		Bot:      bApi,
		TokenBot: cfg.TeleToken,
		MessChan: mesChan,
		LogFile:  logF,
	}

	return app, err
}

func (app *App) Run() {
	for update := range *app.MessChan {
		if update.InlineQuery != nil && update.InlineQuery.Query != "" {
			app.myInlineQuery(update)
		} else if update.Message.IsCommand() {
			app.myCommand(update)
		} else if update.Message != nil {
			testSplit := strings.Split(update.Message.Text, " ")
			testSplit[0] = strings.ToLower(testSplit[0])
			if testSplit[0] == "удалить" || testSplit[0] == "delete" {
				app.deleteFileForName(testSplit, update)
			} else if testSplit[0] == strings.ToLower("superUserDeleteAll)") {
				app.superUserCommand(testSplit, update)
			} else {
				app.myInsertFile(update)
			}
		}
	}
}

func (app *App) myInlineQuery(update tgbotapi.Update) {
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

	url := fmt.Sprintf("https://memestore-q0oy.onrender.com/for_telegram?id_user=%s&id_file=%s", idUser, f.ID)

	err = file.AnswerInlineQuery(app.Bot, update.InlineQuery.ID, url, update.InlineQuery.Query)
	log.Info("Yes")
	if err != nil {
		log.Debug(err)
	}
}

func (app *App) myInsertFile(update tgbotapi.Update) {
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

	if err := f.FileDB.DownloadFile(); err != nil {
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

func (app *App) myCommand(update tgbotapi.Update) {
	var msg string
	switch update.Message.Command() {
	case "start":
		msg = "Hi"
	case "help":
		msg = `Вы можете отправить мне только документ(pdf) пока что
и через инлайн запрос потом отправить его любому человеку в любой момент
имя файла надо вводить полностью
@MemesStore_bot название файла`
	case "files":
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

func (app *App) deleteFileForName(arrayText []string, update tgbotapi.Update) {
	for i := 1; i < len(arrayText); i++ {
		err := app.Db.DeleteFile(arrayText[i], update.Message.From.ID)
		if err != nil {
			app.sendMessageFast(update.Message.Chat.ID, "файл "+arrayText[i]+" не удалён")
			continue
		}
		app.sendMessageFast(update.Message.Chat.ID, "удаленно "+arrayText[i])
	}

}

func (app *App) superUserCommand(arrayText []string, update tgbotapi.Update) {
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
