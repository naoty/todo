package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var Undone = cli.Command{
	Name:  "undone",
	Usage: "Undone TODOs",
	Action: func(context *cli.Context) {
		if len(context.Args()) == 0 {
			cli.ShowCommandHelp(context, "undone")
			os.Exit(1)
		}

		nums, err := Atois(context.Args())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		undone := undoneProcess(nums...)
		err = UpdateTodos(undone)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func undoneProcess(nums ...int) process {
	return func(todos []Todo) ([]Todo, error) {
		var err error
		var indices []int

		for _, num := range nums {
			index := num - 1
			if index >= len(todos) {
				err = errors.New("Index out of bounds.")
			}
			indices = append(indices, index)
		}
		if err != nil {
			return nil, err
		}

		newTodos := make([]Todo, len(todos))
		for i, todo := range todos {
			if Contains(indices, i) {
				todo.Done = false
			}
			newTodos[i] = todo
		}
		return newTodos, nil
	}
}
