package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

var Done = cli.Command{
	Name:   "done",
	Usage:  "Done a TODO",
	Action: done,
}

func done(context *cli.Context) {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "add")
		os.Exit(1)
	}

	num, err := strconv.Atoi(context.Args().First())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = DoneTodo(num)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
