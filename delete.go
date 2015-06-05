package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var Delete = cli.Command{
	Name:  "delete",
	Usage: "Delete a TODO",
	Action: func(context *cli.Context) {
		if len(context.Args()) == 0 {
			cli.ShowCommandHelp(context, "delete")
			os.Exit(1)
		}

		// num, err := strconv.Atoi(context.Args()[0])
		nums, err := Atois(context.Args())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		delete := deleteProcess(nums...)
		err = UpdateTodos(delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func deleteProcess(nums ...int) process {
	return func(todos []Todo) ([]Todo, error) {
		var indices []int
		var err error

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

		var newTodos []Todo
		for i, todo := range todos {
			if !Contains(indices, i) {
				newTodos = append(newTodos, todo)
			}
		}

		return newTodos, nil
	}
}
