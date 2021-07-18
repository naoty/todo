class Todo::List
  private attr_reader :arguments, :output, :error_output

  def initialize(arguments:, output: $stdout, error_output: $stderr)
    @arguments = arguments
    @output = output
    @error_output = error_output
  end

  def run(repository:)
    raise NotImplementedError
  end
end
