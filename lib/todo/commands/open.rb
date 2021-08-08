class Todo::Commands::Open < Todo::Commands::Command
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo open <id>
      todo open -h | --help
    
    Options:
      -h --help  Show this message
  TEXT

  def run(repository:)
    if arguments.empty?
      error_output.puts(HELP_MESSAGE)
      exit 1
    end

    if arguments.first == "-h" || arguments.first == "--help"
      output.puts(HELP_MESSAGE)
      return
    end

    repository.open(id: arguments.first.to_i)
  end
end
