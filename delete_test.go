package main

import "testing"

func TestDelete(t *testing.T) {
	todos := []Todo{
		Todo{Title: "sample 1", Done: false},
		Todo{Title: "sample 2", Done: false},
	}
	delete := deleteProcess(2)

	actual, _ := delete(todos)
	expected := todos[0:1]

	if len(actual) != len(expected) {
		t.Errorf("delete(%q) = %q, want %q", todos, actual, expected)
	}
}
