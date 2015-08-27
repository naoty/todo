package formatter

import (
	"fmt"
	"io"

	"github.com/naoty/todo/todo"
)

type NumberedFormatter struct {
	writer io.Writer
	Mode   Mode
}

func NewNumberedFormatter(w io.Writer, m Mode) *NumberedFormatter {
	return &NumberedFormatter{writer: w, Mode: m}
}

func (f *NumberedFormatter) Print(todos []todo.Todo) error {
	for i, todo := range todos {
		if f.Mode == DONE && !todo.Done {
			continue
		}
		if f.Mode == UNDONE && todo.Done {
			continue
		}
		mark := NewMark(todo)
		fmt.Fprintf(f.writer, "%s %03d: %s\n", mark, i+1, todo.Title)
	}

	return nil
}
