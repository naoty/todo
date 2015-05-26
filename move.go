package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var Move = cli.Command{
	Name:  "move",
	Usage: "move a TODO",
	Action: func(context *cli.Context) {
		if len(context.Args()) < 2 {
			cli.ShowCommandHelp(context, "move")
			os.Exit(1)
		}

		nums, err := Atois(context.Args())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		move := moveProcess(nums[0], nums[1])
		err = UpdateTodos(move)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func moveProcess(from, to int) process {
	return func(todos []Todo) ([]Todo, error) {
		fromIndex, toIndex := from-1, to-1
		if fromIndex >= len(todos) || toIndex >= len(todos) {
			return nil, errors.New("Index out of bounds.")
		}

		movedTodo := todos[fromIndex]
		todos = append(todos[:fromIndex], todos[fromIndex+1:]...)
		todos = append(todos[:toIndex], append([]Todo{movedTodo}, todos[toIndex:]...)...)
		return todos, nil
	}
}
