package command

import (
	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

// Clear is a command to clear done todos.
var Clear = cli.Command{
	Name:   "clear",
	Usage:  "Clear done todos",
	Action: clear,
}

func clear(c *cli.Context) error {
	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	todos = clearTodos(todos)

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}

func clearTodos(todos []todo.Todo) []todo.Todo {
	var newTodos []todo.Todo

	for _, todo := range todos {
		if todo.Done {
			continue
		}

		todo.Todos = clearTodos(todo.Todos)
		newTodos = append(newTodos, todo)
	}

	return newTodos
}
