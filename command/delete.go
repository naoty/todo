package command

import (
	"errors"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
	"github.com/naoty/todo/todoutil"
)

var Delete = cli.Command{
	Name:  "delete",
	Usage: "Delete a TODO",
	Action: func(context *cli.Context) {
		status := ExecDelete(context)
		os.Exit(status)
	},
}

func ExecDelete(context *cli.Context) int {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "delete")
		return 1
	}

	nums, err := todoutil.Atois(context.Args())
	if err != nil {
		log.Println(err)
		return 1
	}

	delete := newTodoDeleteProcess(nums...)

	file := todo.OpenFile()
	err = file.Update(delete)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoDeleteProcess(nums ...int) todo.TodoProcess {
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

		newTodos := make([]todo.Todo, 0)
		for i, todo := range todos {
			if !todoutil.ContainsInt(indices, i) {
				newTodos = append(newTodos, todo)
			}
		}

		return newTodos, nil
	}
}
