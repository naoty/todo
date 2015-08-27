package command

import (
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
)

var Add = cli.Command{
	Name:  "add",
	Usage: "Add a TODO",
	Action: func(context *cli.Context) {
		status := ExecAdd(context)
		os.Exit(status)
	},
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "once, o",
			Usage: "Add the TODO only if it exists",
		},
	},
}

func ExecAdd(context *cli.Context) int {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "add")
		return 1
	}

	title := strings.Join(context.Args(), " ")
	add := newTodoAddProcess(title, context.Bool("once"))

	file := todo.OpenFile()
	err := file.Update(add)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoAddProcess(title string, isOnce bool) todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		if isOnce && hasTodo(todos, title) {
			return todos, nil
		}

		todo := todo.Todo{Title: title, Done: false}
		return append(todos, todo), nil
	}
}

func hasTodo(todos []todo.Todo, title string) bool {
	for _, todo := range todos {
		if todo.Title == title {
			return true
		}
	}
	return false
}
