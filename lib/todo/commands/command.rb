class Todo::Commands::Command
  include Todo::Commands::Printable

  attr_reader :arguments, :output, :error_output

  def initialize(arguments: ARGV, output: $stdout, error_output: $stderr)
    @arguments = arguments
    @output = output
    @error_output = error_output
  end

  def run
    raise NotImplementedError, "this method must be overwritten by subclass"
  end
end
