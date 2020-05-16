# todo

## Installation

```
$ brew tap naoty/misc
$ brew install todo
```

## Usage

```bash
# Add a TODO or a sub-TODO
# Each TODO is saved as a text file at `$TODOS_PATH/<id>.md`.
% todo add Learn Golang
% todo add Write some tool in Golang
% todo add -p 1 Try A tour of Go

# Show TODOs
% todo list
[ ] 001: Learn Golang
  [ ] 003: Try A tour of Go
[ ] 002: Write some tool in Golang

# Open a TODO file with `open` command
% todo open 3

# Mark a TODO as done
% todo done 3
% todo list
[ ] 001: Learn Golang
  [x] 003: Try A tour of Go
[ ] 002: Write some tool in Golang

# Delete a TODO file
% todo delete 2
% todo list
[ ] 001: Learn Golang
  [x] 003: Try A tour of Go

# Archive done TODO files
# This doesn't delete TODO files but hides done TODOs from list
% todo archive
% todo list
[ ] 001: Learn Golang
```

## Environment variables
* `TODOS_PATH`: The root path of TODO files (default: `$HOME/.todos`)

## Author

[naoty](https://github.com/naoty)
