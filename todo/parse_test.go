package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/naoty/todo/todo"
)

func TestParse(t *testing.T) {
	testcases := []struct {
		path string
		td   *todo.Todo
	}{
		{"testdata/1.md", todo.New(0, "dummy")},
	}

	for _, testcase := range testcases {
		t.Run(testcase.path, func(t *testing.T) {
			file, err := os.Open(testcase.path)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			content, err := ioutil.ReadAll(file)
			if err != nil {
				t.Fatal(err)
			}

			td, err := todo.Parse(string(content))
			if err != nil {
				t.Fatal(err)
			}

			if *td != *testcase.td {
				t.Errorf("got: %+v, want: %+v", *td, *testcase.td)
			}
		})
	}
}
