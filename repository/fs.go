package repository

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/naoty/todo/todo"
)

// FS represents a repository backed by file system.
type FS struct {
	root string
}

// NewFS returns a new FS.
func NewFS(root string) *FS {
	return &FS{root: root}
}

// List implements Repository interface.
func (repo *FS) List() ([]*todo.Todo, error) {
	todos := []*todo.Todo{}

	err := filepath.Walk(repo.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		id, err := parseID(path)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		td, err := todo.Parse(string(content))
		if err != nil {
			return err
		}

		td.SetID(id)
		todos = append(todos, td)

		return nil
	})

	if err != nil {
		return todos, fmt.Errorf("failed to get todos from %s: %w", repo.root, err)
	}

	return todos, nil
}

func parseID(path string) (int, error) {
	text := strings.TrimRight(filepath.Base(path), filepath.Ext(path))
	return strconv.Atoi(text)
}
