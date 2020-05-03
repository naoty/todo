package todo

import "fmt"

// Todo represents a TODO.
type Todo struct {
	id    int
	title string
}

// New returns a new Todo.
func New(id int, title string) *Todo {
	return &Todo{id: id, title: title}
}

func (td *Todo) String() string {
	return fmt.Sprintf("[ ] %03d: %s", td.id, td.title)
}

// SetID is setter for id field.
func (td *Todo) SetID(id int) {
	td.id = id
}
