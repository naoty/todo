package repository

import "github.com/naoty/todo/todo"

// FS represents a repository backed by file system.
type FS struct {
	root string
}

// NewFS returns a new FS.
func NewFS(root string) *FS {
	return &FS{root: root}
}

// List implements Repository interface.
func (repo *FS) List() []*todo.Todo {
	return nil
}
