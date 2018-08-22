package gif

import (
	"github.com/shuntaka9576/MocSample/converter"
	"image"
	"image/gif"
	"io"
)

func init() {
	converter.ResisterImageType("gif", &Gif{})
}

type Gif struct {
}

func (*Gif) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

func (*Gif) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}
