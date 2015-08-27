package command

import (
	"errors"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
	"github.com/naoty/todo/todoutil"
)

var Done = cli.Command{
	Name:  "done",
	Usage: "Done TODOs",
	Action: func(context *cli.Context) {
		status := ExecDone(context)
		os.Exit(status)
	},
}

func ExecDone(context *cli.Context) int {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "done")
		return 1
	}

	nums, err := todoutil.Atois(context.Args())
	if err != nil {
		log.Println(err)
		return 1
	}

	done := newTodoDoneProcess(nums...)

	file := todo.OpenFile()
	err = file.Update(done)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoDoneProcess(nums ...int) todo.TodoProcess {
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
				newTodo.Done = true
			}
			newTodos[i] = newTodo
		}

		return newTodos, nil
	}
}
