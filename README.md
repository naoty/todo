# todo [![Build Status](https://travis-ci.org/naoty/todo.svg?branch=master)](https://travis-ci.org/naoty/todo)

## Installation

```
$ go get github.com/naoty/todo
```

## Usage

### List

```
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[ ] 003: Publish a blog entry
```

`-m` or `--markdown` flag enables to list TODOs as task lists in markdown.

```
$ todo list -m
- [x] Learn Golang
- [ ] Make a todo management tool just for myself
- [ ] Publish a blog entry
```

`-u` or `--undone` flag enables to list only undone TODOs.

```
$ todo list -u
[ ] 002: Make a todo management tool just for myself
[ ] 003: Publish a blog entry
```

`-d` or `--done` flag enables to list only done TODOs.

```
$ todo list -d
[x] 001: Learn Golang
```

### Add

```
$ todo add Share the entry on Twitter
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[ ] 003: Publish a blog entry
[ ] 004: Share the entry on Twitter
```

### Delete

```
$ todo delete 3
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
```

### Move

```
$ todo move 1 2
$ todo list
[ ] 001: Make a todo management tool just for myself
[x] 002: Learn Golang
[ ] 003: Publish a blog entry
[ ] 004: Share the entry on Twitter
```

### Rename

```
$ todo rename 4 Share the entry on Twitter and Facebook
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[ ] 003: Publish a blog entry
[ ] 004: Share the entry on Twitter and Facebook
```

### Done

```
$ todo done 2 3
$ todo list
[x] 001: Learn Golang
[x] 002: Make a todo management tool just for myself
[x] 003: Publish a blog entry
```

### Undone

```
$ todo undone 1 2
$ todo list
[ ] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[x] 003: Publish a blog entry
```

### Clear done TODOs

```
$ todo done 2
$ todo clear
$ todo list
[ ] 001: Learn Golang
```

## Configuration

```
TODO_PATH: Directory where .todo file saved (Default: HOME)
```

## Author

[naoty](https://github.com/naoty)

