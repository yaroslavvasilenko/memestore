package postgres

import (
	"io"
	"net/http"
	"os"
)

func (f *File) DownloadFile(linkDowload string) error {

	return downloadAny(linkDowload, FilePath+f.ID)
}

func downloadAny(id string, path string) error {
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
