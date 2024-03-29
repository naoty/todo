require "spec_helper"

RSpec.describe Todo::Commands::Archive do
  describe "#run" do
    let(:output) { StringIO.new }
    let(:error_output) { StringIO.new }
    let(:repository) { instance_double(Todo::FileRepository) }

    before do
      allow(repository).to receive(:list).and_return([])
    end

    context "when arguments are empty" do
      it "calls FileRepository#archive" do
        expect(repository).to receive(:archive)
        archive = described_class.new(arguments: [], output: output, error_output: error_output)
        archive.run(repository: repository)
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          archive = described_class.new(arguments: [flag], output: output, error_output: error_output)
          archive.run(repository: repository)
          expect(output.string).to eq(described_class::HELP_MESSAGE)
        end
      end
    end
  end
end
