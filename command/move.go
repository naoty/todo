package command

import (
	"errors"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
	"github.com/naoty/todo/todoutil"
)

var Move = cli.Command{
	Name:  "move",
	Usage: "move a TODO",
	Action: func(context *cli.Context) {
		status := ExecMove(context)
		os.Exit(status)
	},
}

func ExecMove(context *cli.Context) int {
	if len(context.Args()) < 2 {
		cli.ShowCommandHelp(context, "move")
		return 1
	}

	nums, err := todoutil.Atois(context.Args())
	if err != nil {
		log.Println(err)
		return 1
	}

	move := newTodoMoveProcess(nums[0], nums[1])

	file := todo.OpenFile()
	err = file.Update(move)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoMoveProcess(from, to int) todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		fromIndex, toIndex := from-1, to-1
		if fromIndex >= len(todos) || toIndex >= len(todos) {
			return nil, errors.New("Index out of bounds.")
		}

		movedTodo := todos[fromIndex]
		todos = append(todos[:fromIndex], todos[fromIndex+1:]...)
		todos = append(todos[:toIndex], append([]todo.Todo{movedTodo}, todos[toIndex:]...)...)
		return todos, nil
	}
}
