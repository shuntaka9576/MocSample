package converter

import (
	"image"
	"io"
)

type ImageType interface {
	Decode(r io.Reader) (image.Image, error)
	Encode(w io.Writer, m image.Image) error
}

var supportImageTypes = map[string]ImageType{}

func ResisterImageType(imageExt string, imageType ImageType) {
	supportImageTypes[imageExt] = imageType
}
