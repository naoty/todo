package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Author = "Naoto Kaneko"
	app.Email = "naoty.k@gmail.com"
	app.Version = "0.1.0"
	app.Usage = "Manage TODOs"
	app.Commands = []cli.Command{List}
	app.Run(os.Args)
}
