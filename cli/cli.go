package cli

import (
	"errors"
	"flag"
	"github.com/shuntaka9576/MocSample/converter"
	_ "github.com/shuntaka9576/MocSample/imagetype/gif"
	_ "github.com/shuntaka9576/MocSample/imagetype/jpg"
	_ "github.com/shuntaka9576/MocSample/imagetype/png"
	"io"
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

func (*Cli) Run(args []string) error {
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
	converter, err := converter.GetConverter(option.ToExt, option.FromExt, option.Dirpath)
	if err != nil {
		return err
	}

	converter.Convert("")

	return nil
}
