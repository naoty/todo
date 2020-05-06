package todo

import "fmt"

// Todo represents a TODO.
type Todo struct {
	ID    int
	Title string
	State State
}

// State represents the state of TODO.
type State int

const (
	// Undone represents state of undone TODO.
	Undone State = iota

	// Done represents state of done TODO.
	Done

	// Waiting represents state of TODO which I'm waiting for someone to finish.
	Waiting

	// Archived represents state of archived TODO.
	Archived
)

// New returns a new Todo.
func New(id int, title string) *Todo {
	return &Todo{ID: id, Title: title, State: Undone}
}

func (td *Todo) String() string {
	var mark string
	switch td.State {
	case Undone:
		mark = "[ ]"
	case Done:
		mark = "[x]"
	case Waiting:
		mark = "[w]"
	case Archived:
		mark = "[-]"
	}

	return fmt.Sprintf("%s %03d: %s", mark, td.ID, td.Title)
}
