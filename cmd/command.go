package cmd

import (
	"io"
	"strings"

	"github.com/naoty/todo/repository"
)

// Command represents an interface for all commands.
type Command interface {
	Run(args []string) int
}

// CLI represents an I/O against CLI.
type CLI struct {
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}

// CommandFactory represents a factory function for a command.
type CommandFactory func(cli CLI, version string, repo repository.Repository) Command

var commandFactories = map[string]CommandFactory{
	"list":    NewList,
	"add":     NewAdd,
	"delete":  NewDelete,
	"move":    NewMove,
	"open":    NewOpen,
	"done":    NewDone,
	"undone":  NewUndone,
	"wait":    NewWait,
	"archive": NewArchive,
}

// Lookup returns a CommandFactory based on args.
func Lookup(args []string) CommandFactory {
	if len(args) < 2 {
		return NewRoot
	}

	factory, ok := commandFactories[args[1]]
	if !ok {
		return NewRoot
	}

	return factory
}

func usage() string {
	message := `
Usage:
  todo list
  todo add (-p | --parent <id>) <title>
  todo delete <id>
  todo move <id> <position>
  todo open <id>
  todo done <id>
  todo undone <id>
  todo wait <id>
  todo archive <id>
  todo -h | --help
  todo -v | --version

Options:
  -h --help     Show help message
  -v --version  Show version
`

	return strings.Trim(message, "\n")
}
