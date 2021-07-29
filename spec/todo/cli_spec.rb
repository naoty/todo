require "spec_helper"
require "stringio"

RSpec.describe Todo::CLI do
  describe "#run" do
    let(:output) { StringIO.new }
    let(:error_output) { StringIO.new }

    shared_examples "exits with status code 1" do |arguments:|
      it "exits with status code 1" do
        cli = Todo::CLI.new(arguments: arguments, output: output, error_output: error_output)
        expect { cli.run }.to raise_error(an_instance_of(SystemExit).and(having_attributes(status: 1)))
      end
    end

    context "when arguments are empty" do
      it "puts help message to error output" do
        cli = Todo::CLI.new(arguments: [], output: output, error_output: error_output)
        cli.run
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq(Todo::CLI::HELP_MESSAGE)
      end

      include_examples "exits with status code 1", arguments: []
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          cli = Todo::CLI.new(arguments: [flag], output: output, error_output: error_output)
          cli.run
          expect(output.string).to eq(Todo::CLI::HELP_MESSAGE)
        end
      end
    end

    ["-v", "--version"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts version to output" do
          cli = Todo::CLI.new(arguments: [flag], output: output, error_output: error_output)
          cli.run
          expect(output.string).to eq("#{Todo::VERSION}\n")
        end
      end
    end

    context "when arguments include unknown command" do
      it "puts error message to error output" do
        cli = Todo::CLI.new(arguments: ["unknown"], output: output, error_output: error_output)
        cli.run
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq("command not found: unknown\n")
      end

      include_examples "exits with status code 1", arguments: ["unknown"]
    end

    {
      "add" => Todo::Add,
      "list" => Todo::List,
      "delete" => Todo::Delete,
      "done" => Todo::Update,
      "undone" => Todo::Update,
      "wait" => Todo::Update,
      "archive" => Todo::Archive
    }.each do |command, klass|
      context "when arguments include '#{command}' command" do
        it "calls #{klass}#run" do
          instance = instance_double(klass)
          allow(klass).to receive(:new).and_return(instance)
          expect(instance).to receive(:run)

          cli = Todo::CLI.new(arguments: [command], output: output, error_output: error_output)
          cli.run
        end
      end
    end
  end
end
