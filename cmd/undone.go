package cmd

import (
	"fmt"
	"strconv"

	"github.com/naoty/todo/repository"
	"github.com/naoty/todo/todo"
)

// Undone represents `undone` subcommand.
type Undone struct {
	cli  CLI
	repo repository.Repository
}

// NewUndone returns a new Undone.
func NewUndone(cli CLI, version string, repo repository.Repository) Command {
	return &Undone{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Undone) Run(args []string) int {
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
	td.State = todo.Undone

	err = c.repo.Update(td)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	return 0
}
