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

type index struct {
	Todos map[string][]int `json:"todos"`
}

// New returns a new FileSystem.
func New() (*FileSystem, error) {
	root := os.Getenv("TODOS_PATH")
	if root == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		root = filepath.Join(home, ".todos")
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		err := os.Mkdir(root, 0755)
		if err != nil {
			return nil, err
		}
	}

	return &FileSystem{root: root}, nil
}

// Get implements Repository interface.
func (repo *FileSystem) Get(id int) (*todo.Todo, error) {
	filename := fmt.Sprintf("%d.md", id)
	path := filepath.Join(repo.root, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("TODO not found: %d", id)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get TODO: %w", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to get TODO: %w", err)
	}

	td, err := Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("failed to get TODO: %w", err)
	}

	td.ID = id
	return td, nil
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

		td, err := Parse(string(content))
		if err != nil {
			return err
		}

		td.ID = id
		todos[id] = td

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get todos from %s: %w", repo.root, err)
	}

	i, err := repo.readIndex()
	if err != nil {
		return nil, fmt.Errorf("failed to get todos from %s: %w", repo.root, err)
	}

	ordered := orderTodos(todos, i, "")

	return ordered, nil
}

// Add implements Repository interface.
func (repo *FileSystem) Add(title string, parent *int) error {
	todos, err := repo.List()
	if err != nil {
		return fmt.Errorf("failed to get next id: %w", err)
	}

	if *parent > 0 && !isExists(todos, *parent) {
		return fmt.Errorf("parent not found: %d", *parent)
	}

	nextID := maxID(todos) + 1
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

	st, err := repo.readIndex()
	if err != nil {
		return fmt.Errorf("failed to add TODO: %w", err)
	}

	parentID := ""
	if *parent > 0 {
		parentID = fmt.Sprintf("%d", *parent)
	}

	st.Todos[parentID] = append(st.Todos[parentID], nextID)
	err = repo.writeIndex(st)
	if err != nil {
		return fmt.Errorf("failed to add TODO: %w", err)
	}

	return nil
}

// Update implements Repository interface.
func (repo *FileSystem) Update(td *todo.Todo) error {
	filename := fmt.Sprintf("%d.md", td.ID)
	path := filepath.Join(repo.root, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", path)
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to update TODO: %w", err)
	}
	defer file.Close()

	data := Marshal(td)
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to update TODO: %w", err)
	}

	return nil
}

// Delete implements Repository interface.
func (repo *FileSystem) Delete(id int) error {
	filename := fmt.Sprintf("%d.md", id)
	path := filepath.Join(repo.root, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", path)
	}

	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("failed to delete TODO: %w", err)
	}

	index, err := repo.readIndex()
	if err != nil {
		return fmt.Errorf("failed to delete TODO from index: %w", err)
	}

	key := fmt.Sprintf("%d", id)
	subIDs, ok := index.Todos[key]

	if ok {
		for _, id := range subIDs {
			filename := fmt.Sprintf("%d.md", id)
			path := filepath.Join(repo.root, filename)

			if _, err := os.Stat(path); os.IsNotExist(err) {
				return fmt.Errorf("file not found: %s", path)
			}

			err := os.Remove(path)
			if err != nil {
				return fmt.Errorf("failed to delete TODO: %w", err)
			}
		}
	}

	delete(index.Todos, key)

	for k, todos := range index.Todos {
		var ids []int
		for _, _id := range todos {
			if _id == id {
				continue
			}
			ids = append(ids, _id)
		}

		if len(ids) == 0 {
			delete(index.Todos, k)
			continue
		}

		index.Todos[k] = ids
	}

	err = repo.writeIndex(index)
	if err != nil {
		return fmt.Errorf("failed to delete TODO from index: %w", err)
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

	st, err := repo.readIndex()
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

	err = repo.writeIndex(st)
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

func (repo *FileSystem) readIndex() (*index, error) {
	path := filepath.Join(repo.root, "index.json")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &index{
			Todos: map[string][]int{
				"": {},
			},
		}, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read index from %s: %w", path, err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read index from %s: %w", path, err)
	}

	var i index
	err = json.Unmarshal(content, &i)
	if err != nil {
		return nil, fmt.Errorf("failed to read index from %s: %w", path, err)
	}

	return &i, nil
}

func (repo *FileSystem) writeIndex(i *index) error {
	path := filepath.Join(repo.root, "index.json")

	data, err := json.MarshalIndent(*i, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to write index to %s: %w", path, err)
	}

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write index to %s: %w", path, err)
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

func maxID(todos []*todo.Todo) int {
	result := 0

	for _, td := range todos {
		if td.ID > result {
			result = td.ID
		}

		maxSubID := maxID(td.Todos)
		if maxSubID > result {
			result = maxSubID
		}
	}

	return result
}

func orderTodos(todos map[int]*todo.Todo, i *index, key string) []*todo.Todo {
	var result []*todo.Todo

	for _, id := range i.Todos[key] {
		td, ok := todos[id]
		if !ok {
			continue
		}

		newKey := fmt.Sprintf("%d", td.ID)
		td.Todos = orderTodos(todos, i, newKey)
		result = append(result, td)
	}

	return result
}

func isExists(todos []*todo.Todo, id int) bool {
	for _, td := range todos {
		if td.ID == id {
			return true
		}

		result := isExists(td.Todos, id)
		if result {
			return true
		}
	}

	return false
}
