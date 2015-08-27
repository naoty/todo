package todo

import (
	"log"
	"os"
	"path/filepath"
)

type File struct {
	path string
}

type TodoProcess func([]Todo) ([]Todo, error)

func OpenFile() *File {
	return &File{}
}

func (f *File) Path() string {
	if f.path == "" {
		path := os.Getenv("TODO_PATH")
		if path == "" {
			path = os.Getenv("HOME")
		}
		f.path = filepath.Join(path, ".todo")
	}

	return f.path
}

func (f *File) Update(process TodoProcess) error {
	todos, err := f.Read()
	if err != nil {
		return err
	}

	path := f.Path()
	err = os.Remove(path)
	if err != nil {
		return err
	}

	newTodos, err := process(todos)
	if err != nil {
		log.Println(err)
		return f.Write(todos)
	}

	return f.Write(newTodos)
}

func (f *File) Read() ([]Todo, error) {
	path := f.Path()

	if !f.isExist() {
		_, err := os.Create(path)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := NewDecoder(file)
	todos, err := decoder.Decode()

	return todos, err
}

func (f *File) Write(todos []Todo) error {
	path := f.Path()

	if !f.isExist() {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := NewEncoder(file)
	err = encoder.Encode(todos)

	return err
}

func (f *File) HasTodo(title string) bool {
	todos, err := f.Read()
	if err != nil {
		log.Println(err)
		return false
	}

	for _, todo := range todos {
		if todo.Title == title {
			return true
		}
	}

	return false
}

func (f *File) isExist() bool {
	path := f.Path()
	_, err := os.Stat(path)
	return err == nil
}
