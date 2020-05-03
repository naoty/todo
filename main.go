package main

import (
	"os"

	"github.com/naoty/todo/cmd"
)

func main() {
	commandFactory := cmd.Lookup(os.Args)
	command := commandFactory(cmd.CLI{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	})
	status := command.Run(os.Args)
	os.Exit(status)
}
