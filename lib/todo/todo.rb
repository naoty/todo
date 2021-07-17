class Todo::Todo
  attr_reader :id, :title, :status, :body

  def initialize(id:, title:, status: :undone, body: "")
    @id = id
    @title = title
    @status = status
    @body = body
  end
end
