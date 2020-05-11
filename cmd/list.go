package cmd

import (
	"fmt"
	"strings"

	"github.com/naoty/todo/repository"
	"github.com/naoty/todo/todo"
)

// List represents `list` subcommand.
type List struct {
	cli  CLI
	repo repository.Repository
}

// NewList returns a new List.
func NewList(cli CLI, version string, repo repository.Repository) Command {
	return &List{cli: cli, repo: repo}
}

// Run implements Command interface.
func (c *List) Run(args []string) int {
	todos, err := c.repo.List()
	if err != nil {
		fmt.Fprintln(c.cli.ErrorWriter, err)
		return 1
	}

	for _, td := range todos {
		c.printTodos(td, 0)
	}

	return 0
}

func (c *List) printTodos(td *todo.Todo, level int) {
	var mark string
	switch td.State {
	case todo.Undone:
		mark = "[ ]"
	case todo.Done:
		mark = "[x]"
	case todo.Waiting:
		mark = "[w]"
	case todo.Archived:
		return
	}

	indent := strings.Repeat(" ", level*2)
	fmt.Fprintf(c.cli.Writer, "%s%s %03d: %s\n", indent, mark, td.ID, td.Title)

	for _, sub := range td.Todos {
		c.printTodos(sub, level+1)
	}
}
