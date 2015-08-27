package command

import (
	"errors"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
	"github.com/naoty/todo/todoutil"
)

var Undone = cli.Command{
	Name:  "undone",
	Usage: "Undone TODOs",
	Action: func(context *cli.Context) {
		status := ExecUndone(context)
		os.Exit(status)
	},
}

func ExecUndone(context *cli.Context) int {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "undone")
		return 1
	}

	nums, err := todoutil.Atois(context.Args())
	if err != nil {
		log.Println(err)
		return 1
	}

	undone := newTodoUndoneProcess(nums...)

	file := todo.OpenFile()
	err = file.Update(undone)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoUndoneProcess(nums ...int) todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		indices := make([]int, 0)

		var err error
		for _, num := range nums {
			index := num - 1
			if index >= len(todos) {
				err = errors.New("Index out of bounds.")
			}
			indices = append(indices, index)
		}
		if err != nil {
			return nil, err
		}

		newTodos := make([]todo.Todo, len(todos))
		for i, todo := range todos {
			newTodo := todo
			if todoutil.ContainsInt(indices, i) {
				newTodo.Done = false
			}
			newTodos[i] = newTodo
		}

		return newTodos, nil
	}
}
