package todo

import (
	"io"

	"github.com/ymotongpoo/goltsv"
)

type Decoder struct {
	reader io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{reader: r}
}

func (d *Decoder) Decode() ([]Todo, error) {
	reader := goltsv.NewReader(d.reader)
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
