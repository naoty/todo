require "spec_helper"
require "stringio"

RSpec.describe Todo::Delete do
  describe "#run" do
    let(:output) { StringIO.new }
    let(:error_output) { StringIO.new }
    let(:repository) { instance_double(Todo::FileRepository) }

    context "when arguments are empty" do
      it "puts help message to error output" do
        delete = Todo::Delete.new(arguments: [], output: output, error_output: error_output)
        delete.run(repository: repository)
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq(Todo::Delete::HELP_MESSAGE)
      end

      it "exits with status code 1" do
        add = Todo::Delete.new(arguments: [], output: output, error_output: error_output)
        expect {
          add.run(repository: repository)
        }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          add = Todo::Delete.new(arguments: [flag], output: output, error_output: error_output)
          add.run(repository: repository)
          expect(output.string).to eq(Todo::Delete::HELP_MESSAGE)
        end
      end
    end

    context "when arguments include IDs" do
      it "calls Todo::FileRepository#delete" do
        delete = Todo::Delete.new(arguments: ["1"], output: output, error_output: error_output)
        expect(repository).to receive(:delete).with(ids: [1])
        delete.run(repository: repository)
      end
    end
  end
end
