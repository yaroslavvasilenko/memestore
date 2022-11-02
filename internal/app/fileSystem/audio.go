package fileSystem

type Audio struct {
	IDowload
	ID   string
	Name string
	Size int
}

func (d *Audio) DownloadFile() {

}
