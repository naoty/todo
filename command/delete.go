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
		todos = trashTodos(todos, orders)
	}
	todos = cleanTrashedTodos(todos)

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}

func trashTodos(todos []todo.Todo, orders []int) []todo.Todo {
	if len(orders) == 0 {
		return todos
	}

	i := orders[0] - 1
	if i > len(todos)-1 {
		return todos
	}

	todo := todos[i]
	if len(todo.Todos) == 0 || len(orders) == 1 {
		todo.Trashed = true
	} else {
		todo.Todos = trashTodos(todo.Todos, orders[1:])
	}
	todos[i] = todo

	return todos
}

func cleanTrashedTodos(todos []todo.Todo) []todo.Todo {
	newTodos := []todo.Todo{}

	for _, todo := range todos {
		if todo.Trashed {
			continue
		}

		todo.Todos = cleanTrashedTodos(todo.Todos)
		newTodos = append(newTodos, todo)
	}

	return newTodos
}
