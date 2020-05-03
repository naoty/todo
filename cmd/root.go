package cmd

import (
	"fmt"

	"github.com/naoty/todo/repository"
	"github.com/spf13/pflag"
)

// Root represents a command run when no subcommand is passed.
type Root struct {
	cli  CLI
	meta Metadata
}

// NewRoot returns a *Root.
func NewRoot(cli CLI, meta Metadata, repo repository.Repository) Command {
	return &Root{cli: cli, meta: meta}
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
		fmt.Fprintln(c.cli.Writer, c.meta.Version)
		return 0
	}

	fmt.Fprintln(c.cli.ErrorWriter, usage())
	return 1
}
