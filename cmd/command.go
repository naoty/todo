package cmd

import (
	"strings"

	"github.com/naoty/todo/repository"
)

// Command represents an interface for all commands.
type Command interface {
	Run(args []string) int
}

// CommandFactory represents a factory function for a command.
type CommandFactory func(cli CLI, config Config, repo repository.Repository) Command

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
