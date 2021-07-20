class Todo::List
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo list
      todo -h | --help
    
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
    id_width = todos.map { |todo| todo.id.digits.length }.max

    todos.each do |todo|
      output.puts(format_todo(todo, id_width: id_width))
    end
  end

  private

  def format_todo(todo, id_width:)
    indent = " " * 2
    right_aligned_id = todo.id.to_s.rjust(id_width, " ")
    decorated_title =
      case todo.state
      when :undone then todo.title
      when :waiting then "\e[2m#{todo.title}\e[0m" # dim
      when :done then "\e[2;9m#{todo.title}\e[0m" # dim + strikethrough
      end
    "#{indent}#{right_aligned_id} | #{decorated_title}"
  end
end
