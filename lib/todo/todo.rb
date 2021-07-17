class Todo::Todo
  def initialize(id:, title:, status: :undone, body: "")
    @id = id
    @title = title
    @status = status
    @body = body
  end
end
