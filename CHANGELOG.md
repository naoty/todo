# Change Log

## Unreleased

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

## Added
* `next` command to show a next undone todo.

## Changed
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
