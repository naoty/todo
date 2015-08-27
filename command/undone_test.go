package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestUndone(t *testing.T) {
	todos := []todo.Todo{
		todo.Todo{Title: "dummy1", Done: true},
		todo.Todo{Title: "dummy2", Done: true},
		todo.Todo{Title: "dummy3", Done: true},
	}

	undone := newTodoUndoneProcess(2, 3)

	actual, _ := undone(todos)
	expected := []todo.Todo{
		todo.Todo{Title: "dummy1", Done: true},
		todo.Todo{Title: "dummy2", Done: false},
		todo.Todo{Title: "dummy3", Done: false},
	}

	if expected[0].Done != true {
		t.Errorf("done(%q) = %q, want %q", todos, actual, expected)
	}
	if expected[1].Done != false || expected[2].Done != false {
		t.Errorf("done(%q) = %q, want %q", todos, actual, expected)
	}
}
