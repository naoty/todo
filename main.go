package main

import (
	"os"

	"github.com/naoty/todo/command"
	"github.com/urfave/cli"
)

// Version is the version of this application.
var Version = "0.2.0"

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Version = Version
	app.Usage = "Manage todos"
	app.Author = "Naoto Kaneko"
	app.Email = "naoty.k@gmail.com"
	app.Commands = []cli.Command{
		command.Add,
		command.List,
		command.Done,
	}
	app.Run(os.Args)
}
