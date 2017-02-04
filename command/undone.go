package command

import (
	"strconv"
	"strings"

	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

// Undone is a command to mark a todo as undone.
var Undone = cli.Command{
	Name:   "undone",
	Usage:  "Mark todos as undone",
	Action: undone,
}

func undone(c *cli.Context) error {
	if c.NArg() == 0 {
		cli.ShowCommandHelp(c, "undone")
		return nil
	}

	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	for _, arg := range c.Args() {
		var orders []int
		for _, id := range strings.Split(arg, "-") {
			order, err2 := strconv.Atoi(id)
			if err2 == nil {
				orders = append(orders, order)
			}
		}

		todos = undoneTodos(todos, orders)
	}

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}

func undoneTodos(todos []todo.Todo, orders []int) []todo.Todo {
	if len(orders) == 0 {
		return todos
	}

	i := orders[0] - 1
	if i >= len(todos) {
		return todos
	}

	todo := todos[i]
	if len(todo.Todos) == 0 || len(orders) == 1 {
		todo.Done = false
	} else {
		todo.Todos = undoneTodos(todo.Todos, orders[1:])
	}
	todos[i] = todo

	return todos
}
