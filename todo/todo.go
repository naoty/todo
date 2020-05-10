package todo

import "fmt"

// Todo represents a TODO.
type Todo struct {
	ID    int
	Title string
	State State
	Body  string
	Todos []*Todo
}

// State represents the state of TODO.
type State string

const (
	// Undone represents state of undone TODO.
	Undone State = "UNDONE"

	// Done represents state of done TODO.
	Done State = "DONE"

	// Waiting represents state of TODO which I'm waiting for someone to finish.
	Waiting State = "WAITING"

	// Archived represents state of archived TODO.
	Archived State = "ARCHIVED"
)

func (td *Todo) String() string {
	return fmt.Sprintf("%#v", td)
}
