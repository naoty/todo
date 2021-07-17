class Todo::Add
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo add <title>
      todo add -h | --help
    
    Options:
      -h --help  Show thid message
  TEXT

  private attr_reader :arguments, :output, :error_output

  def initialize(arguments:, output: $stdout, error_output: $stderr)
    @arguments = arguments
    @output = output
    @error_output = error_output
  end

  def run(repository:)
    if arguments.empty?
      error_output.puts(HELP_MESSAGE)
      exit 1
    end

    if arguments.first == "-h" || arguments.first == "--help"
      output.puts(HELP_MESSAGE)
      return
    end

    repository.create(title: arguments.first)
  end
end
