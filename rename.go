package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
)

var Rename = cli.Command{
	Name:   "rename",
	Usage:  "Rename a TODO",
	Action: rename,
}

func rename(context *cli.Context) {
	if len(context.Args()) < 2 {
		cli.ShowCommandHelp(context, "rename")
		os.Exit(1)
	}

	num, err := strconv.Atoi(context.Args().First())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	newTitle := strings.Join(context.Args()[1:], " ")
	err = RenameTodo(num, newTitle)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
