package command

import (
	"strconv"
	"strings"

	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

// Rename is a command to rename a todo.
var Rename = cli.Command{
	Name:   "rename",
	Usage:  "Rename a todo",
	Action: rename,
}

func rename(c *cli.Context) error {
	if c.NArg() < 2 {
		cli.ShowCommandHelp(c, "rename")
		return nil
	}

	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	var orders []int
	for _, id := range strings.Split(c.Args()[0], "-") {
		if order, err2 := strconv.Atoi(id); err2 == nil {
			orders = append(orders, order)
		}
	}

	title := strings.Join(c.Args()[1:], " ")
	todos = renameTodos(todos, title, orders)

	err = writeTodos(todos, path)
	return err
}

func renameTodos(todos []todo.Todo, title string, orders []int) []todo.Todo {
	if len(todos) == 0 {
		return todos
	}

	i := orders[0] - 1
	if i >= len(todos) {
		return todos
	}

	todo := todos[i]
	if len(orders) == 1 {
		todo.Title = title
	} else {
		todo.Todos = renameTodos(todo.Todos, title, orders[1:])
	}
	todos[i] = todo

	return todos
}
