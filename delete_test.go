package main

import "testing"

func TestDelete(t *testing.T) {
	todos := []Todo{
		Todo{Title: "sample 1", Done: false},
		Todo{Title: "sample 2", Done: false},
		Todo{Title: "sample 3", Done: false},
	}
	delete := deleteProcess(2, 3)

	actual, _ := delete(todos)
	expected := []Todo{
		Todo{Title: "sample 1", Done: false},
	}

	if len(actual) != len(expected) {
		t.Errorf("delete(%q) = %q, want %q", todos, actual, expected)
	}
}
