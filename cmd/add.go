package cmd

import (
	"fmt"
	"strings"

	"github.com/naoty/todo/repository"
	"github.com/spf13/pflag"
)

// Add represents `add` subcommand.
type Add struct {
	cli  CLI
	repo repository.Repository
}

// NewAdd returns a new Add.
func NewAdd(cli CLI, version string, repo repository.Repository) Command {
	return &Add{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Add) Run(args []string) int {
	flagset := pflag.NewFlagSet("add", pflag.ExitOnError)
	parent := flagset.IntP("parent", "p", 0, "")

	flagset.Parse(args)

	if flagset.NArg() < 3 {
		fmt.Fprintf(c.cli.ErrorWriter, usage())
		return 1
	}

	title := strings.Join(flagset.Args()[2:], " ")
	err := c.repo.Add(title, parent)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	return 0
}
