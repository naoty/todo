package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

var Move = cli.Command{
	Name:   "move",
	Usage:  "move a TODO",
	Action: move,
}

func move(context *cli.Context) {
	if len(context.Args()) < 2 {
		cli.ShowCommandHelp(context, "move")
		os.Exit(1)
	}

	from, err := strconv.Atoi(context.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	to, err := strconv.Atoi(context.Args()[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = MoveTodo(from, to)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
