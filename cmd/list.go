package cmd

import (
	"fmt"

	"github.com/naoty/todo/repository"
)

// List represents `list` subcommand.
type List struct {
	cli  CLI
	repo repository.Repository
}

// NewList returns a new List.
func NewList(cli CLI, meta Metadata, repo repository.Repository) Command {
	return &List{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *List) Run(args []string) int {
	todos, err := c.repo.List()
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	for _, td := range todos {
		fmt.Fprintln(c.cli.Writer, td)
	}

	return 0
}
