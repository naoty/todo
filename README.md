# todo

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

### Done

```
$ todo done 3
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[x] 003: Publish a blog entry
```

### Undone

```
$ todo undone 1
$ todo list
[ ] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[ ] 003: Publish a blog entry
```

### Clear done TODOs

```
$ todo done 2
$ todo clear
$ todo list
[ ] 003: Publish a blog entry
```

## Configuration

```
TODO_PATH: Directory where .todo file saved (Default: HOME)
```

## Author

[naoty](https://github.com/naoty)

