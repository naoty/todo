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

    if result.has_key?(:help)
      output.puts(HELP_MESSAGE)
      return
    elsif result.has_key?(:invalid_id)
      error_output.puts("id is invalid: #{result[:invalid_id]}")
      exit 1
    elsif result.has_key?(:invalid_position)
      error_output.puts("position is invalid: #{result[:invalid_position]}")
      exit 1
    elsif result.has_key?(:empty_parent_id)
      error_output.puts("parent id is empty")
      exit 1
    elsif result.has_key?(:invalid_parent_id)
      error_output.puts("parent id is invalid: #{result[:invalid_parent_id]}")
      exit 1
    elsif [:id, :position, :parent_id].all? { |key| result.has_key?(key) }
      repository.move(
        id: result[:id].to_i,
        parent_id: result[:parent_id].to_i,
        position: result[:position].to_i
      )
    elsif [:id, :position].all? { |key| result.has_key?(key) }
      repository.move(
        id: result[:id].to_i,
        position: result[:position].to_i
      )
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
