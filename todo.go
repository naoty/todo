package main

import (
	"bytes"
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
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	err = removeFile()
	if err != nil {
		return err
	}

	index := num - 1
	todos = append(todos[:index], todos[index+1:]...)
	return WriteTodos(todos)
}

func MoveTodo(from, to int) error {
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	err = removeFile()
	if err != nil {
		return err
	}

	fromIndex, toIndex := from-1, to-1
	movedTodo := todos[fromIndex]
	todos = append(todos[:fromIndex], todos[fromIndex+1:]...)
	todos = append(todos[:toIndex], append([]Todo{movedTodo}, todos[toIndex:]...)...)
	return WriteTodos(todos)
}

func RenameTodo(num int, title string) error {
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	err = removeFile()
	if err != nil {
		return err
	}

	index := num - 1
	todos[index].Title = title
	return WriteTodos(todos)
}

func DoneTodo(num int) error {
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	err = removeFile()
	if err != nil {
		return err
	}

	newTodos := make([]Todo, len(todos))
	index := num - 1
	for i, todo := range todos {
		if i == index {
			todo.Done = true
		}
		newTodos[i] = todo
	}
	return WriteTodos(newTodos)
}

func UndoneTodo(num int) error {
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	err = removeFile()
	if err != nil {
		return err
	}

	newTodos := make([]Todo, len(todos))
	index := num - 1
	for i, todo := range todos {
		if i == index {
			todo.Done = false
		}
		newTodos[i] = todo
	}
	return WriteTodos(newTodos)
}

func ClearTodos() error {
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	err = removeFile()
	if err != nil {
		return err
	}

	var newTodos []Todo
	for _, todo := range todos {
		if !todo.Done {
			newTodos = append(newTodos, todo)
		}
	}
	return WriteTodos(newTodos)
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

func removeFile() error {
	path := getTodosPath()
	err := os.Remove(path)
	return err
}
