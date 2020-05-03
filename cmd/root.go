package cmd

import (
	"fmt"

	"github.com/spf13/pflag"
)

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
	flagset := pflag.NewFlagSet("", pflag.ExitOnError)
	help := flagset.BoolP("help", "h", false, "")

	flagset.Parse(args)

	if *help {
		fmt.Fprintln(c.Writer, usage())
		return 0
	}

	return 0
}
