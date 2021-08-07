class Todo::Add
  include Todo::Printable

  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo add <title> (<position>) (-p | --parent <id>) (-o | --open)
      todo add -h | --help
    
    Options:
      -h --help    Show thid message
      -p --parent  Parent TODO ID
      -o --open    Open TODO file after create
  TEXT

  private attr_reader :arguments, :output, :error_output

  def initialize(arguments:, output: $stdout, error_output: $stderr)
    @arguments = arguments
    @output = output
    @error_output = error_output
  end

  def run(repository:)
    result = parse_arguments(arguments)
    case result
    in { help: true }
      output.puts(HELP_MESSAGE)
      return
    in { position: position, parent_id: parent_id, title: title, open: open_flag }
      todo = repository.create(title: title, position: position, parent_id: parent_id&.to_i)
      repository.open(id: todo.id) if open_flag
    else
      error_output.puts(HELP_MESSAGE)
      exit 1
    end

    todos = repository.list
    print_todos(todos, indent_width: 2)
  end

  private

  def parse_arguments(arguments)
    result = {position: nil, parent_id: nil, open: false}
    arguments_copy = arguments.dup
    arguments_copy.each.with_index do |argument, index|
      case argument
      when "-h", "--help"
        result[:help] = true
      when "-p", "--parent"
        result[:parent_id] = arguments_copy.delete_at(index + 1)
      when "-o", "--open"
        result[:open] = true
      else
        if result[:title].nil?
          result[:title] = argument
        else
          result[:position] = argument.to_i
        end
      end
    end
    result
  end
end
