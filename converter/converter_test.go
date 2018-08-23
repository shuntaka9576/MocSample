package converter_test

import (
	"github.com/shuntaka9576/MocSample/imagetypes"
	"github.com/shuntaka9576/MocSample/imagetypes/gif"
	"github.com/shuntaka9576/MocSample/imagetypes/png"
	"github.com/shuntaka9576/MocSample/imagetypes/jpg"
	"testing"
)

func TestGetConverter(t *testing.T) {
	//tests := []struct {
	//	name string
	//	input input
	//}
	imagetypes.ResisterImageType("png", &png.Png{})
	imagetypes.ResisterImageType("jpg", &jpg.Jpeg{})
	imagetypes.ResisterImageType("jpeg", &jpg.Jpeg{})
	imagetypes.ResisterImageType("gif", &gif.Gif{})
}
