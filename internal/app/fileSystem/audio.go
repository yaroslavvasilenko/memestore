package fileSystem

import (
	"memestore/pkg/mongodb"
)

type Audio struct {
	IFile
	ID   string
	Name string
	Size int
}

func (a *Audio) DownloadFile() error {
	err := dowl(a.ID, audioPath+makeRandom())
	if err != nil {
		return err
	}
	return nil
}

func (a *Audio) InsertDB(m *mongodb.Collection) error {
	audio := struct {
		Id   string
		Name string
		Size int
	}{
		a.ID,
		a.Name,
		a.Size,
	}
	err := m.Audio.InsertAudio(audio)
	if err != nil {
		return err
	}
	return nil
}
