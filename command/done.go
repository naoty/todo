package command

import (
	"strconv"

	"github.com/urfave/cli"
)

// Done is a command to mark a todo as done.
var Done = cli.Command{
	Name:   "done",
	Usage:  "Mark a todo as done",
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
		order, err2 := strconv.Atoi(arg)

		if err2 != nil {
			continue
		}

		i := order - 1
		if i >= len(todos) {
			continue
		}

		todo := todos[i]
		todo.Done = true
		todos[i] = todo
	}

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}
