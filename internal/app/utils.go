package app

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"memestore/internal/app/fileSystem"
	"memestore/pkg/postgres"
)

func (app *App) makeTypeFile(m *models.Message) fileSystem.ITypeFile {
	if m.Document != nil {
		return &fileSystem.Document{
			ID:       m.Document.FileID,
			Name:     m.Document.FileName,
			Size:     m.Document.FileSize,
			MimeType: m.Document.MimeType,
		}
	} else if m.Audio != nil {
		return &fileSystem.Audio{
			ID:       m.Audio.FileID,
			Name:     m.Caption,
			Size:     m.Audio.FileSize,
			MimeType: m.Audio.MimeType,
		}
	} else if m.Photo != nil {

		return &fileSystem.Photo{
			ID:       m.Photo[0].FileID,
			Name:     m.Caption,
			Size:     m.Photo[0].FileSize,
			MimeType: "image/png",
		}
	} else if m.Video != nil {
		return &fileSystem.Video{
			ID:       m.Video.FileID,
			Name:     m.Video.FileName,
			Size:     m.Video.FileSize,
			MimeType: m.Video.MimeType,
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
	case postgres.TyVideo:
		return &fileSystem.Video{
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

func (app *App) sendMessageFast(chatID int, textMessage string) error {
	_, err := app.Bot.SendMessage(context.TODO(), &bot.SendMessageParams{
		ChatID: chatID,
		Text:   textMessage,
	})
	if err != nil {
		return err
	}
	return nil
}

func IsCommand(m *models.Message) bool {
	if m.Entities == nil || len(m.Entities) == 0 {
		return false
	}

	return (m.Entities)[0].Offset == 0 && (m.Entities)[0].Type == "bot_command"
}
