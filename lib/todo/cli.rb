class Todo::CLI
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo list
      todo add <title>
      todo -h | --help
      todo -v | --version
    
    Options:
      -h --help     Show this message
      -v --version  Show version
  TEXT

  private attr_reader :arguments, :output, :error_output

  def initialize(arguments: ARGV, output: $stdout, error_output: $stderr)
    @arguments = arguments
    @output = output
    @error_output = error_output
  end

  def run
    if arguments.empty?
      error_output.puts(HELP_MESSAGE)
      exit 1
    end

    if arguments.first == "-h" || arguments.first == "--help"
      output.puts(HELP_MESSAGE)
      return
    end

    if arguments.first == "-v" || arguments.first == "--version"
      output.puts(Todo::VERSION)
      return
    end

    command = build_command(name: arguments.first, arguments: arguments[1..])
    repository = Todo::FileRepository.new(root_path: ENV["TODOS_PATH"] || Pathname.pwd)
    command.run(repository: repository)
  rescue CommandNotFound => exception
    error_output.puts(exception.message)
    exit 1
  end

  private

  class CommandNotFound < StandardError
    attr_reader :unknown_name

    def initialize(unknown_name:)
      super
      @unknown_name = unknown_name
    end

    def message
      "command not found: #{unknown_name}"
    end
  end

  def build_command(name:, arguments:)
    case name
    when "add"
      Todo::Add.new(arguments: arguments)
    when "list"
      Todo::List.new(arguments: arguments)
    else
      raise CommandNotFound.new(unknown_name: name)
    end
  end
end
