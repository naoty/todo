package filesystem_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/naoty/todo/repository/filesystem"
	"github.com/naoty/todo/todo"
)

func TestGet(t *testing.T) {
	testcases := []struct {
		input  int
		output *todo.Todo
		err    error
	}{
		{1, &todo.Todo{ID: 1, Title: "dummy", State: todo.Undone}, nil},
		{2, &todo.Todo{ID: 2, Title: "dummy", State: todo.Done}, nil},
		{3, &todo.Todo{ID: 3, Title: "dummy", State: todo.Waiting}, nil},
		{1000, nil, filesystem.ErrTODONotFound},
	}

	repo, err := filesystem.New("./testdata")
	if err != nil {
		t.Fatalf("failed to initialize repository: %v", err)
	}

	for _, testcase := range testcases {
		name := fmt.Sprintf("ID:%d", testcase.input)
		t.Run(name, func(t *testing.T) {
			td, err := repo.Get(testcase.input)

			if err != nil {
				if !errors.Is(err, testcase.err) {
					t.Errorf("got: %v, want: %v", err, testcase.err)
				}
				return
			}

			if !td.Equal(testcase.output) {
				t.Errorf("got: %+v, want: %+v", td, testcase.output)
			}
		})
	}
}
