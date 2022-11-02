package fileSystem

import (
	"crypto/rand"
	"io"
	"math/big"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type IDowload interface {
	DownloadFile()
}

func Dowl(id string, path string) {
	resp, _ := http.Get(id)
	defer resp.Body.Close()
	out, err := os.Create(path)
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Error(err)
	}
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
