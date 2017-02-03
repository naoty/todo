package command

import (
	"strconv"

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
		order, err2 := strconv.Atoi(arg)

		if err2 != nil {
			continue
		}

		i := order - 1
		if i >= len(todos) {
			continue
		}

		todo := todos[i]
		todo.Done = false
		todos[i] = todo
	}

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}
