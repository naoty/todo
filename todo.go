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
	path := getTodosPath()
	if !fileIsExist(path) {
		err := createNewFile(path)
		if err != nil {
			return nil, err
		}
	}

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

func AppendTodo(todo Todo) error {
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	var data []map[string]string
	for _, todo := range todos {
		data = append(data, todo.Encode())
	}
	data = append(data, todo.Encode())

	path := getTodosPath()
	if !fileIsExist(path) {
		err := createNewFile(path)
		if err != nil {
			return err
		}
	}

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

func getTodosPath() string {
	path := os.Getenv("TODO_PATH")
	if path == "" {
		path = os.Getenv("HOME")
	}

	return filepath.Join(path, ".todo")
}

func fileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func createNewFile(path string) error {
	_, err := os.Create(path)
	return err
}
