class Todo::Todo
  attr_accessor :state, :subtodos, :parent
  attr_reader :id, :title, :body

  def initialize(id:, title:, state: :undone, body: "", subtodos: [])
    @id = id
    @title = title
    @state = state
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
