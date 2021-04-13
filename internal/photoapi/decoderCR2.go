package photoapi

import (
	"log"

	"github.com/nf/cr2"
)

type CR2Decoder struct{}

func (_ CR2Decoder) Decode(photoData *PhotoData) error {
	log.Println("Decoding CR2")
	i, err := cr2.Decode(photoData.Reader)

	photoData.Image = i

	return err
}
