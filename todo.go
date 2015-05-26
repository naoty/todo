package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ymotongpoo/goltsv"
)

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

func ReadTodos() ([]Todo, error) {
	if !fileIsExist() {
		err := createNewFile()
		if err != nil {
			return nil, err
		}
	}

	path := getTodosPath()
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(data)
	reader := goltsv.NewReader(buf)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	todos := []Todo{}
	for i, record := range records {
		var title string
		var done bool

		for k, v := range record {
			switch k {
			case "title":
				title = v
			case "done":
				done = (v == "true")
			}
		}

		todo := Todo{Number: i + 1, Title: title, Done: done}
		todos = append(todos, todo)
	}

	return todos, nil
}

func WriteTodos(todos []Todo) error {
	var data []map[string]string
	for _, todo := range todos {
		data = append(data, todo.Encode())
	}

	if !fileIsExist() {
		err := createNewFile()
		if err != nil {
			return err
		}
	}

	path := getTodosPath()
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	writer := goltsv.NewWriter(file)

	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}

func AppendTodo(todo Todo) error {
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	todos = append(todos, todo)
	return WriteTodos(todos)
}

func DeleteTodo(num int) error {
	return rewriteFile(func(todos []Todo) ([]Todo, error) {
		index := num - 1
		if index >= len(todos) {
			return nil, errors.New("Index out of bounds.")
		}

		return append(todos[:index], todos[index+1:]...), nil
	})
}

func MoveTodo(from, to int) error {
	return rewriteFile(func(todos []Todo) ([]Todo, error) {
		fromIndex, toIndex := from-1, to-1
		if fromIndex >= len(todos) || toIndex >= len(todos) {
			return nil, errors.New("Index out of bounds.")
		}

		movedTodo := todos[fromIndex]
		todos = append(todos[:fromIndex], todos[fromIndex+1:]...)
		todos = append(todos[:toIndex], append([]Todo{movedTodo}, todos[toIndex:]...)...)
		return todos, nil
	})
}

func RenameTodo(num int, title string) error {
	return rewriteFile(func(todos []Todo) ([]Todo, error) {
		index := num - 1
		if index >= len(todos) {
			return nil, errors.New("Index out of bounds.")
		}

		todos[index].Title = title
		return todos, nil
	})
}

func DoneTodo(nums ...int) error {
	return rewriteFile(func(todos []Todo) ([]Todo, error) {
		var err error
		var indices []int
		for _, num := range nums {
			index := num - 1
			if index >= len(todos) {
				err = errors.New("Index out of bounds.")
			}
			indices = append(indices, index)
		}
		if err != nil {
			return nil, err
		}

		newTodos := make([]Todo, len(todos))
		for i, todo := range todos {
			if contains(indices, i) {
				todo.Done = true
			}
			newTodos[i] = todo
		}
		return newTodos, nil
	})
}

func UndoneTodo(nums ...int) error {
	return rewriteFile(func(todos []Todo) ([]Todo, error) {
		var err error
		var indices []int
		for _, num := range nums {
			index := num - 1
			if index >= len(todos) {
				err = errors.New("Index out of bounds.")
			}
			indices = append(indices, index)
		}
		if err != nil {
			return nil, err
		}

		newTodos := make([]Todo, len(todos))
		for i, todo := range todos {
			if contains(indices, i) {
				todo.Done = false
			}
			newTodos[i] = todo
		}
		return newTodos, nil
	})
}

func ClearTodos() error {
	return rewriteFile(func(todos []Todo) ([]Todo, error) {
		var newTodos []Todo
		for _, todo := range todos {
			if !todo.Done {
				newTodos = append(newTodos, todo)
			}
		}
		return newTodos, nil
	})
}

func getTodosPath() string {
	path := os.Getenv("TODO_PATH")
	if path == "" {
		path = os.Getenv("HOME")
	}

	return filepath.Join(path, ".todo")
}

func fileIsExist() bool {
	path := getTodosPath()
	_, err := os.Stat(path)
	return err == nil
}

func createNewFile() error {
	path := getTodosPath()
	_, err := os.Create(path)
	return err
}

func rewriteFile(f func([]Todo) ([]Todo, error)) error {
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	err = removeFile()
	if err != nil {
		return err
	}

	newTodos, err := f(todos)
	if err != nil {
		// Recover removed todos
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return WriteTodos(todos)
	}

	return WriteTodos(newTodos)
}

func removeFile() error {
	path := getTodosPath()
	err := os.Remove(path)
	return err
}

func contains(xs []int, n int) bool {
	for _, x := range xs {
		if x == n {
			return true
		}
	}
	return false
}
