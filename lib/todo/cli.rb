class Todo::CLI
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      blog -h | --help
    
    Options:
      -h --help  Show this message
  TEXT

  private attr_reader :arguments, :output, :error_output

  def initialize(arguments: ARGV, output: $stdin, error_output: $stderr)
    @arguments = arguments
    @output = output
    @error_output = error_output
  end

  def run
    if arguments.first == "-h" || arguments.first == "--help"
      output.puts(HELP_MESSAGE)
      return
    end

    raise NotImplementedError
  end
end
