package jpg

import (
	"github.com/shuntaka9576/MocSample/imagetypes"
	"image"
	"image/jpeg"
	"io"
)

func init() {
	imagetypes.ResisterImageType("jpg", &Jpeg{})
}

type Jpeg struct{}

func (*Jpeg) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

func (*Jpeg) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}