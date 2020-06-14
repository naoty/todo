package repository

import "fmt"

// NotFoundError represents an error occured when tha path is not found.
type NotFoundError struct {
	Path string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("file not found: %s", e.Path)
}
