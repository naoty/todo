module Todo
  VERSION = "0.5.0"

  autoload :Commands, "todo/commands"
  autoload :FileRepository, "todo/file_repository"
  autoload :Todo, "todo/todo"
end
