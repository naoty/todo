class Todo::List
  include Todo::Printable

  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo list
      todo list -h | --help
    
    Options:
      -h --help  Show this message
  TEXT

  private attr_reader :arguments, :output, :error_output

  def initialize(arguments:, output: $stdout, error_output: $stderr)
    @arguments = arguments
    @output = output
    @error_output = error_output
  end

  def run(repository:)
    if arguments.first == "-h" || arguments.first == "--help"
      output.puts(HELP_MESSAGE)
      return
    end

    todos = repository.list
    print_todos(todos, indent_width: 2)
  end
end
