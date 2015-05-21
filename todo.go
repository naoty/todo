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
