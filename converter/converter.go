package converter

import (
	"github.com/shuntaka9576/MocSample/image"
)



type Converter struct {
	From, To image.ImageType
}



func GetConverter(from, to, path string) {
	var converter Converter
	if val, ok := image.SupportImageTypes[from];
		converter.From = val
	}

}
