require "spec_helper"
require "stringio"

RSpec.describe Todo::Add do
  describe "#run" do
    let(:output) { StringIO.new }
    let(:error_output) { StringIO.new }
    let(:repository) { instance_double(Todo::FileRepository) }

    context "when arguments are empty" do
      it "puts help message to error output" do
        add = Todo::Add.new(arguments: [], output: output, error_output: error_output)
        add.run(repository: repository)
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq(Todo::Add::HELP_MESSAGE)
      end

      it "exits with status code 1" do
        add = Todo::Add.new(arguments: [], output: output, error_output: error_output)
        expect {
          add.run(repository: repository)
        }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          add = Todo::Add.new(arguments: [flag], output: output, error_output: error_output)
          add.run(repository: repository)
          expect(output.string).to eq(Todo::Add::HELP_MESSAGE)
        end
      end
    end

    context "when arguments include title" do
      it "calls FileRepository#create with title" do
        add = Todo::Add.new(arguments: ["dummy"], output: output, error_output: error_output)
        expect(repository).to receive(:create).with({title: "dummy"})
        add.run(repository: repository)
      end
    end

    ["-p", "--parent"].each do |option|
      context "when arguments include '#{option}' option" do
        it "calls FileRepository#create with title and parent_id" do
          add = Todo::Add.new(arguments: [option, "1", "dummy"], output: output, error_output: error_output)
          expect(repository).to receive(:create).with({title: "dummy", parent_id: 1})
          add.run(repository: repository)
        end
      end

      context "when arguments include '#{option}' option with invalid value" do
        it "puts help message to error output" do
          add = Todo::Add.new(arguments: [option, "dummy"], output: output, error_output: error_output)
          add.run(repository: repository)
        rescue SystemExit
          # ignore exit
        ensure
          expect(error_output.string).to eq(Todo::Add::HELP_MESSAGE)
        end

        it "exits with status code 1" do
          add = Todo::Add.new(arguments: [option, "dummy"], output: output, error_output: error_output)
          expect {
            add.run(repository: repository)
          }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
        end
      end
    end
  end
end
