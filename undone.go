package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

var Undone = cli.Command{
	Name:   "undone",
	Usage:  "Undone a TODO",
	Action: undone,
}

func undone(context *cli.Context) {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "add")
		os.Exit(1)
	}

	num, err := strconv.Atoi(context.Args().First())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = UndoneTodo(num)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
