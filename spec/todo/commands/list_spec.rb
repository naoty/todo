require "spec_helper"
require "stringio"

RSpec.describe Todo::Commands::List do
  let(:output) { StringIO.new }
  let(:error_output) { StringIO.new }
  let(:repository) { instance_double(Todo::FileRepository) }

  let(:todos) do
    [
      Todo::Todo.new(id: 1, title: "dummy", state: :undone)
    ]
  end

  describe "#run" do
    context "when arguments are empty" do
      let(:arguments) { [] }

      it "calls FileRepository#list" do
        expect(repository).to receive(:list).and_return(todos)
        list = described_class.new(arguments: arguments, output: output, error_output: error_output)
        list.run(repository: repository)
      end

      it "calls Printable#print_todos" do
        allow(repository).to receive(:list).and_return(todos)
        list = described_class.new(arguments: arguments, output: output, error_output: error_output)
        expect(list).to receive(:print_todos)
        list.run(repository: repository)
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include #{flag} flag" do
        let(:arguments) { [flag] }

        it "puts help message" do
          allow(repository).to receive(:list).and_return(todos)
          list = described_class.new(arguments: arguments, output: output, error_output: error_output)
          list.run(repository: repository)
          expect(output.string).to eq(described_class::HELP_MESSAGE)
        end
      end
    end
  end
end
