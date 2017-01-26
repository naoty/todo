package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Usage = "manage todos"
	app.Run(os.Args)
}
