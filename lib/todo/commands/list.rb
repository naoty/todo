class Todo::Commands::List < Todo::Commands::Command
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo list
      todo list -h | --help
    
    Options:
      -h --help  Show this message
  TEXT

  def run(repository:)
    if arguments.first == "-h" || arguments.first == "--help"
      output.puts(HELP_MESSAGE)
      return
    end

    todos = repository.list
    print_todos(todos, indent_width: 2)
  end
end
