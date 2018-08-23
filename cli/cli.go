package cli

import (
	"errors"
	"flag"
	"fmt"
	"github.com/shuntaka9576/MocSample/converter"
	_ "github.com/shuntaka9576/MocSample/imagetypes/gif"
	_ "github.com/shuntaka9576/MocSample/imagetypes/jpg"
	_ "github.com/shuntaka9576/MocSample/imagetypes/png"
	"io"
	"io/ioutil"
	"path/filepath"
)

type Cli struct {
	OutStream, ErrStream io.Writer
	Name, Version        string
}

type Option struct {
	Dirpath, FromExt, ToExt string
}

func NewApp(in, out io.Writer) *Cli {
	return &Cli{OutStream: in, ErrStream: out}
}

func (c *Cli) Run(args []string) error {
	var option Option
	flag.StringVar(&option.FromExt, "f", "png", "")
	flag.StringVar(&option.ToExt, "t", "jpg", "")
	flag.Parse()

	switch {
	case len(flag.Args()) == 0:
		option.Dirpath = "."
	case len(flag.Args()) == 1:
		option.Dirpath = flag.Arg(0)
	default:
		return errors.New("dir argument error occurred")
	}
	fmt.Println(option.ToExt, option.FromExt, option.Dirpath)

	convert, err := converter.GetConverter(option.ToExt, option.FromExt)
	if err != nil {
		return err
	}

	filepaths := dirwalk(option.Dirpath)
	for _, path := range filepaths {
		path, err = filepath.Abs(path)
		if err != nil {
			return err
		}
		err := convert.Convert(path, ".")
		if err != nil {
			return err
		}
	}
	return nil
}

// Recursively directory search function
func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
