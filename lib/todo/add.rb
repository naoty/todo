class Todo::Add
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo add (-p | --parent <id>) <title>
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
    result = parse_arguments(arguments)
    case result
    in { help: true }
      output.puts(HELP_MESSAGE)
    in { parent_id: parent_id, title: title }
      repository.create(title: title, parent_id: parent_id.to_i)
    in { title: title }
      repository.create(title: title)
    else
      error_output.puts(HELP_MESSAGE)
      exit 1
    end
  end

  private

  def parse_arguments(arguments)
    result = {}
    arguments_copy = arguments.dup
    arguments_copy.each.with_index do |argument, index|
      case argument
      when "-h", "--help"
        result[:help] = true
      when "-p", "--parent"
        result[:parent_id] = arguments_copy.delete_at(index + 1)
      else
        result[:title] = argument
      end
    end
    result
  end
end
