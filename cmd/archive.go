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
		err := c.runArchive(td)
		if err != nil {
			fmt.Fprintln(c.cli.ErrorWriter, err)
			return 1
		}
	}

	return 0
}

func (c *Archive) runArchive(td *todo.Todo) error {
	if td.State == todo.Done {
		err := c.repo.Archive(td.ID)
		if err != nil {
			return err
		}
	}

	for _, sub := range td.Todos {
		err := c.runArchive(sub)
		if err != nil {
			return err
		}
	}

	return nil
}
