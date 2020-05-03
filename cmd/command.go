package cmd

import (
	"io"
	"strings"
)

// Command represents an interface for all commands.
type Command interface {
	Run(args []string) int
}

// CLI represents an I/O against CLI.
type CLI struct {
	Version     string
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}

// CommandFactory represents a factory function for a command.
type CommandFactory func(cli CLI) Command

var commandFactories = map[string]CommandFactory{
	"list": NewList,
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
  todo -h | --help
  todo -v | --version

Options:
  -h --help     Show help message
  -v --version  Show version
`

	return strings.Trim(message, "\n")
}
