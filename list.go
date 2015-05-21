package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

const (
	ENOENT = 2
)

var List = cli.Command{
	Name:   "list",
	Usage:  "list TODOs",
	Action: list,
}

func list(context *cli.Context) {
	todos, err := ReadTodos()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ENOENT)
	}

	var formatter TodoFormatter = StandardFormatter{Out: os.Stdout, Err: os.Stderr}
	for _, todo := range todos {
		formatter.Println(todo)
	}
}
