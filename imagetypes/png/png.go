package png

import (
	"github.com/shuntaka9576/MocSample/imagetypes"
	"image"
	"image/png"
	"io"
)

func init() {
	imagetypes.ResisterImageType("jpg", &Png{})
	imagetypes.ResisterImageType("jpeg", &Png{})
}

type Png struct{}

func (*Png) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func (*Png) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
