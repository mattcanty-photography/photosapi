package photoapi

type ImageSource interface {
	GetImage(photoData *PhotoData) error
}

type ImageDecoder interface {
	Decode(photoData *PhotoData) error
}
