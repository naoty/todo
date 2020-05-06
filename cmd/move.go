package cmd

import (
	"fmt"

	"github.com/naoty/todo/repository"
)

// Move represents `move` subcommand.
type Move struct {
	cli CLI
}

// NewMove returns a new Move.
func NewMove(cli CLI, meta Metadata, repo repository.Repository) Command {
	return &Move{}
}

// Run implements Command interface.
func (c *Move) Run(args []string) int {
	if len(args) < 4 {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	return 0
}
