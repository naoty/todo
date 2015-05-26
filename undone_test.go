package main

import "testing"

func TestUndone(t *testing.T) {
	todos := []Todo{
		Todo{Title: "sample 1", Done: true},
		Todo{Title: "sample 2", Done: true},
		Todo{Title: "sample 2", Done: true},
	}
	undone := undoneProcess(2, 3)

	actual, _ := undone(todos)
	expected := []Todo{
		Todo{Title: "sample 1", Done: true},
		Todo{Title: "sample 2", Done: false},
		Todo{Title: "sample 2", Done: false},
	}

	for i, todo := range actual {
		switch i {
		case 0:
			if !todo.Done {
				t.Errorf("undone(%q) = %q, want %q", todos, actual, expected)
				break
			}
		default:
			if todo.Done {
				t.Errorf("undone(%q) = %q, want %q", todos, actual, expected)
				break
			}
		}
	}
}
