require "spec_helper"

RSpec.describe Todo::Commands::Open do
  describe "#run" do
    let(:output) { StringIO.new }
    let(:error_output) { StringIO.new }
    let(:repository) { instance_double(Todo::FileRepository) }

    context "when arguments are empty" do
      it "puts help message to error output" do
        command = described_class.new(arguments: [], output: output, error_output: error_output)
        command.run(repository: repository)
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq(described_class::HELP_MESSAGE)
      end

      it "exits with status code 1" do
        command = described_class.new(arguments: [], output: output, error_output: error_output)
        expect {
          command.run(repository: repository)
        }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          command = described_class.new(arguments: [flag], output: output, error_output: error_output)
          command.run(repository: repository)
          expect(output.string).to eq(described_class::HELP_MESSAGE)
        end
      end
    end

    context "when arguments include an ID" do
      it "calls FileRepository#open with the ID" do
        expect(repository).to receive(:open).with(id: 1)
        command = described_class.new(arguments: ["1"], output: output, error_output: error_output)
        command.run(repository: repository)
      end
    end
  end
end
