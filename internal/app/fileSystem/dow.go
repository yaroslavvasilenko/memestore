package fileSystem

import (
	"crypto/rand"
	"io"
	"math/big"
	"net/http"
	"os"

	"memestore/pkg/mongodb"
)

const (
	documentPath = "./store/document/"
	audioPath    = "./store/audio/"
)

type IFile interface {
	DownloadFile() error
	InsertDB(m *mongodb.Collection) error
}

func dowl(id string, path string) error {
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
