package main

import "testing"

func TestRename(t *testing.T) {
	todos := []Todo{
		Todo{Title: "sample", Done: false},
	}
	rename := renameProcess(1, "new sample")

	actual, _ := rename(todos)
	expected := []Todo{
		Todo{Title: "new sample", Done: false},
	}

	if actual[0].Title != expected[0].Title {
		t.Errorf("rename(%q) = %q, want %q", todos, actual, expected)
	}
}
