package command

import (
	"fmt"
	"strings"

	"github.com/naoty/todo/todo"
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
	buf := formatTodos(todos, "")
	fmt.Printf("%v\n", strings.Join(buf, "\n"))

	return nil
}

func formatTodos(todos []todo.Todo, indent string) []string {
	buf := []string{}
	for i, todo := range todos {
		var mark string
		if todo.Done {
			mark = "[x]"
		} else {
			mark = "[ ]"
		}
		buf = append(buf, fmt.Sprintf("%s%s %03d: %s", indent, mark, i+1, todo.Title))

		extraIndent := strings.Repeat(" ", 2)
		subtodos := formatTodos(todo.Todos, indent+extraIndent)
		buf = append(buf, subtodos...)
	}
	return buf
}
