package cli

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/shuntaka9576/MocSample/converter"
	_ "github.com/shuntaka9576/MocSample/imagetypes/gif"
	_ "github.com/shuntaka9576/MocSample/imagetypes/jpg"
	_ "github.com/shuntaka9576/MocSample/imagetypes/png"
)

// TODO add user definition variables

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
	flags := flag.NewFlagSet("imgConverter", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)
	flags.StringVar(&fromExt, "f", "png", "")
	flags.StringVar(&toExt, "t", "jpg", "")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprintf(c.ErrStream, err.Error())
		return 1
	}

	nonflagArgs := flags.Args()
	switch {
	case len(nonflagArgs) <= 0:
		targetDir = "."
	case len(nonflagArgs) == 1:
		targetDir = nonflagArgs[0]
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
	var createdImageFileNames []string
	for _, path := range filepaths {
		path, err = filepath.Abs(path)
		if err != nil {
			fmt.Fprintf(c.ErrStream, err.Error())
			return 1
		}
		convertedImageName := filepath.Base(path[:len(path)-len(filepath.Ext(path))] + "_c." + toExt)

		// Duplication check
		convertedImageName = checkSameFileName(createdImageFileNames, convertedImageName, 0)

		// Convert image file
		outfilname, err := convert.Convert(path, filepath.Join(outdir, convertedImageName))
		if err != nil {
			fmt.Fprintf(c.ErrStream, err.Error())
			return 1
		}
		if outfilname == "" {
			continue
		}
		createdImageFileNames = append(createdImageFileNames, convertedImageName)
		fmt.Fprintf(c.OutStream, "Convert Succeeded![%s -> %s]\n", path, filepath.Join(outdir, convertedImageName))
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

// Returns file names that have never been created
func checkSameFileName(createdImages []string, convertedImageName string, count int) string {
	for _, imageName := range createdImages {
		if imageName == convertedImageName {
			count++
			convertedImageName = convertedImageName[:strings.LastIndex(convertedImageName, "c")+1] + strconv.Itoa(count) + filepath.Ext(convertedImageName)
			convertedImageName = checkSameFileName(createdImages, convertedImageName, count)
		}
	}
	return convertedImageName
}
