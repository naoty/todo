# todo [![Build Status](https://travis-ci.org/naoty/todo.svg?branch=master)](https://travis-ci.org/naoty/todo)

## Installation

```
$ go get github.com/naoty/todo
```

## Usage

### List

```
$ todo list
[ ] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
```

### Add/Delete

```
$ todo add Get things done
$ todo list
[ ] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[ ] 003: Get things done
```

```
$ todo delete 3
$ todo list
[ ] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
```

### Done/Undone

```
$ todo done 1 2
$ todo list
[x] 001: Learn Golang
[x] 002: Make a todo management tool just for myself
```

```
$ todo undone 2
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
```

### Move

```
$ todo move 2 1
$ todo list
[ ] 001: Make a todo management tool just for myself
[x] 002: Learn Golang
```

### Clear

```
$ todo clear
$ todo list
[ ] 001: Make a todo management tool just for myself
```

### Subtodo

```
$ todo add -p 1 Try 'A Tour of Go'
$ todo add -p 1 Read 'Effective Go'
$ todo list
[ ] 001: Learn Golang
  [ ] 001: Try 'A Tour of Go'
  [ ] 002: Read 'Effective Go'
[ ] 002: Make a todo management tool just for myself
```

```
$ todo done 1-1
$ todo list
[ ] 001: Learn Golang
  [x] 001: Try 'A Tour of Go'
  [ ] 002: Read 'Effective Go'
[ ] 002: Make a todo management tool just for myself
```

## Configuration

```
TODO_PATH: Directory where .todo.json file saved (Default: HOME)
```

## Author

[naoty](https://github.com/naoty)
