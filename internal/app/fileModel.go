package app

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	memeModels "github.com/yaroslavvasilenko/meme_store_models"
	"memestore/internal/app/fileSystem"
)

type fileModel struct {
	FileDB    *memeModels.File
	FileTgAPI fileSystem.ITypeFile
}

func newFullFile(file fileSystem.ITypeFile) *fileModel {
	return &fileModel{
		FileDB:    file.GiveFile(),
		FileTgAPI: file,
	}
}

func (app *App) Download(fileDb *memeModels.File) error {
	f, _ := app.Bot.GetFile(context.TODO(), &bot.GetFileParams{
		FileID: fileDb.ID,
	})
	linkForDownload := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", app.TokenBot, f.FilePath)

	return fileDb.DownloadFile(linkForDownload)
}
