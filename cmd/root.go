package cmd

import "fmt"

// Root represents a command run when no subcommand is passed.
type Root struct {
	CLI
}

// NewRoot returns a *Root.
func NewRoot(cli CLI) Command {
	return &Root{cli}
}

// Run implements Command interface.
func (c *Root) Run(args []string) int {
	fmt.Fprintln(c.Writer, "TODO: implement root")
	return 0
}
