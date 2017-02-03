package command

import (
	"strconv"

	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

// Delete is a command to delete todos.
var Delete = cli.Command{
	Name:   "delete",
	Usage:  "Delete todos",
	Action: delete,
}

func delete(c *cli.Context) error {
	if c.NArg() == 0 {
		cli.ShowCommandHelp(c, "delete")
		return nil
	}

	orders := make([]int, len(c.Args()))
	for _, arg := range c.Args() {
		order, err := strconv.Atoi(arg)
		if err == nil {
			orders = append(orders, order)
		}
	}

	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	var newTodos []todo.Todo
	for i, todo := range todos {
		order := i + 1
		if !contains(orders, order) {
			newTodos = append(newTodos, todo)
		}
	}

	err = writeTodos(newTodos, path)
	if err != nil {
		return err
	}

	return nil
}

func contains(ns []int, n int) bool {
	for _, e := range ns {
		if e == n {
			return true
		}
	}
	return false
}
