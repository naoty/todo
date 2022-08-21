require "spec_helper"
require "stringio"

RSpec.describe Todo::Commands::Add do
  describe "#run" do
    let(:output) { StringIO.new }
    let(:error_output) { StringIO.new }
    let(:repository) { instance_double(Todo::FileRepository) }

    before do
      allow(repository).to receive(:list).and_return([])
    end

    context "when arguments are empty" do
      it "puts help message to error output" do
        add = described_class.new(arguments: [], output: output, error_output: error_output)
        add.run(repository: repository)
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq(described_class::HELP_MESSAGE)
      end

      it "exits with status code 1" do
        add = described_class.new(arguments: [], output: output, error_output: error_output)
        expect {
          add.run(repository: repository)
        }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          add = described_class.new(arguments: [flag], output: output, error_output: error_output)
          add.run(repository: repository)
          expect(output.string).to eq(described_class::HELP_MESSAGE)
        end
      end
    end

    context "when arguments include title" do
      it "calls FileRepository#create with title" do
        expect(repository).to receive(:create).with({title: "dummy", tags: [], position: nil, parent_id: nil})
        add = described_class.new(arguments: ["dummy"], output: output, error_output: error_output)
        add.run(repository: repository)
      end
    end

    context "when arguments include title and position" do
      it "calls FileRepository#create with title and position" do
        expect(repository).to receive(:create).with({title: "dummy", tags: [], position: 0, parent_id: nil})
        add = described_class.new(arguments: ["dummy", "0"], output: output, error_output: error_output)
        add.run(repository: repository)
      end
    end

    ["-t", "--tag"].each do |option|
      context "when arguments include '#{option}' option" do
        it "calls FileRepository#create with title and tag" do
          expect(repository).to receive(:create).with({title: "dummy", tags: ["dummy"], position: nil, parent_id: nil})
          add = described_class.new(arguments: [option, "dummy", "dummy"])
          add.run(repository: repository)
        end
      end

      context "when arguments include multiple '#{option}' options" do
        it "calls FileRepository#create with title and multiple tags" do
          expect(repository).to receive(:create).with({title: "dummy", tags: ["dummy1", "dummy2"], position: nil, parent_id: nil})
          add = described_class.new(arguments: [option, "dummy1", option, "dummy2", "dummy"])
          add.run(repository: repository)
        end
      end
    end

    ["-p", "--parent"].each do |option|
      context "when arguments include '#{option}' option" do
        it "calls FileRepository#create with title and parent_id" do
          expect(repository).to receive(:create).with({title: "dummy", tags: [], position: nil, parent_id: 1})
          add = described_class.new(arguments: [option, "1", "dummy"], output: output, error_output: error_output)
          add.run(repository: repository)
        end
      end

      context "when arguments include '#{option}' option with invalid value" do
        it "puts help message to error output" do
          add = described_class.new(arguments: [option, "dummy"], output: output, error_output: error_output)
          add.run(repository: repository)
        rescue SystemExit
          # ignore exit
        ensure
          expect(error_output.string).to eq(described_class::HELP_MESSAGE)
        end

        it "exits with status code 1" do
          add = described_class.new(arguments: [option, "dummy"], output: output, error_output: error_output)
          expect {
            add.run(repository: repository)
          }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
        end
      end
    end

    ["-o", "--open"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        let(:arguments) { ["dummy", flag] }
        let(:todo) { Todo::Todo.new(id: 1, title: "dummy") }

        it "calls FileRepository#open" do
          allow(repository).to receive(:create).and_return(todo)
          expect(repository).to receive(:open).with(id: 1)
          add = described_class.new(arguments: arguments, output: output, error_output: error_output)
          add.run(repository: repository)
        end
      end
    end
  end
end
