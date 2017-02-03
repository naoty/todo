package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/naoty/todo/todo"
)

func todoFilePath() string {
	dir := os.Getenv("TODO_PATH")
	if dir == "" {
		dir = os.Getenv("HOME")
	}
	path := filepath.Join(dir, ".todo.json")
	return path
}

func readTodos(path string) ([]todo.Todo, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read data: %v", err)
	}

	var todos []todo.Todo
	if err = json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("Failed to decode data: %v", err)
	}
	return todos, nil
}

func writeTodos(todos []todo.Todo, path string) error {
	indent := strings.Repeat(" ", 4)
	json, err := json.MarshalIndent(todos, "", indent)
	if err != nil {
		return fmt.Errorf("Failed to encode data: %v", err)
	}

	ioutil.WriteFile(path, json, 0644)

	return nil
}
