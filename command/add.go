package command

import (
	"strings"

	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

// Add is a command to add a todo.
var Add = cli.Command{
	Name:   "add",
	Usage:  "Add a todo",
	Action: add,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "parent, p",
			Usage: "Specify the parent todo",
		},
	},
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
	todo := todo.Todo{Title: title, Done: false, Todos: []todo.Todo{}}

	p := c.Int("parent")
	if p == 0 {
		todos = append(todos, todo)
	} else if len(todos) > p-1 {
		parent := todos[p-1]
		parent.Todos = append(parent.Todos, todo)
		todos[p-1] = parent
	}

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}
