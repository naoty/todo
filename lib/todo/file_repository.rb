require "json"
require "pathname"

class Todo::FileRepository
  private attr_reader :root_path

  def initialize(root_path:)
    @root_path = Pathname.new(root_path)
    setup
  end

  private

  def setup
    create_index_if_not_exist
    create_archived_directory_if_not_exist
  end

  def create_index_if_not_exist
    index_path = root_path.join("index.json")
    return if index_path.exist?

    index_json = JSON.pretty_generate({
      todos: {},
      archived: {},
      metadata: {
        lastId: 0,
        missingIds: []
      }
    })
    index_path.open("wb") { |file| file.puts(index_json) }
  end

  def create_archived_directory_if_not_exist
    archived_path = root_path.join("archived")
    return if archived_path.exist?

    archived_path.mkdir
  end
end
