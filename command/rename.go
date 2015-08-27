package command

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
)

var Rename = cli.Command{
	Name:  "rename",
	Usage: "Rename a TODO",
	Action: func(context *cli.Context) {
		status := ExecRename(context)
		os.Exit(status)
	},
}

func ExecRename(context *cli.Context) int {
	if len(context.Args()) < 2 {
		cli.ShowCommandHelp(context, "rename")
		return 1
	}

	num, err := strconv.Atoi(context.Args()[0])
	if err != nil {
		log.Println(err)
		return 1
	}

	title := strings.Join(context.Args()[1:], " ")
	rename := newTodoRenameProcess(num, title)

	file := todo.OpenFile()
	err = file.Update(rename)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoRenameProcess(num int, title string) todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		i := num - 1
		if i >= len(todos) {
			return nil, errors.New("Index out of bounds.")
		}

		newTodo := todos[i]
		newTodo.Title = title

		newTodos := make([]todo.Todo, len(todos))
		copy(newTodos, todos)
		newTodos[i] = newTodo

		return newTodos, nil
	}
}
