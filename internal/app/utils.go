package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"memestore/internal/app/fileSystem"
)

func (app *App) linkForDownload(id string) string {
	file := tgbotapi.FileConfig{
		FileID: id,
	}

	f, _ := app.Bot.GetFile(file)
	return f.Link(app.TokenBot)
}

func (app *App) makeTypeFile(m *tgbotapi.Message) fileSystem.ITypeFile {
	if m.Document != nil {
		return &fileSystem.Document{
			ID:   app.linkForDownload(m.Document.FileID),
			Name: m.Document.FileName,
			Size: m.Document.FileSize,
		}
	} else if m.Audio != nil {
		return &fileSystem.Audio{
			ID:   app.linkForDownload(m.Audio.FileID),
			Name: m.Audio.FileName,
			Size: m.Audio.FileSize,
		}
	}
	return nil
}
