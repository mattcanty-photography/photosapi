package photoapi

import (
	"image/jpeg"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nfnt/resize"
	"github.com/rs/zerolog"
)

func (c *Configuration) thumbnailHandler(w http.ResponseWriter, r *http.Request) {

	photoData := &PhotoData{
		ID: chi.URLParam(r, "ID"),
	}

	err := GetDecoderByExtension(photoData)
	if err != nil {
		log.Fatal(err)
	}

	err = c.ImageSource.GetImage(photoData)
	if err != nil {
		log.Fatal(err)
	}

	err = photoData.Decoder.Decode(photoData)
	if err != nil {
		log.Fatal(err)
	}

	photoData.Image = resize.Resize(400, 0, photoData.Image, resize.Lanczos3)

	err = jpeg.Encode(w, photoData.Image, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func notFoundPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	w.Write([]byte("Not found"))
}

func errorHandler(w http.ResponseWriter, oplog zerolog.Logger, err error) {
	w.Write([]byte(err.Error()))
}
