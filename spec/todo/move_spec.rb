require "spec_helper"
require "stringio"

RSpec.describe Todo::Move do
  describe "#run" do
    let(:output) { StringIO.new }
    let(:error_output) { StringIO.new }
    let(:repository) { instance_double(Todo::FileRepository) }

    before do
      allow(repository).to receive(:list).and_return([])
    end

    context "when arguments are empty" do
      it "puts help message to error output" do
        move = Todo::Move.new(arguments: [], output: output, error_output: error_output)
        move.run(repository: repository)
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq(Todo::Move::HELP_MESSAGE)
      end

      it "exits with status code 1" do
        move = Todo::Move.new(arguments: [], output: output, error_output: error_output)
        expect {
          move.run(repository: repository)
        }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          move = Todo::Move.new(arguments: [flag], output: output, error_output: error_output)
          move.run(repository: repository)
          expect(output.string).to eq(Todo::Move::HELP_MESSAGE)
        end
      end
    end

    context "when arguments include only id" do
      it "puts help message to error output" do
        move = Todo::Move.new(arguments: ["1"], output: output, error_output: error_output)
        move.run(repository: repository)
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq(Todo::Move::HELP_MESSAGE)
      end

      it "exits with status code 1" do
        move = Todo::Move.new(arguments: ["1"], output: output, error_output: error_output)
        expect {
          move.run(repository: repository)
        }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
      end
    end

    context "when arguments include both id and position" do
      context "when id is invalid" do
        it "puts message to error output" do
          move = Todo::Move.new(arguments: ["dummy", "2"], output: output, error_output: error_output)
          move.run(repository: repository)
        rescue SystemExit
          # ignore exit
        ensure
          expect(error_output.string).to eq("id is invalid: dummy\n")
        end

        it "exits with status code 1" do
          move = Todo::Move.new(arguments: ["dummy", "2"], output: output, error_output: error_output)
          expect {
            move.run(repository: repository)
          }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
        end
      end

      context "when position is invalid" do
        it "puts message to error output" do
          move = Todo::Move.new(arguments: ["1", "dummy"], output: output, error_output: error_output)
          move.run(repository: repository)
        rescue SystemExit
          # ignore exit
        ensure
          expect(error_output.string).to eq("position is invalid: dummy\n")
        end

        it "exits with status code 1" do
          move = Todo::Move.new(arguments: ["1", "dummy"], output: output, error_output: error_output)
          expect {
            move.run(repository: repository)
          }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
        end
      end

      context "when both id and position is valid" do
        it "calls FileRepository#move with id and position" do
          expect(repository).to receive(:move).with(id: 1, position: 2)
          move = Todo::Move.new(arguments: ["1", "2"], output: output, error_output: error_output)
          move.run(repository: repository)
        end
      end
    end

    ["-p", "--parent"].each do |option|
      context "when arguments include '#{option} option'" do
        it "calls# FileRepository#move with id and position and parent_id" do
          expect(repository).to receive(:move).with(id: 1, position: 2, parent_id: 3)
          move = Todo::Move.new(arguments: ["1", "2", option, "3"], output: output, error_output: error_output)
          move.run(repository: repository)
        end
      end

      context "when arguments include '#{option}' option without value" do
        it "puts message to error output" do
          move = Todo::Move.new(arguments: ["1", "2", option], output: output, error_output: error_output)
          move.run(repository: repository)
        rescue SystemExit
          # ignore exit
        ensure
          expect(error_output.string).to eq("parent id is empty\n")
        end

        it "exits with status code 1" do
          move = Todo::Move.new(arguments: ["1", "2", option], output: output, error_output: error_output)
          expect {
            move.run(repository: repository)
          }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
        end
      end

      context "when arguments include '#{option}' option with invalid value" do
        it "puts help message to error output" do
          move = Todo::Move.new(arguments: ["1", "2", option, "dummy"], output: output, error_output: error_output)
          move.run(repository: repository)
        rescue SystemExit
          # ignore exit
        ensure
          expect(error_output.string).to eq("parent id is invalid: dummy\n")
        end

        it "exits with status code 1" do
          move = Todo::Move.new(arguments: ["1", "2", option, "dummy"], output: output, error_output: error_output)
          expect {
            move.run(repository: repository)
          }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
        end
      end
    end
  end
end
