package formatter

import (
	"github.com/naoty/todo/todo"
)

const (
	DoneMark   = "[x]"
	UndoneMark = "[ ]"
)

func NewMark(todo todo.Todo) string {
	if todo.Done {
		return DoneMark
	} else {
		return UndoneMark
	}
}
