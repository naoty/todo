package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	// Get a path for a todo file
	dir := os.Getenv("TODO_PATH")
	if dir == "" {
		dir = os.Getenv("HOME")
	}

	path := filepath.Join(dir, filename)

	// Read data
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Failed to read data: %v", err)
	}

	// Decode data
	var todos []todo.Todo
	if err = json.Unmarshal(data, &todos); err != nil {
		return fmt.Errorf("Failed to decode data: %v", err)
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
