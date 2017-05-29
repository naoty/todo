package todo

// Todo represents a todo.
type Todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Todos []Todo `json:"todos"`

	// Marked to be deleted
	Trashed bool `json:"trashed"`
}
