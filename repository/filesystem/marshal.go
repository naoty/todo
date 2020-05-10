package filesystem

import (
	"fmt"
	"strings"

	"github.com/naoty/todo/todo"
)

// Marshal converts a TODO into byte slice.
func Marshal(td *todo.Todo) []byte {
	var lines []string
	lines = append(lines, "---")
	lines = append(lines, fmt.Sprintf("title: %s", td.Title))

	switch td.State {
	case todo.Undone:
		lines = append(lines, "state: undone")
	case todo.Done:
		lines = append(lines, "state: done")
	case todo.Waiting:
		lines = append(lines, "state: waiting")
	case todo.Archived:
		lines = append(lines, "state: archived")
	}

	lines = append(lines, "---")
	lines = append(lines, td.Body)
	text := strings.Join(lines, "\n")

	return []byte(text)
}
