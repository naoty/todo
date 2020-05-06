package cmd

import (
	"fmt"

	"github.com/naoty/todo/repository"
	"github.com/naoty/todo/todo"
)

// List represents `list` subcommand.
type List struct {
	cli  CLI
	repo repository.Repository
}

// NewList returns a new List.
func NewList(cli CLI, version string, repo repository.Repository) Command {
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
		var mark string
		switch td.State {
		case todo.Undone:
			mark = "[ ]"
		case todo.Done:
			mark = "[x]"
		case todo.Waiting:
			mark = "[w]"
		case todo.Archived:
			mark = "[-]"
		}

		fmt.Fprintf(c.cli.Writer, "%s %03d: %s\n", mark, td.ID, td.Title)
	}

	return 0
}
