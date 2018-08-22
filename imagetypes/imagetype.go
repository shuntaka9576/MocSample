package imagetype

import (
	"github.com/pkg/errors"
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

func CheckSupportImageType(extension string) (imagetype ImageType, err error) {
	imagetype, ok := supportImageTypes[extension]
	if !ok {
		return imagetype, errors.New("not found option:" + extension)
	}
	return imagetype, nil
}
