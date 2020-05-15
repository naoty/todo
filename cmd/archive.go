package cmd

import (
	"fmt"

	"github.com/naoty/todo/repository"
	"github.com/naoty/todo/todo"
)

// Archive represents `archive` subcommand.
type Archive struct {
	cli  CLI
	repo repository.Repository
}

// NewArchive returns a new Archive.
func NewArchive(cli CLI, version string, repo repository.Repository) Command {
	return &Archive{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Archive) Run(args []string) int {
	todos, err := c.repo.List()
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	for _, td := range todos {
		if td.State != todo.Done {
			continue
		}

		err := c.repo.Archive(td.ID)
		if err != nil {
			fmt.Fprintln(c.cli.ErrorWriter, err)
			return 1
		}
	}

	return 0
}
