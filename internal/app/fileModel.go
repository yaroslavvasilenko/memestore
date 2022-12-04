package app

import (
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