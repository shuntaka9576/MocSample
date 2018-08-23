package cli

import (
	"flag"
	"fmt"
	"github.com/shuntaka9576/MocSample/converter"
	_ "github.com/shuntaka9576/MocSample/imagetypes/gif"
	_ "github.com/shuntaka9576/MocSample/imagetypes/jpg"
	_ "github.com/shuntaka9576/MocSample/imagetypes/png"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Cli struct {
	OutStream, ErrStream io.Writer
	Name, Version        string
}

func NewApp(in, out io.Writer) *Cli {
	return &Cli{OutStream: in, ErrStream: out}
}

func (c *Cli) Run(args []string) int {
	outdir, err := initDir()
	if err != nil {
		fmt.Fprintf(c.ErrStream, err.Error())
		return 1
	}

	var fromExt, toExt, targetDir string
	flag.StringVar(&fromExt, "f", "png", "")
	flag.StringVar(&toExt, "t", "jpg", "")
	flag.Parse()

	switch {
	case len(flag.Args()) == 0:
		targetDir = "."
	case len(flag.Args()) == 1:
		targetDir = flag.Arg(0)
	default:
		fmt.Fprintf(c.ErrStream, "dir argument error occurred\n")
		return 1
	}

	convert, err := converter.GetConverter(fromExt, toExt)
	if err != nil {
		fmt.Fprintf(c.ErrStream, err.Error())
		return 1
	}

	filepaths := dirwalk(targetDir)
	for _, path := range filepaths {
		path, err = filepath.Abs(path)
		if err != nil {
			fmt.Fprintf(c.ErrStream, err.Error())
			return 1
		}

		var createdImageFileNames map[string]bool
		if "."+fromExt == filepath.Ext(path) {
			convertedImageName := filepath.Join(outdir, filepath.Base(path[:len(path)-len(filepath.Ext(path))]+"_c."+toExt))
			if checkImageFileName(createdImageFileNames, convertedImageName, 0) {
				err := convert.Convert(path, convertedImageName)
			}
			if err != nil {
				fmt.Fprintf(c.ErrStream, err.Error())
				return 1
			}
		}
	}
	return 0
}

func initDir() (string, error) {
	outdir := "converted_" + time.Now().Format("20060102-150405")
	if err := os.Mkdir(outdir, 0777); err != nil {
	}
	initialDir, err := filepath.Abs(outdir)
	if err != nil {
		return initialDir, err
	}
	return initialDir, nil
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

func checkImageFileName(createdImages map[string]bool, convertedImageName string, count int) bool {
	switch count {
	case 0:
		_, ok := createdImages[convertedImageName]
		if !ok {
		}
		return ok

	}

}
