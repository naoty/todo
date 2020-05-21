package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/naoty/todo/cmd"
	"github.com/naoty/todo/repository/mock"
	"github.com/naoty/todo/todo"
)

func TestRootRun(t *testing.T) {
	testcases := []struct {
		args    []string
		message string
	}{
		{[]string{"todo", "-v"}, "0.0.0\n"},
	}

	for _, testcase := range testcases {
		name := strings.Join(testcase.args, " ")
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBufferString("")
			command := cmd.NewRoot(cmd.CLI{Writer: buf}, "0.0.0", mock.New([]*todo.Todo{}))
			command.Run(testcase.args)
			got := buf.String()
			if got != testcase.message {
				t.Errorf("got: %s, want: %s", got, testcase.message)
			}
		})
	}
}
