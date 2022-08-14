module Todo
  VERSION = "0.5.1"

  autoload :Commands, "todo/commands"
  autoload :FileRepository, "todo/file_repository"
  autoload :Todo, "todo/todo"
end
