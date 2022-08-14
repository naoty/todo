require "fileutils"
require "json"
require "pathname"
require "yaml"

class Todo::FileRepository
  private attr_reader :root_path, :error_output

  def initialize(root_path:, error_output:)
    @root_path = Pathname.new(root_path)
    @error_output = error_output

    setup
  end

  def list(id: nil)
    index = load_index
    todos = (index[:todos][id.to_s.to_sym] || []).map do |id|
      todo_path = root_path.join("#{id}.md")

      unless todo_path.exist?
        error_output.puts("todo file is not found: #{todo_path}")
        next nil
      end

      todo = decode(id: id, text: todo_path.read)

      if todo.nil?
        error_output.puts("todo file is broken: #{todo_path}")
        next nil
      end

      list(id: todo.id).each { |subtodo| todo.append_subtodo(subtodo) }
      todo
    end
    todos.compact
  end

  def create(title:, tags: [], position: nil, parent_id: nil)
    next_id = load_next_id

    todo = Todo::Todo.new(id: next_id, title: title, state: :undone, tags: tags, body: "")
    todo_path = root_path.join("#{todo.id}.md")
    encoded_todo = encode(todo)
    todo_path.open("wb") { |file| file.puts(encoded_todo) }

    insert_index = case position
      when nil then -1
      when ..0 then position
      else position - 1
    end

    index = load_index
    index[:todos][parent_id.to_s.to_sym] ||= []
    index[:todos][parent_id.to_s.to_sym].insert(insert_index, todo.id).compact!
    save_index(index)

    todo
  end

  def delete(ids:)
    index = load_index

    ids_to_be_deleted = ids.dup
    ids.each do |id|
      subtodo_ids = index[:todos][id.to_s.to_sym]
      next if subtodo_ids.nil?

      ids_to_be_deleted -= subtodo_ids
    end

    ids_to_be_deleted.each do |id|
      todo_path = root_path.join("#{id}.md")

      unless todo_path.exist?
        error_output.puts("todo file is not found: #{todo_path}")
        next
      end

      todo_path.delete

      index[:todos].each do |parent_id, subtodo_ids|
        if parent_id.to_s.to_i == id
          delete(ids: subtodo_ids)
          index[:todos].delete(id.to_s.to_sym)
        end

        if subtodo_ids.include?(id)
          index[:todos][parent_id].delete(id)
          index[:metadata][:missingIds] << id
        end
      end
    end

    save_index(index)
  end

  def update(ids:, state:)
    index = load_index

    ids_to_be_updated = ids.dup
    ids.each do |id|
      subtodo_ids = index[:todos][id.to_s.to_sym]
      next if subtodo_ids.nil?

      ids_to_be_updated -= subtodo_ids
    end

    ids_to_be_updated.each do |id|
      todo_path = root_path.join("#{id}.md")

      unless todo_path.exist?
        error_output.puts("todo file is not found: #{todo_path}")
        next
      end

      todo = decode(id: id, text: todo_path.read)
      todo.state = state
      encoded_todo = encode(todo)
      todo_path.open("wb") { |file| file.puts(encoded_todo) }

      subtodo_ids = index[:todos][id.to_s.to_sym]
      next if subtodo_ids.nil?

      update(ids: subtodo_ids, state: state)
    end
  end

  def archive(todos: list)
    index = load_index

    todos.each do |todo|
      if todo.should_be_archived?
        todo_path = root_path.join("#{todo.id}.md")
        archived_todo_path = root_path.join("archived", "#{todo.id}.md")
        FileUtils.mv(todo_path, archived_todo_path)

        index[:todos][todo.parent&.id.to_s.to_sym].delete(todo.id)
        index[:todos].delete(todo.parent&.id.to_s.to_sym) if index[:todos][todo.parent&.id.to_s.to_sym].empty?
        index[:archived][todo.parent&.id.to_s.to_sym] ||= []
        index[:archived][todo.parent&.id.to_s.to_sym] << todo.id
      end

      archive(todos: todo.subtodos)
    end

    save_index(index)
  end

  def move(id:, position:, parent_id: nil)
    if id == parent_id
      error_output.puts("moving a todo under itself is forbidden")
      return
    end

    index = load_index

    index[:archived].each do |_, subtodos|
      if subtodos.include?(id)
        error_output.puts("moving an archived todo is forbidden")
        return nil
      end

      if subtodos.include?(parent_id)
        error_output.puts("moving a todo under an archived todo is forbidden")
        return nil
      end
    end

    todo_found = false
    parent_found = false
    index[:todos].each do |parent_key, subtodos|
      if subtodos.include?(id)
        todo_found = true
        subtodos.delete(id)
        index[:todos].delete(parent_key) if subtodos.empty?
      end

      if subtodos.include?(parent_id)
        parent_found = true
      end
    end

    if !todo_found
      error_output.puts("todo is not found: #{id}")
      return
    end

    if !parent_id.nil? && !parent_found
      error_output.puts("parent is not found: #{parent_id}")
      return
    end

    insert_index = position > 0 ? position - 1 : position
    index[:todos][parent_id.to_s.to_sym] ||= []
    index[:todos][parent_id.to_s.to_sym].insert(insert_index, id).compact!

    save_index(index)
  end

  def open(id:)
    todo_path = root_path.join("#{id}.md")
    todo_path = root_path.join("archived", "#{id}.md") unless todo_path.exist?

    unless todo_path.exist?
      error_output.puts("todo is not found: #{id}")
      return
    end

    system("open #{todo_path}")
  end

  private

  def setup
    create_index_if_not_exist
    create_archived_directory_if_not_exist
  end

  def create_index_if_not_exist
    index_path = root_path.join("index.json")
    return if index_path.exist?

    save_index(default_index)
  end

  def create_archived_directory_if_not_exist
    archived_path = root_path.join("archived")
    return if archived_path.exist?

    archived_path.mkdir
  end

  def load_next_id
    index = load_index
    missing_ids = index.dig(:metadata, :missingIds)

    if missing_ids.empty?
      last_id = index.dig(:metadata, :lastId) || 0
      next_id = last_id + 1
      index[:metadata][:lastId] = next_id
    else
      next_id = index[:metadata][:missingIds].shift
    end

    save_index(index)
    next_id
  end

  def default_index
    {
      todos: {},
      archived: {},
      metadata: {
        lastId: 0,
        missingIds: []
      }
    }
  end

  def load_index
    @loaded_index ||= begin
      index_path = root_path.join("index.json")
      JSON.parse(index_path.read, symbolize_names: true)
    rescue JSON::ParserError
      error_output.puts("index file is broken: #{index_path}")
      default_index
    end
  end

  def save_index(index)
    index_json = JSON.pretty_generate(index)
    root_path.join("index.json").open("wb") { |file| file.puts(index_json) }
  end

  def encode(todo)
    metadata = {
      title: todo.title.match?(/[\[\]:]/) ? %("#{todo.title}") : todo.title,
      state: todo.state
    }

    unless todo.tags.empty?
      tags_string = todo.tags.map { |tag| %("#{tag}") }.join(", ")
      metadata[:tags] = %([#{tags_string}])
    end

    front_matter = metadata.map { |key, value| "#{key}: #{value}" }.join("\n")

    <<~TEXT
      ---
      #{front_matter}
      ---

      #{todo.body}
    TEXT
  end

  def decode(id:, text:)
    parts = text.split("---", 3).map(&:strip)

    return nil if parts.length < 3

    front_matter = YAML.safe_load(parts[1], symbolize_names: true)
    Todo::Todo.new(
      id: id,
      title: front_matter[:title],
      state: front_matter[:state]&.to_sym || :undone, # TODO: handle unknown status
      tags: front_matter[:tags] || [],
      body: parts[2]
    )
  rescue Psych::SyntaxError
    nil
  end
end
