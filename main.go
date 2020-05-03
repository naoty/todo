package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/naoty/todo/cmd"
	"github.com/naoty/todo/repository"
)

// Version is the version of this application.
// Makefile injects git tag into this value.
var Version = "0.0.0"

func main() {
	commandFactory := cmd.Lookup(os.Args)
	stdio := cmd.CLI{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
	config := cmd.Config{
		TodosPath: ensureTodosPath(),
		Version:   Version,
	}
	repo := repository.NewFS(ensureTodosPath())
	command := commandFactory(stdio, config, repo)
	status := command.Run(os.Args)
	os.Exit(status)
}

func ensureTodosPath() string {
	home := os.Getenv("TODOS_PATH")

	if home == "" {
		var err error
		home, err = os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	path := filepath.Join(home, ".todos")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	return path
}
