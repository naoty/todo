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
		cli.ShowCommandHelp(context, "undone")
		os.Exit(1)
	}

	var err error
	nums := []int{}

	for _, arg := range context.Args() {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		nums = append(nums, num)
	}

	err = UndoneTodo(nums...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
