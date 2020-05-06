package cmd

import (
	"fmt"
	"strconv"

	"github.com/naoty/todo/repository"
)

// Open represents `open` subcommand.
type Open struct {
	cli  CLI
	repo repository.Repository
}

// NewOpen returns a new Open.
func NewOpen(cli CLI, version string, repo repository.Repository) Command {
	return &Open{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Open) Run(args []string) int {
	if len(args) < 3 {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	err = c.repo.Open(id)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	return 0
}
