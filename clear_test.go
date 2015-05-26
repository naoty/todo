package main

import "testing"

func TestClear(t *testing.T) {
	todos := []Todo{
		Todo{Title: "sample 1", Done: false},
		Todo{Title: "sample 2", Done: true},
		Todo{Title: "sample 3", Done: true},
	}

	actual, _ := clear(todos)
	expected := todos[0:1]

	if len(actual) != len(expected) {
		t.Errorf("delete(%q) = %q, want %q", todos, actual, expected)
	}
}
