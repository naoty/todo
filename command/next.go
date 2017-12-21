package command

import (
	"fmt"

	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

var Next = cli.Command{
	Name:   "next",
	Usage:  "Show a next undone todo",
	Action: next,
}

func next(c *cli.Context) error {
	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	todo, found := nextTodoFromTodos(todos)
	if found {
		fmt.Println(todo.Title)
	}

	return nil
}

func nextTodoFromTodos(todos []todo.Todo) (todo.Todo, bool) {
	for _, todo := range todos {
		if todo.Done {
			continue
		}

		if len(todo.Todos) > 0 {
			subtodo, found := nextTodoFromTodos(todo.Todos)
			if found {
				return subtodo, true
			}
			return todo, true
		}

		return todo, true
	}

	return todo.Todo{}, false
}
