# Change Log

## Unreleased

### Fixed
* Double-quote titles with square brackets or colons to decode them correctly.
* Fix default path for TODOs from current directory to $HOME/.todos.

## 0.5.0 - 2021-08-07

### Added
* `--open`, `-o` flag is added to `todo add` to open a TODO file after creating it.
* `<position>` argument is added to `todo add` to create a TODO at given position.

### Changed
* All codes are rewrited in Ruby.
* `list` command shows TODOs in a new format.
* `done`, `undone`, `wait` command updates the state of subtodos recursively.

## 0.4.2 - 2020-06-14

### Changed
* `archive` command continues to archive all done TODOs even if some TODO files are not found.

## 0.4.1

### Added
* `archive` command archives all done sub-TODOs of undone TODOs.

### Changed
* Remove all IDs from index.json when there are not corresponding files under `$TODOS_PATH`.

## 0.4.0

### Changed
* Rewrite all codes from scratch.

## 0.3.1

### Fixed
* Fix `next` command to show a todo when the subtodos of the todo are all done.

## 0.3.0

### Added
* `next` command to show a next undone todo.

### Changed
* `done` command without any orders marks a next undone todo as done.

## 0.2.1

### Fixed
* Fix a bug to delete multiple TODOs.

## 0.2.0

### Added
* Support subtodos.

### Changed
* Change the name of the file where todos saved from `.todo` to `.todo.json`.
* Change the format of the file where todos saved from LTSV to JSON to support subtodos.
