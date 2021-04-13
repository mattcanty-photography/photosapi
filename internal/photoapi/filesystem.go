package photoapi

import (
	"os"
	"path"
)

type FileSystemImageSource struct {
	Directory string
}

func (s FileSystemImageSource) GetImage(photoData *PhotoData) error {
	f, err := os.Open(path.Join(s.Directory, photoData.ID))
	if err != nil {
		return err
	}
	// defer f.Close()

	photoData.Reader = f

	return err
}
