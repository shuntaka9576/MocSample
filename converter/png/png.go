package png

import (
	"github.com/shuntaka9576/MocSample/converter"
	"image"
	"image/png"
	"io"
)

func init() {
	converter.ResisterImageType("jpg", &Png{})
}

type Png struct{}

func (*Png) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func (*Png) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
