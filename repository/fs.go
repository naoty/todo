package repository

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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

// Add implements Repository interface.
func (repo *FS) Add(title string) error {
	todos, err := repo.List()
	if err != nil {
		return fmt.Errorf("failed to get next id: %w", err)
	}

	lastID := 0
	for _, td := range todos {
		id := td.ID()
		if id > lastID {
			lastID = id
		}
	}

	nextID := lastID + 1
	filename := fmt.Sprintf("%d.md", nextID)
	path := filepath.Join(repo.root, filename)

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("already exist: %s", path)
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to add TODO: %w", err)
	}
	defer file.Close()

	content := newContent(title)
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to add TODO: %w", err)
	}

	return nil
}

// Open implements Repository interface.
func (repo *FS) Open(id int) error {
	filename := fmt.Sprintf("%d.md", id)
	path := filepath.Join(repo.root, filename)

	cmd := exec.Command("open", path)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", path, err)
	}

	return nil
}

func parseID(path string) (int, error) {
	text := strings.TrimRight(filepath.Base(path), filepath.Ext(path))
	return strconv.Atoi(text)
}

func newContent(title string) string {
	return strings.TrimLeft(fmt.Sprintf(`
---
title: %s
---


`, title), "\n")
}
