package command

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
)

var Clear = cli.Command{
	Name:  "clear",
	Usage: "Clear done TODOs",
	Action: func(context *cli.Context) {
		status := ExecClear(context)
		os.Exit(status)
	},
}

func ExecClear(context *cli.Context) int {
	clear := newTodoClearProcess()
	file := todo.OpenFile()
	file.Update(clear)

	return 0
}

func newTodoClearProcess() todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		newTodos := make([]todo.Todo, 0)
		for _, todo := range todos {
			if !todo.Done {
				newTodos = append(newTodos, todo)
			}
		}
		return newTodos, nil
	}
}
