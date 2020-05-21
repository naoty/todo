package mock

import (
	"errors"
	"fmt"

	"github.com/naoty/todo/todo"
)

// Mock represents a mock repository for TODOs.
type Mock struct {
	todos []*todo.Todo
}

// New returns a new Mock.
func New(todos []*todo.Todo) *Mock {
	return &Mock{todos: todos}
}

// Get implements Repository interface.
func (repo *Mock) Get(id int) (*todo.Todo, error) {
	for _, td := range repo.todos {
		if td.ID == id {
			return td, nil
		}

		for _, sub := range td.Todos {
			td, err := repo.Get(sub.ID)
			if td != nil {
				return td, nil
			}

			if err != nil {
				return nil, err
			}
		}
	}

	return nil, fmt.Errorf("TODO not found: %d", id)
}

// List implements Repository interface.
func (repo *Mock) List() ([]*todo.Todo, error) {
	return repo.todos, nil
}

// Add implements Repository interface.
func (repo *Mock) Add(title string, parent *int) error {
	td := &todo.Todo{
		ID:    0,
		Title: title,
		State: todo.Undone,
		Body:  "",
		Todos: []*todo.Todo{},
	}

	// TODO: support adding a sub-TODO

	repo.todos = append(repo.todos, td)

	return nil
}

// Update implements Repository interface.
func (repo *Mock) Update(td *todo.Todo) error {
	index := -1

	// TODO: support updating a sub-TODO

	for i, _td := range repo.todos {
		if _td.ID == td.ID {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("TODO not found: %d", td.ID)
	}

	repo.todos = append(repo.todos[:index], append([]*todo.Todo{td}, repo.todos[index+1:]...)...)

	return nil
}

// Delete implements Repository interface.
func (repo *Mock) Delete(id int) error {
	index := -1

	// TODO: support updating a sub-TODO

	for i, _td := range repo.todos {
		if _td.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("TODO not found: %d", id)
	}

	repo.todos = append(repo.todos[:index], repo.todos[index+1:]...)

	return nil
}

// Open implements Repository interface.
func (repo *Mock) Open(id int) error {
	return errors.New("mock repository doesn't support openning a TODO")
}

// Move implements Repository interface.
func (repo *Mock) Move(id int, parent *int, position int) error {
	var td *todo.Todo
	from := -1
	to := position - 1

	// TODO: support updating a sub-TODO

	for i, _td := range repo.todos {
		if _td.ID == id {
			td = _td
			from = i
			break
		}
	}

	if td == nil {
		return fmt.Errorf("TODO not found: %d", id)
	}

	repo.todos = append(repo.todos[:from], repo.todos[from+1:]...)
	repo.todos = append(repo.todos[:to], append([]*todo.Todo{td}, repo.todos[to:]...)...)

	return nil
}

// Archive implements Repository interface.
func (repo *Mock) Archive(id int) error {
	return errors.New("mock repository doesn't support archiving TODOs")
}
