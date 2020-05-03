package main

import (
	"os"

	"github.com/naoty/todo/cmd"
)

// Version is the version of this application.
// Makefile injects git tag into this value.
var Version = "0.0.0"

func main() {
	commandFactory := cmd.Lookup(os.Args)
	stdio := cmd.CLI{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
	config := cmd.Config{
		Version: Version,
	}
	command := commandFactory(stdio, config)
	status := command.Run(os.Args)
	os.Exit(status)
}
