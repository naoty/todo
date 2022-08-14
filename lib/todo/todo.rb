class Todo::Todo
  attr_accessor :state, :subtodos, :parent
  attr_reader :id, :title, :tags, :body

  def initialize(id:, title:, state: :undone, tags: [], body: "", subtodos: [])
    @id = id
    @title = title
    @state = state
    @tags = tags
    @body = body
    @subtodos = subtodos
  end

  def append_subtodo(subtodo)
    subtodos << subtodo
    subtodo.parent = self
  end

  def done?
    state == :done
  end

  def should_be_archived?
    (!parent.nil? && parent.should_be_archived?) || done?
  end
end
