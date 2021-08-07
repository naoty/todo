module Todo::Printable
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
      when :waiting then "\e[2m#{todo.title}\e[0m" # dim
      when :done then "\e[2;9m#{todo.title}\e[0m" # dim + strikethrough
      end
    "#{indent}#{right_aligned_id} | #{decorated_title}"
  end
end
