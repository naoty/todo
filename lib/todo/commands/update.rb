class Todo::Commands::Update < Todo::Commands::Command
  attr_reader :state, :name

  def initialize(arguments:, state:, name: nil, output: $stdout, error_output: $stderr)
    super(arguments: arguments, output: output, error_output: error_output)
    @state = state
    @name = name || state
  end

  def run(repository:)
    if arguments.empty?
      error_output.puts(help_message)
      exit 1
    end

    if arguments.first == "-h" || arguments.first == "--help"
      output.puts(help_message)
      return
    end

    repository.update(ids: arguments.map(&:to_i), state: state)

    todos = repository.list
    print_todos(todos, indent_width: 2)
  end

  def help_message
    <<~TEXT
      Usage:
        todo #{name} <id>...
        todo #{name} -h | --help
      
      Options:
        -h --help  Show this message
    TEXT
  end
end
