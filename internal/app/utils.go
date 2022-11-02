package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"memestore/internal/app/fileSystem"
)

func (app *App) anyDow(id string) string {
	file := tgbotapi.FileConfig{
		FileID: id,
	}

	f, _ := app.Bot.GetFile(file)
	return f.Link("5496447413:AAENVSjTJw_3Uk7CUEzoNX23XC185eY7hH8")
}

func (app *App) makeTypeDownload(m *tgbotapi.Message) fileSystem.IDowload {
	if m.Document != nil {
		return &fileSystem.Document{
			ID:   app.anyDow(m.Document.FileID),
			Name: m.Document.FileName,
			Size: m.Document.FileSize,
		}
	} else if m.Audio != nil {
		return &fileSystem.Audio{
			ID:   app.anyDow(m.Audio.FileID),
			Name: m.Audio.FileName,
			Size: m.Audio.FileSize,
		}
	}
	return nil
}