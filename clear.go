package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

var Clear = cli.Command{
	Name:   "clear",
	Usage:  "Clear done TODOs",
	Action: clear,
}

func clear(context *cli.Context) {
	err := ClearTodos()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
