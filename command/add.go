package command

import (
	"strings"

	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

const filename = ".todo.json"

// Add is a command to add a todo.
var Add = cli.Command{
	Name:   "add",
	Usage:  "Add a todo",
	Action: add,
}

func add(c *cli.Context) error {
	if c.NArg() < 1 {
		cli.ShowCommandHelp(c, "add")
		return nil
	}

	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	title := strings.Join(c.Args(), " ")
	todo := todo.Todo{Title: title, Done: false}
	todos = append(todos, todo)

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}
