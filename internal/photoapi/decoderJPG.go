package photoapi

import (
	image "image/jpeg"
	"log"
)

type JPGDecoder struct{}

func (_ JPGDecoder) Decode(photoData *PhotoData) error {
	log.Println("Decoding JPG")
	i, err := image.Decode(photoData.Reader)

	photoData.Image = i

	return err
}
