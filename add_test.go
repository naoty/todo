package main

import "testing"

func TestAdd(t *testing.T) {
	todo := Todo{Title: "sample", Done: false}
	todos := []Todo{todo}
	add := addProcess("new sample", false)

	actual, _ := add(todos)
	expected := []Todo{todo, Todo{Title: "new sample", Done: false}}

	if len(actual) != len(expected) {
		t.Errorf("add(%q) = %q, want %q", todo, actual, expected)
	}
}

func TestAddOnce(t *testing.T) {
	todo := Todo{Title: "sample", Done: false}
	todos := []Todo{todo}
	add := addProcess("sample", true)

	actual, _ := add(todos)
	expected := []Todo{todo}

	if len(actual) != len(expected) {
		t.Errorf("add(%q) = %q, want %q", todo, actual, expected)
	}
}
