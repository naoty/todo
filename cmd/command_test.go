package cmd_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/naoty/todo/cmd"
	"github.com/naoty/todo/repository/mock"
	"github.com/naoty/todo/todo"
)

func TestLookup(t *testing.T) {
	testcases := []struct {
		args    []string
		cmdType string
	}{
		{[]string{"todo"}, "Root"},
		{[]string{"todo", "list"}, "List"},
		{[]string{"todo", "add"}, "Add"},
		{[]string{"todo", "delete"}, "Delete"},
		{[]string{"todo", "move"}, "Move"},
		{[]string{"todo", "open"}, "Open"},
		{[]string{"todo", "done"}, "Done"},
		{[]string{"todo", "undone"}, "Undone"},
		{[]string{"todo", "wait"}, "Wait"},
		{[]string{"todo", "archive"}, "Archive"},
	}

	cli := cmd.CLI{}
	repo := mock.New([]*todo.Todo{})

	for _, testcase := range testcases {
		name := strings.Join(testcase.args, " ")
		t.Run(name, func(t *testing.T) {
			factory := cmd.Lookup(testcase.args)
			command := factory(cli, "0.0.0", repo)
			cmdType := reflect.TypeOf(command).Elem().Name()
			if cmdType != testcase.cmdType {
				t.Errorf("got: %s, want: %s", cmdType, testcase.cmdType)
			}
		})
	}
}
