package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
)

var Rename = cli.Command{
	Name:  "rename",
	Usage: "Rename a TODO",
	Action: func(context *cli.Context) {
		if len(context.Args()) < 2 {
			cli.ShowCommandHelp(context, "rename")
			os.Exit(1)
		}

		num, err := strconv.Atoi(context.Args()[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		title := strings.Join(context.Args()[1:], " ")
		rename := renameProcess(num, title)
		err = UpdateTodos(rename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func renameProcess(num int, title string) process {
	return func(todos []Todo) ([]Todo, error) {
		index := num - 1
		if index >= len(todos) {
			return nil, errors.New("Index out of bounds.")
		}

		todos[index].Title = title
		return todos, nil
	}
}
