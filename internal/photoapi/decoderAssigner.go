package photoapi

import (
	"errors"
	"fmt"
	"log"
	"path"
	"strings"
)

func GetDecoderByExtension(photoData *PhotoData) error {
	extension := strings.TrimSpace(strings.ToLower(path.Ext(photoData.ID)))
	log.Printf("Extension: '%s'", extension)
	switch extension {
	case ".jpg":
		photoData.Decoder = &JPGDecoder{}
	case ".jpeg":
		photoData.Decoder = &JPGDecoder{}
	case ".cr2":
		photoData.Decoder = &CR2Decoder{}
	default:
		return errors.New(fmt.Sprintf("Extension '%s' not recognised.", extension))
	}
	return nil
}
