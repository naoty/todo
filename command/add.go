package command

import (
	"strconv"
	"strings"

	"github.com/naoty/todo/todo"
	"github.com/urfave/cli"
)

// Add is a command to add a todo.
var Add = cli.Command{
	Name:   "add",
	Usage:  "Add a todo",
	Action: add,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "parent, p",
			Usage: "Specify the parent todo",
		},
	},
}

func add(c *cli.Context) error {
	if c.NArg() < 1 {
		cli.ShowCommandHelp(c, "add")
		return nil
	}

	path := todoFilePath()
	todos, err := readTodos(path)
	if err != nil {
		return err
	}

	title := strings.Join(c.Args(), " ")
	todo := todo.Todo{Title: title, Done: false, Todos: []todo.Todo{}}

	var orders []int
	parent := c.String("parent")
	for _, id := range strings.Split(parent, "-") {
		if order, err2 := strconv.Atoi(id); err2 == nil {
			orders = append(orders, order)
		}
	}

	todos = addTodo(todo, todos, orders)

	err = writeTodos(todos, path)
	if err != nil {
		return err
	}

	return nil
}

func addTodo(todo todo.Todo, todos []todo.Todo, orders []int) []todo.Todo {
	if len(orders) == 0 {
		todos = append(todos, todo)
		return todos
	}

	i := orders[0] - 1
	if i >= len(todos) {
		return todos
	}

	t := todos[i]
	if len(orders) == 1 {
		t.Todos = append(todo.Todos, todo)
	} else {
		t.Todos = addTodo(todo, t.Todos, orders[1:])
	}
	todos[i] = t

	return todos
}
