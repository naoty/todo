package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var Add = cli.Command{
	Name:   "add",
	Usage:  "Add a TODO",
	Action: add,
}

func add(context *cli.Context) {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "add")
		os.Exit(1)
	}
	newTitle := strings.Join(context.Args(), " ")

	todo := Todo{Title: newTitle, Done: false}
	err := AppendTodo(todo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
