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

	// Append a todo
	title := strings.Join(c.Args(), " ")
	todo := todo.Todo{Title: title, Done: false}
	todos = append(todos, todo)

	// Encode data
	indent := strings.Repeat(" ", 4)
	json, err := json.MarshalIndent(todos, "", indent)
	if err != nil {
		return fmt.Errorf("Failed to encode data: %v", err)
	}

	// Write json
	ioutil.WriteFile(path, json, 0644)

	return nil
}
