class Todo::Commands::Move < Todo::Commands::Command
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo move <id> <position> (-p | --parent <parent id>)
      todo move -h | --help
    
    Options:
      -h --help  Show thid message
  TEXT

  def run(repository:)
    result = parse_arguments

    case result
    in { help: true }
      output.puts(HELP_MESSAGE)
      return
    in { invalid_id: id }
      error_output.puts("id is invalid: #{id}")
      exit 1
    in { invalid_position: position }
      error_output.puts("position is invalid: #{position}")
      exit 1
    in { empty_parent_id: true }
      error_output.puts("parent id is empty")
      exit 1
    in { invalid_parent_id: parent_id }
      error_output.puts("parent id is invalid: #{parent_id}")
      exit 1
    in { id: id, position: position, parent_id: parent_id }
      repository.move(id: id.to_i, parent_id: parent_id.to_i, position: position.to_i)
    in { id: id, position: position }
      repository.move(id: id.to_i, position: position.to_i)
    else
      error_output.puts(HELP_MESSAGE)
      exit 1
    end

    todos = repository.list
    print_todos(todos, indent_width: 2)
  end

  private

  def parse_arguments
    result = {}

    arguments.each.with_index do |argument, index|
      case argument
      when "-h", "--help"
        result[:help] = true
      when "-p", "--parent"
        if arguments[index + 1].nil?
          result[:empty_parent_id] = true
          break
        end

        if /\d+/.match?(arguments[index + 1])
          result[:parent_id] = arguments.delete_at(index + 1)
          next
        end

        result[:invalid_parent_id] = arguments.delete_at(index + 1)
      else
        if result[:id].nil?
          if /\d+/.match?(argument)
            result[:id] = argument
            next
          end

          result[:invalid_id] = argument
          break
        end

        if /\d+/.match?(argument)
          result[:position] = argument
          next
        end

        result[:invalid_position] = argument
      end
    end

    result
  end
end
