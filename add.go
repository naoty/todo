package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var Add = cli.Command{
	Name:  "add",
	Usage: "Add a TODO",
	Action: func(context *cli.Context) {
		if len(context.Args()) == 0 {
			cli.ShowCommandHelp(context, "add")
			os.Exit(1)
		}

		title := strings.Join(context.Args(), " ")
		add := addProcess(title)
		err := UpdateTodos(add)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func addProcess(title string) process {
	return func(todos []Todo) ([]Todo, error) {
		todo := Todo{Title: title, Done: false}
		return append(todos, todo), nil
	}
}
