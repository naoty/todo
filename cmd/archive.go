package cmd

import (
	"fmt"
	"strconv"

	"github.com/naoty/todo/repository"
	"github.com/naoty/todo/todo"
)

// Archive represents `archive` subcommand.
type Archive struct {
	cli  CLI
	repo repository.Repository
}

// NewArchive returns a new Archive.
func NewArchive(cli CLI, version string, repo repository.Repository) Command {
	return &Archive{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *Archive) Run(args []string) int {
	if len(args) == 2 {
		return c.archiveTodos(args)
	}

	return c.archiveTodo(args)
}

func (c *Archive) archiveTodo(args []string) int {
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	td, err := c.repo.Get(id)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	err = c.archive(td)
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, usage())
		return 1
	}

	return 0
}

func (c *Archive) archiveTodos(args []string) int {
	todos, err := c.repo.List()
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	for _, td := range todos {
		err := c.archive(td)
		if err != nil {
			fmt.Fprintln(c.cli.ErrorWriter, err)
			return 1
		}
	}

	return 0
}

func (c *Archive) archive(td *todo.Todo) error {
	if td.State == todo.Done {
		td.State = todo.Archived
		err := c.repo.Update(td)
		if err != nil {
			return err
		}
	}

	for _, sub := range td.Todos {
		err := c.archive(sub)
		if err != nil {
			return err
		}
	}

	return nil
}
