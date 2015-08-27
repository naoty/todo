package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestClear(t *testing.T) {
	todos := []todo.Todo{
		todo.Todo{Title: "dummy1", Done: false},
		todo.Todo{Title: "dummy2", Done: true},
		todo.Todo{Title: "dummy3", Done: true},
	}

	clear := newTodoClearProcess()

	actual, _ := clear(todos)
	expected := []todo.Todo{
		todo.Todo{Title: "dummy1", Done: false},
	}

	if len(actual) != len(expected) {
		t.Errorf("delete(%q) = %q, want %q", todos, actual, expected)
	}
}
