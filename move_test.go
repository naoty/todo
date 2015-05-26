package main

import "testing"

func TestMove(t *testing.T) {
	todos := []Todo{
		Todo{Number: 1, Title: "sample 1", Done: false},
		Todo{Number: 2, Title: "sample 2", Done: false},
	}
	move := moveProcess(2, 1)

	actual, _ := move(todos)
	expected := []Todo{
		Todo{Number: 2, Title: "sample 2", Done: false},
		Todo{Number: 1, Title: "sample 1", Done: false},
	}

	for i, todo := range actual {
		another := expected[i]
		if todo.Number != another.Number {
			t.Errorf("move(%q) = %q, want %q", todos, actual, expected)
			break
		}
	}
}
