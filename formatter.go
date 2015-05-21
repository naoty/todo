package main

import (
	"fmt"
	"io"
)

const (
	DoneMark   = "[x]"
	UndoneMark = "[ ]"
)

type TodoFormatter interface {
	Println(todo Todo)
}

type StandardFormatter struct {
	Out, Err io.Writer
}

func (f StandardFormatter) Println(todo Todo) {
	var mark string
	if todo.Done {
		mark = DoneMark
	} else {
		mark = UndoneMark
	}
	fmt.Fprintf(f.Out, "%s %03d: %s\n", mark, todo.Number, todo.Title)
}

type MarkdownFormatter struct {
	Out, Err io.Writer
}

func (f MarkdownFormatter) Println(todo Todo) {
	var mark string
	if todo.Done {
		mark = DoneMark
	} else {
		mark = UndoneMark
	}
	fmt.Fprintf(f.Out, "- %s %s\n", mark, todo.Title)
}
