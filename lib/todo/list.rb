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
    puts_todos(todos, indent_width: 2)
  end

  private

  def puts_todos(todos, indent_width:)
    indent = " " * indent_width
    id_width = todos.map { |todo| todo.id.digits.length }.max

    todos.each do |todo|
      output.puts(format_todo(todo, indent: indent, id_width: id_width))
      puts_todos(todo.subtodos, indent_width: indent_width + id_width + 3) # " | " is 3 chars
    end
  end

  def format_todo(todo, indent:, id_width:)
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
