require "json"
require "pathname"

class Todo::FileRepository
  private attr_reader :root_path

  def initialize(root_path:)
    @root_path = Pathname.new(root_path)
    setup
  end

  def create(title:)
    todo = Todo::Todo.new(id: next_id, title: title, status: :undone, body: "")
    todo_path = root_path.join("#{todo.id}.md")
    encoded_todo = encode(todo)
    todo_path.open("wb") { |file| file.puts(encoded_todo) }

    index = load_index
    index[:todos][:""] ||= []
    index[:todos][:""] << todo.id
    index[:metadata][:lastId] = todo.id
    save_index(index)
  end

  private

  def setup
    create_index_if_not_exist
    create_archived_directory_if_not_exist
  end

  def create_index_if_not_exist
    index_path = root_path.join("index.json")
    return if index_path.exist?

    save_index({
      todos: {},
      archived: {},
      metadata: {
        lastId: 0,
        missingIds: []
      }
    })
  end

  def create_archived_directory_if_not_exist
    archived_path = root_path.join("archived")
    return if archived_path.exist?

    archived_path.mkdir
  end

  def next_id
    index = load_index
    last_id = index.dig(:metadata, :lastId) || 0
    last_id + 1
  end

  def load_index
    JSON.parse(root_path.join("index.json").read, symbolize_names: true)
  end

  def save_index(index)
    index_json = JSON.pretty_generate(index)
    root_path.join("index.json").open("wb") { |file| file.puts(index_json) }
  end

  def encode(todo)
    <<~TEXT
      ---
      title: #{todo.title}
      status: #{todo.status}
      ---

      #{todo.body}
    TEXT
  end
end
