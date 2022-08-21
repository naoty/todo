require "pathname"

class Todo::Commands::Root < Todo::Commands::Command
  HELP_MESSAGE = <<~TEXT.freeze
    Usage:
      todo list
      todo add <title> (<position>) (-t | --tag <tag>)... (-p | --parent <id>) (-o | --open)
      todo open <id>
      todo move <id> <position> (-p | --parent <id>)
      todo delete <id>...
      todo done <id>...
      todo undone <id>...
      todo wait <id>...
      todo archive
      todo -h | --help
      todo -v | --version

    Options:
      -h --help     Show this message
      -v --version  Show version
  TEXT

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
    repository = Todo::FileRepository.new(
      root_path: ENV["TODOS_PATH"] || default_root_path,
      error_output: error_output
    )
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
      Todo::Commands::Add.new(arguments: arguments, output: output, error_output: error_output)
    when "list"
      Todo::Commands::List.new(arguments: arguments, output: output, error_output: error_output)
    when "open"
      Todo::Commands::Open.new(arguments: arguments, output: output, error_output: error_output)
    when "move"
      Todo::Commands::Move.new(arguments: arguments, output: output, error_output: error_output)
    when "delete"
      Todo::Commands::Delete.new(arguments: arguments, output: output, error_output: error_output)
    when "done"
      Todo::Commands::Update.new(arguments: arguments, state: :done, output: output, error_output: error_output)
    when "undone"
      Todo::Commands::Update.new(arguments: arguments, state: :undone, output: output, error_output: error_output)
    when "wait"
      Todo::Commands::Update.new(arguments: arguments, state: :waiting, name: "wait", output: output, error_output: error_output)
    when "archive"
      Todo::Commands::Archive.new(arguments: arguments, output: output, error_output: error_output)
    else
      raise CommandNotFound.new(unknown_name: name)
    end
  end

  def default_root_path
    Pathname.new(ENV.fetch("HOME")).join(".todos").to_s
  end
end
