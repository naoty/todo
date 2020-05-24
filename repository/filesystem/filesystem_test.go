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

func TestDelete(t *testing.T) {
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

	err = repo.Delete(1)
	if err != nil {
		t.Fatal("failed to delete a TODO")
	}

	if _, err := os.Stat("./testdata/sandbox/1.md"); os.IsExist(err) {
		t.Error("1.md is not deleted")
	}
}

func TestMove(t *testing.T) {
	testcases := []struct {
		id       int
		parent   int
		position int
		want     []int
	}{
		{3, 0, 1, []int{3, 1, 2}},
	}

	for _, testcase := range testcases {

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
		for i := 0; i < 3; i++ {
			err = repo.Add("dummy", &parent)
			if err != nil {
				t.Fatal("failed to add a TODO")
			}
		}

		err = repo.Move(testcase.id, &testcase.parent, testcase.position)
		if err != nil {
			t.Fatal("failed to move a TODO")
		}

		todos, err := repo.List()

		ids := make([]int, len(todos))
		for i, td := range todos {
			ids[i] = td.ID
		}

		if !reflect.DeepEqual(ids, testcase.want) {
			t.Errorf("got: %v, want: %v", ids, testcase.want)
		}
	}
}

func TestArchive(t *testing.T) {
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

	err = repo.Archive(1)
	if err != nil {
		t.Fatal("failed to archive a TODO")
	}

	if _, err := os.Stat("./testdata/sandbox/archived/1.md"); os.IsNotExist(err) {
		t.Fatal("archived TODO is deleted")
	}

	todos, err := repo.List()
	if len(todos) != 0 {
		t.Errorf("got: %d, want: 0", len(todos))
	}
}
