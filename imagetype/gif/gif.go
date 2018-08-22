package gif

import (
	"github.com/shuntaka9576/MocSample/imagetype"
	"image"
	"image/gif"
	"io"
)

func init() {
	imagetype.ResisterImageType("gif", &Gif{})
}

type Gif struct{}

func (*Gif) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

func (*Gif) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}
