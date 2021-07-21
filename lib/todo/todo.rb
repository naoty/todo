class Todo::Todo
  attr_accessor :subtodos
  attr_reader :id, :title, :state, :body

  def initialize(id:, title:, state: :undone, body: "", subtodos: [])
    @id = id
    @title = title
    @state = state
    @body = body
    @subtodos = subtodos
  end
end
