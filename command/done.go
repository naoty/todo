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
	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	if c.NArg() == 0 {
		todos, _ = doneNextTodoFromTodos(todos)
	} else {
		for _, arg := range c.Args() {
			orders := splitOrder(arg)
			todos = doneTodos(todos, orders)
		}
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

func doneNextTodoFromTodos(todos []todo.Todo) ([]todo.Todo, bool) {
	done := false

	for i, todo := range todos {
		if todo.Done {
			continue
		}

		if len(todo.Todos) > 0 {
			todo.Todos, done = doneNextTodoFromTodos(todo.Todos)
			if !done {
				todo.Done = true
				done = true
			}
		} else {
			todo.Done = true
			done = true
		}
		todos[i] = todo

		if done {
			break
		}
	}

	return todos, done
}
