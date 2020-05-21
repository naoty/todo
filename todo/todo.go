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
)

func (td *Todo) String() string {
	return fmt.Sprintf("%#v", td)
}

// Equal compares self and another TODO deeply.
func (td *Todo) Equal(another *Todo) bool {
	if another == nil {
		return false
	}

	if td.ID != another.ID {
		return false
	}

	if td.Title != another.Title {
		return false
	}

	if td.State != another.State {
		return false
	}

	if td.Body != another.Body {
		return false
	}

	if len(td.Todos) != len(another.Todos) {
		return false
	}

	for i, sub := range td.Todos {
		equal := sub.Equal(another.Todos[i])
		if !equal {
			return false
		}
	}

	return true
}
