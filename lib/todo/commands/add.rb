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
    if result.has_key?(:help)
      output.puts(HELP_MESSAGE)
      return
    elsif [:tags, :position, :parent_id, :title, :open].all? { |key| result.has_key?(key) }
      todo = repository.create(
        title: result[:title],
        tags: result[:tags],
        position: result[:position],
        parent_id: result[:parent_id]&.to_i
      )
      repository.open(id: todo.id) if result[:open]
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
