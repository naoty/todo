package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

var Delete = cli.Command{
	Name:   "delete",
	Usage:  "Delete a TODO",
	Action: delete,
}

func delete(context *cli.Context) {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "add")
		os.Exit(1)
	}

	num, err := strconv.Atoi(context.Args().First())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = DeleteTodo(num)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
