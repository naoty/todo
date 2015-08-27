package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Author = "Naoto Kaneko"
	app.Email = "naoty.k@gmail.com"
	app.Version = "0.1.0"
	app.Usage = "Manage TODOs"
	app.Commands = []cli.Command{
		command.List,
		command.Add,
		command.Delete,
		command.Move,
		command.Done,
		command.Undone,
		command.Clear,
		command.Rename,
	}
	app.Run(os.Args)
}
