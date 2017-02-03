package command

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

// List is a command to list todos
var List = cli.Command{
	Name:   "list",
	Usage:  "List todos",
	Action: list,
}

func list(c *cli.Context) error {
	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	// Print todos
	buf := []string{}
	for i, todo := range todos {
		var mark string
		if todo.Done {
			mark = "[x]"
		} else {
			mark = "[ ]"
		}
		buf = append(buf, fmt.Sprintf("%s %03d: %s", mark, i+1, todo.Title))
	}
	fmt.Printf("%v\n", strings.Join(buf, "\n"))

	return nil
}
