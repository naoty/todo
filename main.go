package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/naoty/todo/cmd"
	"github.com/naoty/todo/repository/filesystem"
)

// Version is the version of this application.
// Makefile injects git tag into this value.
var Version = "0.0.0"

func main() {
	stdio := cmd.CLI{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	root := rootPath()
	repo, err := filesystem.New(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	commandFactory := cmd.Lookup(os.Args)
	command := commandFactory(stdio, Version, repo)
	status := command.Run(os.Args)
	os.Exit(status)
}

func rootPath() string {
	root := os.Getenv("TODOS_PATH")
	if root == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		root = filepath.Join(home, ".todos")
	}

	return root
}
