module Todo::Commands::Printable
  def output
    raise NotImplementedError
  end

  def print_todos(todos, indent_width: 0)
    indent = " " * indent_width
    id_width = todos.map { |todo| todo.id.digits.length }.max

    todos.each do |todo|
      output.puts(format_todo(todo, indent: indent, id_width: id_width))
      print_todos(todo.subtodos, indent_width: indent_width + id_width + 3) # " | " is 3 chars
    end
  end

  private

  def format_todo(todo, indent:, id_width:)
    right_aligned_id = todo.id.to_s.rjust(id_width, " ")

    decorated_title =
      case todo.state
      when :undone then todo.title
      when :waiting then "\e[2m#{todo.title}" # dim
      when :done then "\e[2;9m#{todo.title}" # dim + strikethrough
      end

    decorated_tags = todo.tags
      .map { |tag| "\e[36m##{tag}" }
      .join(" ")

    result = "#{indent}#{right_aligned_id} | #{decorated_title}"
    result += " #{decorated_tags}" unless todo.tags.empty?
    result += "\e[0m" if [:waiting, :done].include?(todo.state) || !todo.tags.empty?
    result
  end
end
