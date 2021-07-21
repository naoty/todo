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

	maxWidth := 0
	for _, td := range todos {
		idString := fmt.Sprintf("%d", td.ID)
		if len(idString) > maxWidth {
			maxWidth = len(idString)
		}
	}

	for _, td := range todos {
		c.printTodos(td, 0, maxWidth)
	}

	return 0
}

func (c *List) printTodos(td *todo.Todo, level, width int) {
	var decoratedTitle string
	switch td.State {
	case todo.Undone:
		decoratedTitle = td.Title
	case todo.Done:
		decoratedTitle = fmt.Sprintf("\033[2;9m%s\033[0m", td.Title)
	case todo.Waiting:
		decoratedTitle = fmt.Sprintf("\033[2m%s\033[0m", td.Title)
	}

	indent := strings.Repeat(" ", (level+1)*2)
	fmt.Fprintf(c.cli.Writer, "%s%*d | %s\n", indent, width, td.ID, decoratedTitle)

	maxWidth := 0
	for _, td := range td.Todos {
		idString := fmt.Sprintf("%d", td.ID)
		if len(idString) > maxWidth {
			maxWidth = len(idString)
		}
	}

	for _, sub := range td.Todos {
		c.printTodos(sub, level+1, maxWidth)
	}
}
