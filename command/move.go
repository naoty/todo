package command

import (
	"errors"
	"strconv"

	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

// Move is a command to move a todo.
var Move = cli.Command{
	Name:   "move",
	Usage:  "Move a command",
	Action: move,
}

func move(c *cli.Context) error {
	if c.NArg() < 2 {
		cli.ShowCommandHelp(c, "move")
		return nil
	}

	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	from, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		return err
	}

	to, err := strconv.Atoi(c.Args()[1])
	if err != nil {
		return err
	}

	fromIndex, toIndex := from-1, to-1
	if fromIndex >= len(todos) || toIndex >= len(todos) {
		return errors.New("Index out of bounds")
	}

	t := todos[fromIndex]
	todos = append(todos[:fromIndex], todos[fromIndex+1:]...)
	todos = append(todos[:toIndex], append([]todo.Todo{t}, todos[toIndex:]...)...)

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}
