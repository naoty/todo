package cmd

import (
	"fmt"
	"strconv"

	"github.com/naoty/todo/repository"
)

// Delete represents `delete` subcommand.
type Delete struct {
	cli  CLI
	repo repository.Repository
}

// NewDelete returns a new Delete.
func NewDelete(cli CLI, version string, repo repository.Repository) Command {
	return &Delete{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Delete) Run(args []string) int {
	if len(args) < 3 {
		fmt.Fprint(c.cli.ErrorWriter, usage())
		return 1
	}

	for _, arg := range args[2:] {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(c.cli.ErrorWriter, err)
			return 1
		}

		err = c.repo.Delete(id)
		if err != nil {
			fmt.Fprintln(c.cli.ErrorWriter, err)
			return 1
		}
	}

	return 0
}
