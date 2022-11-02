package fileSystem

type Document struct {
	IDowload
	ID   string
	Name string
	Size int
}

func (d *Document) DownloadFile() {

}
