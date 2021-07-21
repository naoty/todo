module Todo
  VERSION = "0.5.0"

  autoload :Add, "todo/add"
  autoload :CLI, "todo/cli"
  autoload :Delete, "todo/delete"
  autoload :FileRepository, "todo/file_repository"
  autoload :List, "todo/list"
  autoload :Todo, "todo/todo"
end
