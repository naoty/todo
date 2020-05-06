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

// Metadata represents metadata of this application.
type Metadata struct {
	Version string
}

// CommandFactory represents a factory function for a command.
type CommandFactory func(cli CLI, meta Metadata, repo repository.Repository) Command

var commandFactories = map[string]CommandFactory{
	"add":  NewAdd,
	"list": NewList,
	"open": NewOpen,
	"move": NewMove,
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
  todo add <title>
  todo list
  todo open <id>
  todo move <id> <position>
  todo -h | --help
  todo -v | --version

Options:
  -h --help     Show help message
  -v --version  Show version
`

	return strings.Trim(message, "\n")
}
