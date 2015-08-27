package command

import (
	"log"
	"os"

	"github.com/codegangsta/cli"

	"github.com/naoty/todo/formatter"
	"github.com/naoty/todo/todo"
)

var List = cli.Command{
	Name:  "list",
	Usage: "List TODOs",
	Action: func(context *cli.Context) {
		status := ExecList(context)
		os.Exit(status)
	},
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "undone, u",
			Usage: "Print only undone TODOs",
		},
		cli.BoolFlag{
			Name:  "done, d",
			Usage: "Print only done TODOs",
		},
		cli.BoolFlag{
			Name:  "markdown, m",
			Usage: "Print TODOs as task lists in markdown",
		},
	},
}

func ExecList(context *cli.Context) int {
	file := todo.OpenFile()

	var f formatter.Formatter
	mode := formatter.NewMode(context.Bool("done"), context.Bool("undone"))
	if context.Bool("markdown") {
		f = formatter.NewMarkdownFormatter(os.Stdout, mode)
	} else {
		f = formatter.NewNumberedFormatter(os.Stdout, mode)
	}

	todos, err := file.Read()
	if err != nil {
		log.Println(err)
		return 1
	}

	err = f.Print(todos)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
