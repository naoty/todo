package cmd

import "github.com/naoty/todo/repository"

// Open represents `open` subcommand.
type Open struct {
	cli CLI
}

// NewOpen returns a new Open.
func NewOpen(cli CLI, meta Metadata, repo repository.Repository) Command {
	return &Open{cli: cli}
}

// Run implements Command interface.
func (c *Open) Run(args []string) int {
	return 0
}
