package command

import (
	"fmt"

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

	for _, todo := range todos {
		if todo.Done {
			continue
		}
		fmt.Printf("%s\n", todo.Title)
		return nil
	}

	return nil
}
