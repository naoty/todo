package main

type Todo struct {
	Number int
	Title  string
	Done   bool
}

func (todo Todo) Encode() map[string]string {
	m := make(map[string]string)
	m["title"] = todo.Title
	if todo.Done {
		m["done"] = "true"
	} else {
		m["done"] = "false"
	}
	return m
}

type process func([]Todo) ([]Todo, error)
