class Todo::Commands::Archive < Todo::Commands::Command
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo archive
      todo archive -h | --help
    
    Options:
      -h --help  Show thid message
  TEXT

  def run(repository:)
    if arguments.first == "-h" || arguments.first == "--help"
      output.puts(HELP_MESSAGE)
      return
    end

    repository.archive

    todos = repository.list
    print_todos(todos, indent_width: 2)
  end
end
