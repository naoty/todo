package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ymotongpoo/goltsv"
)

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

func UpdateTodos(p process) error {
	todos, err := ReadTodos()
	if err != nil {
		return err
	}

	err = removeFile()
	if err != nil {
		return err
	}

	newTodos, err := p(todos)
	if err != nil {
		// Recover removed todos
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return WriteTodos(todos)
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
