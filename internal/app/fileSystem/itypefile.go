package fileSystem

import (
	"crypto/rand"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"math/big"
	"net/http"
	"os"

	"gorm.io/gorm"
)

const (
	filePath = "./store/"
)

type ITypeFile interface {
	DownloadFile() error
	InsertDB(db *gorm.DB, idUser int) error
	DeleteDB(db *gorm.DB, idUser int) error
	AnswerInlineQuery(bot *tgbotapi.BotAPI, inlineQueryId, url, description string) error
}

func downloadAny(id string, path string) error {
	resp, err := http.Get(id)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func makeRandom() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	idRune := make([]rune, 16)
	for i := range idRune {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			nBig = big.NewInt(0)

		}
		idRune[i] = letterRunes[nBig.Int64()]
	}
	return string(idRune)
}
