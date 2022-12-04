package app

import (
	"memestore/internal/app/fileSystem"
	"memestore/pkg/postgres"
)

type instansFullFile struct {
	FileDB    *postgres.File
	FileTgAPI fileSystem.ITypeFile
}

func newFullFile(file fileSystem.ITypeFile) *instansFullFile {
	return &instansFullFile{
		FileDB:    file.GiveFile(),
		FileTgAPI: file,
	}
}
