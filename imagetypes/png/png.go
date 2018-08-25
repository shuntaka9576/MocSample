package png

import (
	"github.com/shuntaka9576/MocSample/imagetypes"
	"image"
	"image/png"
	"io"
)

func init() {
	init := &Png{[]string{".png"}}
	imagetypes.ResisterImageType(init)
}

type Png struct{
	extStrs []string
}

func (*Png) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func (*Png) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

func (g *Png) CheckExtStr(checkExt string) bool {
	for _, ext := range g.extStrs {
		if ext == checkExt {
			return true
			break
		}
	}
	return false
}
