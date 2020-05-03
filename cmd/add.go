package cmd

import "github.com/naoty/todo/repository"

// Add represents `add` subcommand.
type Add struct {
}

// NewAdd returns a new Add.
func NewAdd(cli CLI, meta Metadata, repo repository.Repository) Command {
	return &Add{}
}

// Run implements Command interface.
func (c *Add) Run(args []string) int {
	return 0
}
