package app

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"memestore/internal/app/fileSystem"
	"memestore/pkg/postgres"
)

type fileModel struct {
	FileDB    *postgres.File
	FileTgAPI fileSystem.ITypeFile
}

func newFullFile(file fileSystem.ITypeFile) *fileModel {
	return &fileModel{
		FileDB:    file.GiveFile(),
		FileTgAPI: file,
	}
}

func (app *App) Download(fileDb *postgres.File) error {
	f, _ := app.Bot.GetFile(context.TODO(), &bot.GetFileParams{
		FileID: fileDb.ID,
	})
	linkForDownload := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", app.TokenBot, f.FilePath)

	return fileDb.DownloadFile(linkForDownload)
}
