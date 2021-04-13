package photoapi

import (
	"image"
	"io"
)

type PhotoData struct {
	ID      string
	Reader  io.Reader
	Image   image.Image
	Decoder ImageDecoder
}
