package repository

import "github.com/naoty/todo/todo"

// Repository represents an interface to get and set TODOs.
type Repository interface {
	List() ([]*todo.Todo, error)
	Add(title string) error
}
