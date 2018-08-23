package converter

import (
	"github.com/shuntaka9576/MocSample/imagetypes"
	"os"
)

type Converter struct {
	From, To imagetypes.ImageType
}

func GetConverter(from, to string) (converter Converter, err error) {
	converter.From, err = imagetypes.CheckSupportImageType(from)
	if err != nil {
		return converter, err
	}

	converter.To, err = imagetypes.CheckSupportImageType(to)
	if err != nil {
		return converter, err
	}
	return converter, nil
}

func (c *Converter) Convert(inputImagePath, outputPath string) error {
	// decode
	file, err := os.Open(inputImagePath)
	if err != nil {
		return err
	}
	defer file.Close()
	decodeImage, err := c.From.Decode(file)
	if err != nil {
		return err
	}

	// encode
	outfile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	err = c.To.Encode(outfile, decodeImage)
	if err != nil {
		return err
	}

	return nil
}
