package main

import (
	"os"

	"github.com/urfave/cli"
)

// Version is the version of this application.
var Version = "0.2.0"

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Version = Version
	app.Usage = "manage todos"
	app.Run(os.Args)
}
