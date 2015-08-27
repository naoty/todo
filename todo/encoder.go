package todo

import (
	"io"

	"github.com/ymotongpoo/goltsv"
)

type Encoder struct {
	writer io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{writer: w}
}

func (e *Encoder) Encode(todos []Todo) error {
	data := make([]map[string]string, len(todos))
	for i, todo := range todos {
		data[i] = e.encodeTodo(todo)
	}

	writer := goltsv.NewWriter(e.writer)
	err := writer.WriteAll(data)

	return err
}

func (e *Encoder) encodeTodo(todo Todo) map[string]string {
	m := make(map[string]string)

	m["title"] = todo.Title

	if todo.Done {
		m["done"] = "true"
	} else {
		m["done"] = "false"
	}

	return m
}
