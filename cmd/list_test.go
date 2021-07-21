package cmd_test

import (
	"bytes"
	"testing"

	"github.com/naoty/todo/cmd"
	"github.com/naoty/todo/repository/mock"
	"github.com/naoty/todo/todo"
)

func TestListRun(t *testing.T) {
	testcases := []struct {
		name   string
		input  []*todo.Todo
		output string
	}{
		{"empty", []*todo.Todo{}, ""},
		{"undone_todo", []*todo.Todo{{ID: 1, Title: "dummy", State: todo.Undone}}, "  1 | dummy\n"},
		{"done_todo", []*todo.Todo{{ID: 1, Title: "dummy", State: todo.Done}}, "  1 | \033[2;9mdummy\033[0m\n"},
		{"waiting_todo", []*todo.Todo{{ID: 1, Title: "dummy", State: todo.Waiting}}, "  1 | \033[2mdummy\033[0m\n"},
		{"right-aligned ID", []*todo.Todo{{ID: 1, Title: "dummy", State: todo.Undone}, {ID: 10, Title: "dummy", State: todo.Undone}}, "   1 | dummy\n  10 | dummy\n"},
		{"subtodo", []*todo.Todo{{ID: 1, Title: "dummy", State: todo.Undone, Todos: []*todo.Todo{{ID: 2, Title: "dummy", State: todo.Undone}}}}, "  1 | dummy\n    2 | dummy\n"},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			buf := bytes.NewBufferString("")
			command := cmd.NewList(cmd.CLI{Writer: buf}, "", mock.New(testcase.input))
			command.Run([]string{"todo", "list"})
			got := buf.String()
			if got != testcase.output {
				t.Errorf("got: %s, want: %s", got, testcase.output)
			}
		})
	}
}
