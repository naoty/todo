package cmd

import (
	"fmt"
	"strings"

	"github.com/naoty/todo/repository"
)

// Add represents `add` subcommand.
type Add struct {
	cli  CLI
	repo repository.Repository
}

// NewAdd returns a new Add.
func NewAdd(cli CLI, version string, repo repository.Repository) Command {
	return &Add{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Add) Run(args []string) int {
	if len(args) < 3 {
		fmt.Fprintf(c.cli.ErrorWriter, usage())
		return 1
	}

	title := strings.Join(args[2:], " ")
	err := c.repo.Add(title)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	return 0
}
