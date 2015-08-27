package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestDelete(t *testing.T) {
	todos := []todo.Todo{
		todo.Todo{Title: "dummy1", Done: false},
		todo.Todo{Title: "dummy2", Done: false},
		todo.Todo{Title: "dummy3", Done: false},
	}

	delete := newTodoDeleteProcess(2, 3)

	actual, _ := delete(todos)
	expected := []todo.Todo{
		todo.Todo{Title: "dummy1", Done: false},
	}

	if len(actual) != len(expected) {
		t.Errorf("delete(%q) = %q, want %q", todos, actual, expected)
	}
}
