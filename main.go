package main

import (
	"github.com/shuntaka9576/MocSample/cli"
	"os"
)

const (
	Version = "v0.1.0"
	Name    = "imageConverter"
)

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.Cli {
	app := cli.NewApp(os.Stdin, os.Stderr)
	app.Version = Version
	app.Name = Name
	return app
}
