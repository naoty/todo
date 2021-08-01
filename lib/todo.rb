module Todo
  VERSION = "0.5.0"

  autoload :Add, "todo/add"
  autoload :Archive, "todo/archive"
  autoload :CLI, "todo/cli"
  autoload :Delete, "todo/delete"
  autoload :FileRepository, "todo/file_repository"
  autoload :List, "todo/list"
  autoload :Move, "todo/move"
  autoload :Todo, "todo/todo"
  autoload :Update, "todo/update"
end
