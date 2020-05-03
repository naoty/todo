package cmd

import "fmt"

// List represents `list` subcommand.
type List struct {
	cli    CLI
	config Config
}

// NewList returns a new List.
func NewList(cli CLI, config Config) Command {
	return &List{cli: cli, config: config}
}

// Run implements Command interface.
func (c *List) Run(args []string) int {
	fmt.Fprintln(c.cli.Writer, "TODO: implement list")
	return 0
}
