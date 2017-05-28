package command

import (
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

	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	for _, arg := range c.Args() {
		orders := splitOrder(arg)
		todos = deleteTodo(todos, orders)
	}

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}

func deleteTodo(todos []todo.Todo, orders []int) []todo.Todo {
	if len(orders) == 0 || len(todos) == 0 {
		return todos
	}

	i := orders[0] - 1
	if i >= len(todos) {
		return todos
	}

	if len(orders) == 1 {
		if len(todos) == 1 {
			return []todo.Todo{}
		}

		if i+1 >= len(todos) {
			return todos[:i]
		}

		return append(todos[:i], todos[i+1:]...)
	}

	todo := todos[i]
	todo.Todos = deleteTodo(todo.Todos, orders[1:])
	todos[i] = todo

	return todos
}
