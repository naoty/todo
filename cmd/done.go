package cmd

import (
	"fmt"
	"strconv"

	"github.com/naoty/todo/repository"
	"github.com/naoty/todo/todo"
)

// Done represents `done` subcommand.
type Done struct {
	cli  CLI
	repo repository.Repository
}

// NewDone returns a new Done.
func NewDone(cli CLI, version string, repo repository.Repository) Command {
	return &Done{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Done) Run(args []string) int {
	if len(args) < 3 {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	td, err := c.repo.Get(id)
	td.State = todo.Done

	err = c.repo.Update(td)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	return 0
}
