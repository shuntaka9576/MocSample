package main

import (
	"github.com/shuntaka9576/MocSample/cli"
	"os"
	"fmt"
)

const (
	Version = "v0.1.0"
	Name    = "imageConverter"
)

func main() {
	err := newApp().Run(os.Args)
	if err != nil{
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func newApp() *cli.Cli {
	app := cli.NewApp(os.Stdin, os.Stderr)
	app.Version = Version
	app.Name = Name
	return app
}
