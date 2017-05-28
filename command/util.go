package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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
	initFileIfNotExist(path)
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
	initFileIfNotExist(path)
	indent := strings.Repeat(" ", 4)
	json, err := json.MarshalIndent(todos, "", indent)
	if err != nil {
		return fmt.Errorf("Failed to encode data: %v", err)
	}

	ioutil.WriteFile(path, json, 0644)

	return nil
}

func initFileIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	f, err := os.Create(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	_, err = f.Write([]byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func splitOrder(order string) []int {
	orders := []int{}
	for _, id := range strings.Split(order, "-") {
		order, err := strconv.Atoi(id)
		if err == nil {
			orders = append(orders, order)
		}
	}
	return orders
}
