package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var Clear = cli.Command{
	Name:  "clear",
	Usage: "Clear done TODOs",
	Action: func(context *cli.Context) {
		err := UpdateTodos(clear)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func clear(todos []Todo) ([]Todo, error) {
	var newTodos []Todo
	for _, todo := range todos {
		if !todo.Done {
			newTodos = append(newTodos, todo)
		}
	}
	return newTodos, nil
}
