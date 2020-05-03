package todo

import (
	"errors"
	"strings"

	"gopkg.in/yaml.v2"
)

const separator = "---"

type metadata struct {
	Title string
}

// Parse parses given text into a *Todo.
func Parse(text string) (*Todo, error) {
	frontmatter, _, err := splitFrontmatter(text)
	if err != nil {
		return nil, err
	}

	var meta metadata
	err = yaml.Unmarshal([]byte(frontmatter), &meta)
	if err != nil {
		return nil, err
	}

	return &Todo{id: 0, title: meta.Title}, nil
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
