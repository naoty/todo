package cmd

import (
	"fmt"
	"strconv"

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
	if len(args) < 3 {
		fmt.Fprintf(c.cli.ErrorWriter, usage())
		return 1
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	td, err := c.repo.Get(id)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	td.State = todo.Archived
	err = c.repo.Update(td)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	return 0
}
