class Todo::Commands::Add < Todo::Commands::Command
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo add <title> (<position>) (-t | --tag <tag>)... (-p | --parent <id>) (-o | --open)
      todo add -h | --help

    Options:
      -h --help    Show thid message
      -t --tag     Add tag
      -p --parent  Parent TODO ID
      -o --open    Open TODO file after create
  TEXT

  def run(repository:)
    result = parse_arguments(arguments)
    case result
    in { help: true }
      output.puts(HELP_MESSAGE)
      return
    in { tags: tags, position: position, parent_id: parent_id, title: title, open: open_flag }
      todo = repository.create(
        title: title,
        tags: tags,
        position: position,
        parent_id: parent_id&.to_i
      )
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
    result = {tags: [], position: nil, parent_id: nil, open: false}
    arguments_copy = arguments.dup
    arguments_copy.each.with_index do |argument, index|
      case argument
      when "-h", "--help"
        result[:help] = true
      when "-t", "--tag"
        result[:tags] << arguments_copy.delete_at(index + 1)
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
