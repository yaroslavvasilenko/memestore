package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"memestore/internal/app/fileSystem"
	"memestore/pkg/postgres"
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
			ID:       app.linkForDownload(m.Document.FileID),
			Name:     m.Document.FileName,
			Size:     m.Document.FileSize,
			MimeType: m.Document.MimeType,
		}
	} else if m.Audio != nil {
		return &fileSystem.Audio{
			ID:       app.linkForDownload(m.Audio.FileID),
			Name:     m.Caption,
			Size:     m.Audio.FileSize,
			MimeType: m.Audio.MimeType,
		}
	} else if m.Photo != nil {
		ph := *m.Photo
		phOne := ph[1]
		return &fileSystem.Photo{
			ID:       app.linkForDownload(phOne.FileID),
			Name:     m.Caption,
			Size:     phOne.FileSize,
			MimeType: "image/png",
		}
	}
	return nil
}

func makeTypeFileForDB(file *postgres.File) fileSystem.ITypeFile {
	switch file.TypeFile {
	case postgres.TyDocument:
		return &fileSystem.Document{
			ID:       file.ID,
			Name:     file.Name,
			Size:     file.Size,
			IdUser:   file.IdUser,
			MimeType: file.MimeType,
		}

	case postgres.TyAudio:
		return &fileSystem.Audio{
			ID:       file.ID,
			Name:     file.Name,
			Size:     file.Size,
			IdUser:   file.IdUser,
			MimeType: file.MimeType,
		}
	case postgres.TyPhoto:
		return &fileSystem.Photo{
			ID:       file.ID,
			Name:     file.Name,
			Size:     file.Size,
			IdUser:   file.IdUser,
			MimeType: file.MimeType,
		}

	default:
		return nil

	}

}

func (app *App) sendMessageFast(chatID int64, textMessage string) error {
	msg := tgbotapi.NewMessage(chatID, textMessage)
	_, err := app.Bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
