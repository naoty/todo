class Todo::CLI
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      blog -h | --help
      blog -v | --version
    
    Options:
      -h --help     Show this message
      -v --version  Show version
  TEXT

  private attr_reader :arguments, :output, :error_output

  def initialize(arguments: ARGV, output: $stdin, error_output: $stderr)
    @arguments = arguments
    @output = output
    @error_output = error_output
  end

  def run
    if arguments.empty?
      error_output.puts(HELP_MESSAGE)
      exit 1
    end

    if arguments.first == "-h" || arguments.first == "--help"
      output.puts(HELP_MESSAGE)
      return
    end

    if arguments.first == "-v" || arguments.first == "--version"
      output.puts(Todo::VERSION)
      return
    end

    raise NotImplementedError
  end
end
