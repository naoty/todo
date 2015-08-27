package formatter

import (
	"github.com/naoty/todo/todo"
)

type Formatter interface {
	Print(todos []todo.Todo) error
}
