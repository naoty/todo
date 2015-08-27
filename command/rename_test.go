package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestRename(t *testing.T) {
	todos := []todo.Todo{
		todo.Todo{Title: "dummy1", Done: false},
	}

	rename := newTodoRenameProcess(1, "dummy2")

	actual, _ := rename(todos)
	expected := []todo.Todo{
		todo.Todo{Title: "dummy2", Done: false},
	}

	if actual[0].Title != expected[0].Title {
		t.Errorf("rename(%q) = %q, want %q", todos, actual, expected)
	}
}
