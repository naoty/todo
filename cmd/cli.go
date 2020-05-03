package cmd

import "io"

// CLI represents an I/O against CLI.
type CLI struct {
	Reader      io.Reader
	Writer      io.Writer
	ErrorWriter io.Writer
}
