package command

import (
	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

// Done is a command to mark a todo as done.
var Done = cli.Command{
	Name:   "done",
	Usage:  "Mark todos as done",
	Action: done,
}

func done(c *cli.Context) error {
	if c.NArg() == 0 {
		cli.ShowCommandHelp(c, "done")
		return nil
	}

	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	for _, arg := range c.Args() {
		orders := splitOrder(arg)
		todos = doneTodos(todos, orders)
	}

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}

func doneTodos(todos []todo.Todo, orders []int) []todo.Todo {
	if len(orders) == 0 {
		return todos
	}

	i := orders[0] - 1
	if i >= len(todos) {
		return todos
	}

	todo := todos[i]
	if len(todo.Todos) == 0 || len(orders) == 1 {
		todo.Done = true
	} else {
		todo.Todos = doneTodos(todo.Todos, orders[1:])
	}
	todos[i] = todo

	return todos
}
