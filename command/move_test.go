package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestMove(t *testing.T) {
	todos := []todo.Todo{
		todo.Todo{Number: 1, Title: "dummy1", Done: true},
		todo.Todo{Number: 2, Title: "dummy2", Done: true},
	}

	move := newTodoMoveProcess(2, 1)

	actual, _ := move(todos)
	expected := []todo.Todo{
		todo.Todo{Number: 2, Title: "dummy2", Done: true},
		todo.Todo{Number: 1, Title: "dummy1", Done: true},
	}

	if expected[0].Number != 2 || expected[1].Number != 1 {
		t.Errorf("move(%q) = %q, want %q", todos, actual, expected)
	}
}
