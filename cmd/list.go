package cmd

import (
	"fmt"

	"github.com/naoty/todo/repository"
	"github.com/naoty/todo/todo"
)

// List represents `list` subcommand.
type List struct {
	cli    CLI
	config Config
	repo   repository.Repository
}

// NewList returns a new List.
func NewList(cli CLI, config Config, repo repository.Repository) Command {
	return &List{cli: cli, config: config, repo: repo}
}

// Run implements Command interface.
func (c *List) Run(args []string) int {
	todos := []*todo.Todo{
		todo.New(1, "dummy"),
		todo.New(2, "dummy"),
	}

	for _, td := range todos {
		fmt.Fprintln(c.cli.Writer, td)
	}

	return 0
}
