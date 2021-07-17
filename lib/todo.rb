module Todo
  VERSION = "0.5.0"

  autoload :Add, "todo/add"
  autoload :CLI, "todo/cli"
  autoload :FileRepository, "todo/file_repository"
  autoload :Todo, "todo/todo"
end
