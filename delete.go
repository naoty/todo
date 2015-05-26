package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

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

		num, err := strconv.Atoi(context.Args()[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		delete := deleteProcess(num)
		err = UpdateTodos(delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func deleteProcess(num int) process {
	return func(todos []Todo) ([]Todo, error) {
		index := num - 1
		if index >= len(todos) {
			return nil, errors.New("Index out of bounds.")
		}

		return append(todos[:index], todos[index+1:]...), nil
	}
}
