package filesystem_test

import (
	"errors"
	"fmt"
	"os"
	"reflect"
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

	repo, err := filesystem.New("./testdata/todos")
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

func TestList(t *testing.T) {
	repo, err := filesystem.New("./testdata/todos")
	if err != nil {
		t.Fatalf("failed to initialize repository: %v", err)
	}

	todos, err := repo.List()
	if err != nil {
		t.Fatalf("failed to list todos")
	}

	ids := make([]int, 3)
	for i, td := range todos {
		ids[i] = td.ID
	}

	want := []int{2, 1, 3}
	if !reflect.DeepEqual(ids, want) {
		t.Errorf("got: %v, want: %v", ids, want)
	}
}

func TestAdd(t *testing.T) {
	repo, err := filesystem.New("./testdata/sandbox")
	if err != nil {
		t.Fatalf("failed to initialize repository: %v", err)
	}

	t.Cleanup(func() {
		err := os.RemoveAll("./testdata/sandbox")
		if err != nil {
			t.Fatalf("failed to cleanup sandbox: %v", err)
		}
	})

	parent := 0
	err = repo.Add("dummy", &parent)
	if err != nil {
		t.Fatalf("failed to add a TODO")
	}

	td, err := repo.Get(1)
	if err != nil {
		t.Fatalf("failed to get a new TODO")
	}

	if td.Title != "dummy" {
		t.Errorf("got: %s, want: dummy", td.Title)
	}

	if td.State != todo.Undone {
		t.Errorf("got: %s, want: %s", td.State, todo.Undone)
	}

	if td.Body != "" {
		t.Errorf("got: %s, want: ''", td.Body)
	}
}

func TestUpdate(t *testing.T) {
	repo, err := filesystem.New("./testdata/sandbox")
	if err != nil {
		t.Fatalf("failed to initialize repository: %v", err)
	}

	t.Cleanup(func() {
		err := os.RemoveAll("./testdata/sandbox")
		if err != nil {
			t.Fatalf("failed to cleanup sandbox: %v", err)
		}
	})

	parent := 0
	err = repo.Add("dummy", &parent)
	if err != nil {
		t.Fatal("failed to add a TODO")
	}

	td, err := repo.Get(1)
	if err != nil {
		t.Fatal("failed to get a new TODO")
	}

	td.State = todo.Done
	err = repo.Update(td)
	if err != nil {
		t.Fatal("failed to update a TODO")
	}

	td, err = repo.Get(1)
	if err != nil {
		t.Fatal("failed to get a new TODO")
	}

	if td.State != todo.Done {
		t.Errorf("got: %s, want: %s", td.State, todo.Done)
	}
}
