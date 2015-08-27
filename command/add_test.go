package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestAdd(t *testing.T) {
	todo1 := todo.Todo{Title: "dummy1", Done: false}
	todo2 := todo.Todo{Title: "dummy2", Done: false}
	todos := []todo.Todo{todo1}

	add := newTodoAddProcess("dummy2", false)

	actual, _ := add(todos)
	expected := []todo.Todo{todo1, todo2}

	if len(actual) != len(expected) {
		t.Errorf("add(%q) = %q, want %q", todo1, actual, expected)
	}
}

func TestAddOnce(t *testing.T) {
	todo1 := todo.Todo{Title: "dummy", Done: false}
	todos := []todo.Todo{todo1}

	add := newTodoAddProcess("dummy", true)

	actual, _ := add(todos)
	expected := []todo.Todo{todo1}

	if len(actual) != len(expected) {
		t.Errorf("add(%q) = %q, want %q", todo1, actual, expected)
	}
}
