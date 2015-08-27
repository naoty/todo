package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestDone(t *testing.T) {
	todos := []todo.Todo{
		todo.Todo{Title: "dummy1", Done: false},
		todo.Todo{Title: "dummy2", Done: false},
		todo.Todo{Title: "dummy3", Done: false},
	}

	done := newTodoDoneProcess(2, 3)

	actual, _ := done(todos)
	expected := []todo.Todo{
		todo.Todo{Title: "dummy1", Done: false},
		todo.Todo{Title: "dummy2", Done: true},
		todo.Todo{Title: "dummy3", Done: true},
	}

	if expected[0].Done != false {
		t.Errorf("done(%q) = %q, want %q", todos, actual, expected)
	}
	if expected[1].Done != true || expected[2].Done != true {
		t.Errorf("done(%q) = %q, want %q", todos, actual, expected)
	}
}
