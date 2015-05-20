# todo 

## TODOs of todo

- [x] Write README
- [x] List
- [ ] Add
- [ ] Delete
- [ ] Done
- [ ] Undone
- [ ] Clear

## Usage

### List

```bash
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[ ] 003: Publish a blog entry
```

### Add

```bash
$ todo add Share the entry on Twitter
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[ ] 003: Publish a blog entry
[ ] 004: Share the entry on Twitter
```

### Delete

```bash
$ todo delete 3
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
```

### Done

```bash
$ todo done 3
$ todo list
[x] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[x] 003: Publish a blog entry
```

### Undone

```bash
$ todo undone 1
$ todo list
[ ] 001: Learn Golang
[ ] 002: Make a todo management tool just for myself
[ ] 003: Publish a blog entry
```

### Clear done TODOs

```bash
$ todo done 2
$ todo clear
$ todo list
[ ] 003: Publish a blog entry
```

## Author

[naoty](https://github.com/naoty)

