package main

import "testing"

func TestDone(t *testing.T) {
	todos := []Todo{
		Todo{Title: "sample 1", Done: false},
		Todo{Title: "sample 2", Done: false},
		Todo{Title: "sample 2", Done: false},
	}
	done := doneProcess(2, 3)

	actual, _ := done(todos)
	expected := []Todo{
		Todo{Title: "sample 1", Done: false},
		Todo{Title: "sample 2", Done: true},
		Todo{Title: "sample 2", Done: true},
	}

	for i, todo := range actual {
		switch i {
		case 0:
			if todo.Done {
				t.Errorf("done(%q) = %q, want %q", todos, actual, expected)
				break
			}
		default:
			if !todo.Done {
				t.Errorf("done(%q) = %q, want %q", todos, actual, expected)
				break
			}
		}
	}
}
