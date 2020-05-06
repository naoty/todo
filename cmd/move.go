package cmd

import (
	"fmt"
	"strconv"

	"github.com/naoty/todo/repository"
)

// Move represents `move` subcommand.
type Move struct {
	cli  CLI
	repo repository.Repository
}

// NewMove returns a new Move.
func NewMove(cli CLI, meta Metadata, repo repository.Repository) Command {
	return &Move{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Move) Run(args []string) int {
	if len(args) < 4 {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	position, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	if position < 1 {
		fmt.Fprintf(c.cli.ErrorWriter, "position number must be larger than 0: %d", position)
		return 1
	}

	err = c.repo.Move(id, position)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	return 0
}
