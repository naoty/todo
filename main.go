package main

import (
	"fmt"
	"os"

	"github.com/naoty/todo/cmd"
	"github.com/naoty/todo/repository/filesystem"
)

// Version is the version of this application.
// Makefile injects git tag into this value.
var Version = "0.0.0"

func main() {
	stdio := cmd.CLI{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	repo, err := filesystem.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	commandFactory := cmd.Lookup(os.Args)
	command := commandFactory(stdio, Version, repo)
	status := command.Run(os.Args)
	os.Exit(status)
}
