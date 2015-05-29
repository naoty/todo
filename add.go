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
		add := addProcess(title, context.Bool("once"))
		err := UpdateTodos(add)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "once, o",
			Usage: "Add the TODO only if it exists",
		},
	},
}

func addProcess(title string, once bool) process {
	return func(todos []Todo) ([]Todo, error) {
		if once && todoExist(todos, title) {
			return todos, nil
		}

		todo := Todo{Title: title, Done: false}
		return append(todos, todo), nil
	}
}

func todoExist(todos []Todo, title string) bool {
	for _, todo := range todos {
		if todo.Title == title {
			return true
		}
	}
	return false
}
