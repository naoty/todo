package formatter

import (
	"fmt"
	"io"

	"github.com/naoty/todo/todo"
)

type MarkdownFormatter struct {
	writer io.Writer
	Mode   Mode
}

func NewMarkdownFormatter(w io.Writer, m Mode) *MarkdownFormatter {
	return &MarkdownFormatter{writer: w, Mode: m}
}

func (f *MarkdownFormatter) Print(todos []todo.Todo) error {
	for _, todo := range todos {
		if f.Mode == DONE && !todo.Done {
			continue
		}
		if f.Mode == UNDONE && todo.Done {
			continue
		}
		mark := NewMark(todo)
		fmt.Fprintf(f.writer, "- %s %s\n", mark, todo.Title)
	}

	return nil
}
