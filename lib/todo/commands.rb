module Todo::Commands
  autoload :Command, "todo/commands/command"

  autoload :Add, "todo/commands/add"
  autoload :Archive, "todo/commands/archive"
  autoload :CLI, "todo/commands/cli"
  autoload :Delete, "todo/commands/Delete"
  autoload :List, "todo/commands/list"
  autoload :Move, "todo/commands/move"
  autoload :Open, "todo/commands/open"
  autoload :Printable, "todo/commands/printable"
  autoload :Update, "todo/commands/update"
end
