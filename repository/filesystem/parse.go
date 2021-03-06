package filesystem

import (
	"errors"
	"strings"

	"github.com/naoty/todo/todo"
	"gopkg.in/yaml.v2"
)

const separator = "---"

type frontmatter struct {
	Title string
	State string
}

// Parse parses given text into a *Todo.
func Parse(text string) (*todo.Todo, error) {
	fmText, body, err := splitFrontmatter(text)
	if err != nil {
		return nil, err
	}

	var fm frontmatter
	err = yaml.Unmarshal([]byte(fmText), &fm)
	if err != nil {
		return nil, err
	}

	state := todo.Undone
	switch fm.State {
	case "done":
		state = todo.Done
	case "waiting":
		state = todo.Waiting
	}

	return &todo.Todo{ID: 0, Title: fm.Title, State: state, Body: body}, nil
}

func splitFrontmatter(text string) (string, string, error) {
	lines := strings.Split(text, "\n")

	if len(lines) == 0 || lines[0] != separator {
		return "", "", errors.New("frontmatter not found")
	}

	isSeparated := false
	frontmatterLines := []string{}
	bodyLines := []string{}
	for _, line := range lines[1:] {
		if line == separator {
			isSeparated = true
			continue
		}

		if isSeparated {
			bodyLines = append(bodyLines, line)
			continue
		}

		frontmatterLines = append(frontmatterLines, line)
	}

	if !isSeparated {
		return "", "", errors.New("frontmatter not found")
	}

	frontmatter := strings.Join(frontmatterLines, "\n")
	body := strings.Join(bodyLines, "\n")
	return frontmatter, body, nil
}
