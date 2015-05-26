package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

var Done = cli.Command{
	Name:   "done",
	Usage:  "Done TODOs",
	Action: done,
}

func done(context *cli.Context) {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "done")
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

	err = DoneTodo(nums...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
