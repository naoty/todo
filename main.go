package main

import (
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "todo"
	app.Author = "Naoto Kaneko"
	app.Email = "naoty.k@gmail.com"
	app.Version = "0.1.0"
	app.Usage = "Manage TODOs"
	app.Commands = []cli.Command{
		List,
		Add,
		Delete,
		Move,
		Rename,
		Done,
		Undone,
		Clear,
	}
	app.Run(os.Args)
}

// Utility functions

func Atois(words []string) ([]int, error) {
	var nums []int
	for _, word := range words {
		num, err := strconv.Atoi(word)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func Contains(xs []int, n int) bool {
	for _, x := range xs {
		if x == n {
			return true
		}
	}
	return false
}
