package cmd

import (
	"fmt"

	"github.com/naoty/todo/repository"
	"github.com/spf13/pflag"
)

// Root represents a command run when no subcommand is passed.
type Root struct {
	cli     CLI
	version string
}

// NewRoot returns a *Root.
func NewRoot(cli CLI, version string, repo repository.Repository) Command {
	return &Root{cli: cli, version: version}
}

// Run implements Command interface.
func (c *Root) Run(args []string) int {
	flagset := pflag.NewFlagSet("", pflag.ExitOnError)
	help := flagset.BoolP("help", "h", false, "")
	version := flagset.BoolP("version", "v", false, "")

	flagset.Parse(args)

	if *help {
		fmt.Fprintln(c.cli.Writer, usage())
		return 0
	}

	if *version {
		fmt.Fprintln(c.cli.Writer, c.version)
		return 0
	}

	fmt.Fprintln(c.cli.ErrorWriter, usage())
	return 1
}
