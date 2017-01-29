package command

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

const filename = ".todo"

// Add is a command to add a todo.
var Add = cli.Command{
	Name:   "add",
	Usage:  "Add a todo",
	Action: add,
}

func add(c *cli.Context) error {
	if c.NArg() < 2 {
		cli.ShowCommandHelp(c, "add")
		return nil
	}

	dir := os.Getenv("TODO_PATH")

	if dir == "" {
		dir = os.Getenv("HOME")
	}

	if dir == "" && runtime.GOOS == "windows" {
		dir = os.Getenv("APPDATA")
	}

	if dir == "" {
		return errors.New("Failed to get a directory for .todo file")
	}

	path := filepath.Join(dir, filename)

	var f *os.File

	if _, err := os.Stat(path); err != nil {
		f, err = os.Create(path)
		if err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(path, os.O_RDWR, 0644)
		if err != nil {
			return err
		}
	}

	defer f.Close()

	enc := json.NewEncoder(f)

	title := strings.Join(c.Args(), " ")
	todo := todo.Todo{Title: title, Done: false}
	err := enc.Encode(todo)

	if err != nil {
		return cli.NewExitError("Failed to encode a todo", 1)
	}

	return nil
}
