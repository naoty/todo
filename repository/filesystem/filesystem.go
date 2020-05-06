package filesystem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/naoty/todo/todo"
)

// FileSystem represents a repository backed by file system.
type FileSystem struct {
	root string
}

type metadata struct {
	Todos map[string][]int `json:"todos"`
}

// New returns a new FileSystem.
func New() (*FileSystem, error) {
	home := os.Getenv("TODOS_PATH")
	if home == "" {
		var err error
		home, err = os.UserHomeDir()
		if err != nil {
			return nil, err
		}
	}

	root := filepath.Join(home, ".todos")
	if _, err := os.Stat(root); os.IsNotExist(err) {
		err := os.Mkdir(root, 0755)
		if err != nil {
			return nil, err
		}
	}

	return &FileSystem{root: root}, nil
}

// List implements Repository interface.
func (repo *FileSystem) List() ([]*todo.Todo, error) {
	todos := map[int]*todo.Todo{}

	err := filepath.Walk(repo.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".md") {
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
		todos[id] = td

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get todos from %s: %w", repo.root, err)
	}

	st, err := repo.readMetadata()
	if err != nil {
		return nil, fmt.Errorf("failed to get todos from %s: %w", repo.root, err)
	}

	var ordered []*todo.Todo
	for k, ids := range st.Todos {
		// NOTE: when TODOs have subTODOs, k is parent TODO's ID.
		if k != "" {
			continue
		}

		for _, id := range ids {
			td, ok := todos[id]
			if !ok {
				continue
			}

			ordered = append(ordered, td)
		}
	}

	return ordered, nil
}

// Add implements Repository interface.
func (repo *FileSystem) Add(title string) error {
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

	st, err := repo.readMetadata()
	if err != nil {
		return fmt.Errorf("failed to add TODO: %w", err)
	}

	st.Todos[""] = append(st.Todos[""], nextID)
	err = repo.writeMetadata(st)
	if err != nil {
		return fmt.Errorf("failed to add TODO: %w", err)
	}

	return nil
}

// Open implements Repository interface.
func (repo *FileSystem) Open(id int) error {
	filename := fmt.Sprintf("%d.md", id)
	path := filepath.Join(repo.root, filename)

	cmd := exec.Command("open", path)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", path, err)
	}

	return nil
}

// Move implements Repository interface.
func (repo *FileSystem) Move(id, position int) error {
	if position < 1 {
		return fmt.Errorf("position number must be larger than 0: %d", position)
	}

	to := position - 1

	st, err := repo.readMetadata()
	if err != nil {
		return fmt.Errorf("failed to move TODO: %w", err)
	}

	ids, _ := st.Todos[""]
	if to >= len(ids) {
		return fmt.Errorf("position number is too large: %d", position)
	}

	var from int
	for j, _id := range ids {
		if _id == id {
			from = j
		}
	}

	st.Todos[""] = swapped(ids, from, to)

	err = repo.writeMetadata(st)
	if err != nil {
		return fmt.Errorf("failed to move TODO: %w", err)
	}

	return nil
}

func parseID(path string) (int, error) {
	text := strings.TrimRight(filepath.Base(path), filepath.Ext(path))
	return strconv.Atoi(text)
}

func newContent(title string) string {
	return strings.Trim(fmt.Sprintf(`
---
title: %s
state: undone
---
`, title), "\n")
}

func (repo *FileSystem) readMetadata() (*metadata, error) {
	path := filepath.Join(repo.root, "metadata.json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &metadata{
			Todos: map[string][]int{
				"": {},
			},
		}, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata from %s: %w", path, err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata from %s: %w", path, err)
	}

	var meta metadata
	err = json.Unmarshal(content, &meta)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata from %s: %w", path, err)
	}

	return &meta, nil
}

func (repo *FileSystem) writeMetadata(meta *metadata) error {
	path := filepath.Join(repo.root, "metadata.json")

	data, err := json.MarshalIndent(*meta, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to write metadata to %s: %w", path, err)
	}

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write metadata to %s: %w", path, err)
	}

	return nil
}

func swapped(s []int, from, to int) []int {
	if from < 0 || from >= len(s) {
		return s
	}

	if to < 0 || to >= len(s) {
		return s
	}

	copied := make([]int, len(s), cap(s))
	copy(copied, s)

	e := copied[from]
	copied = append(copied[:from], copied[(from+1):]...)
	copied = append(copied[:to], append([]int{e}, copied[to:]...)...)
	return copied
}
