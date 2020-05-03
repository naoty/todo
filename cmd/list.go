package cmd

import "fmt"

// List represents `list` subcommand.
type List struct {
	CLI
}

// NewList returns a new List.
func NewList(cli CLI) Command {
	return &List{cli}
}

// Run implements Command interface.
func (c *List) Run(args []string) int {
	fmt.Fprintln(c.Writer, "TODO: implement list")
	return 0
}
