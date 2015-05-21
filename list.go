package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var List = cli.Command{
	Name:   "list",
	Usage:  "List TODOs",
	Action: list,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "markdown, m",
			Usage: "Print TODOs as task lists in markdown",
		},
	},
}

func list(context *cli.Context) {
	todos, err := ReadTodos()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// var formatter TodoFormatter = StandardFormatter{Out: os.Stdout, Err: os.Stderr}

	var formatter TodoFormatter
	if context.Bool("markdown") {
		formatter = MarkdownFormatter{Out: os.Stdout, Err: os.Stderr}
	} else {
		formatter = StandardFormatter{Out: os.Stdout, Err: os.Stderr}
	}

	for _, todo := range todos {
		formatter.Println(todo)
	}
}
