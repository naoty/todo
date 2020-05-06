package todo

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
